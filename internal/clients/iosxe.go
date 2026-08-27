package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/upbound/provider-upjet-iosxe/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-upjet-iosxe/apis/namespaced/v1beta1"
	"github.com/upbound/provider-upjet-iosxe/config"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal iosxe credentials as JSON"
)

// device is a single IOS-XE device reachable with the credentials of the
// ProviderConfig. Managed resources select a device by its name via their
// spec.forProvider.device field. A resource that does not set a device is
// reconciled against the device configured with the top level host.
type device struct {
	Name    string `json:"name"`
	Host    string `json:"host,omitempty"`
	Managed *bool  `json:"managed,omitempty"`
}

// credentials is the JSON document expected in the secret referenced by a
// ProviderConfig. It mirrors the configuration of the Cisco IOS-XE Terraform
// provider, which this provider embeds and configures in-process.
//
// The schema is documented for users in the "Credentials schema" section of
// README.md; keep the two in sync.
type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Host is the hostname or IP address of the device, optionally with a
	// ":port" suffix. The NETCONF default port is 830.
	Host string `json:"host,omitempty"`
	// Insecure skips SSH host key verification. It defaults to true in the
	// Terraform provider.
	Insecure *bool `json:"insecure,omitempty"`
	// Retries is the number of retries for NETCONF calls.
	Retries *int64 `json:"retries,omitempty"`
	// LockReleaseTimeout is the number of seconds to wait for the device
	// configuration lock to be released.
	LockReleaseTimeout *int64 `json:"lockReleaseTimeout,omitempty"`
	// Devices allows managing multiple devices with a single ProviderConfig.
	Devices []device `json:"devices,omitempty"`
	// SelectedDevices restricts reconciliation to a subset of Devices.
	SelectedDevices []string `json:"selectedDevices,omitempty"`
}

// TerraformSetupBuilder builds a terraform.SetupFn function which returns the
// Terraform provider setup configuration. The Cisco IOS-XE Terraform provider
// is a Terraform Plugin Framework provider that is linked into this binary and
// invoked in-process, so the setup carries the provider instance itself. There
// is no Terraform CLI to run and therefore no provider plugin requirement to
// declare.
func TerraformSetupBuilder() terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			FrameworkProvider: config.FrameworkProvider(),
		}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := credentials{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = providerConfiguration(creds)
		return ps, nil
	}
}

// providerConfiguration converts the credentials into the configuration of the
// embedded Terraform provider. Every attribute of the provider schema is set,
// with a nil value for the ones that are left to the defaults of the Terraform
// provider.
func providerConfiguration(creds credentials) map[string]any {
	devices := make([]any, 0, len(creds.Devices))
	for _, d := range creds.Devices {
		devices = append(devices, map[string]any{
			"name":    d.Name,
			"host":    stringOrNil(d.Host),
			"managed": d.Managed,
		})
	}

	cfg := map[string]any{
		"username":             creds.Username,
		"password":             creds.Password,
		"host":                 stringOrNil(creds.Host),
		"insecure":             creds.Insecure,
		"retries":              creds.Retries,
		"lock_release_timeout": creds.LockReleaseTimeout,
		// Upjet configures the embedded provider once per reconciliation, so a
		// NETCONF session that is kept open between operations would never be
		// closed. Connection reuse is therefore disabled, which makes the
		// provider open and close a session per operation.
		"reuse_connection": false,
		// Manual commit mode requires connection reuse and an explicit commit,
		// which has no declarative equivalent here, so changes are always
		// committed by the operation that makes them.
		"auto_commit":      true,
		"selected_devices": nil,
		"devices":          nil,
	}
	if len(devices) > 0 {
		cfg["devices"] = devices
	}
	if len(creds.SelectedDevices) > 0 {
		selected := make([]any, 0, len(creds.SelectedDevices))
		for _, s := range creds.SelectedDevices {
			selected = append(selected, s)
		}
		cfg["selected_devices"] = selected
	}
	return cfg
}

// stringOrNil returns nil for an empty string so that the attribute is left
// unset in the Terraform provider configuration, which lets the provider fall
// back to its own defaults or environment variables.
func stringOrNil(s string) any {
	if s == "" {
		return nil
	}
	return s
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged: //nolint:staticcheck // still handling cluster-scoped behavior
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // still handling cluster-scoped behavior
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
