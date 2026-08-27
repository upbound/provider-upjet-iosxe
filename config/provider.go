package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-upjet-iosxe/config/resources"
	"github.com/upbound/provider-upjet-iosxe/config/templates"
)

const (
	resourcePrefix = "iosxe"
	modulePath     = "github.com/upbound/provider-upjet-iosxe"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns the cluster-scoped provider configuration.
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("iosxe.upbound.io"),
		// This is a no-fork provider: every resource is reconciled in-process
		// via the Terraform Plugin Framework, so the CLI-based include list is
		// empty and all resources are registered in the Plugin Framework
		// include list. Note that upjet's default CLI include list matches all
		// resources, hence the explicit empty list here.
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginFrameworkIncludeList(resources.IncludeList()),
		ujconfig.WithTerraformPluginFrameworkProvider(FrameworkProvider()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithControllerTemplate(templates.ControllerTemplate),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			resources.SanitizeSensitiveFields(),
		))

	resources.Configure(pc)

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("iosxe.m.upbound.io"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginFrameworkIncludeList(resources.IncludeList()),
		ujconfig.WithTerraformPluginFrameworkProvider(FrameworkProvider()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithControllerTemplate(templates.ControllerTemplate),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			resources.SanitizeSensitiveFields(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	resources.Configure(pc)

	pc.ConfigureResources()
	return pc
}
