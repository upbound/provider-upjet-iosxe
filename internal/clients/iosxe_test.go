package clients

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/upbound/provider-upjet-iosxe/config"
)

const (
	testUsername = "admin"
	testPassword = "t0ps3cr3t11"
	testDevice   = "leaf1"
)

func ptr[T any](v T) *T { return &v }

// providerType returns the Terraform type of the embedded provider's
// configuration schema.
func providerType(t *testing.T) tftypes.Type {
	t.Helper()
	var resp fwprovider.SchemaResponse
	config.FrameworkProvider().Schema(context.Background(), fwprovider.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("cannot read provider schema: %v", resp.Diagnostics)
	}
	return resp.Schema.Type().TerraformType(context.Background())
}

// TestProviderConfigurationIsCompleteAndTyped converts the provider
// configuration exactly the way upjet does at runtime, which fails if an
// attribute is missing or has the wrong type. It also asserts that every
// attribute of the provider schema is covered, so that an attribute added by a
// newer version of the Terraform provider is noticed here rather than at
// runtime.
func TestProviderConfigurationIsCompleteAndTyped(t *testing.T) {
	cases := map[string]credentials{
		"OnlyRequiredFields": {
			Username: testUsername,
			Password: testPassword,
		},
		"SingleDevice": {
			Username:           testUsername,
			Password:           testPassword,
			Host:               "10.0.0.1:830",
			Insecure:           ptr(true),
			Retries:            ptr(int64(5)),
			LockReleaseTimeout: ptr(int64(60)),
		},
		"MultipleDevices": {
			Username: testUsername,
			Password: testPassword,
			Host:     "10.0.0.1",
			Devices: []device{
				{Name: testDevice, Host: "10.0.0.1:830"},
				{Name: "leaf2", Host: "10.0.0.2:830", Managed: ptr(false)},
				{Name: "leaf3"},
			},
			SelectedDevices: []string{testDevice, "leaf2"},
		},
	}

	tfType := providerType(t)
	object, ok := tfType.(tftypes.Object)
	if !ok {
		t.Fatalf("provider schema is a %T, expected an object", tfType)
	}

	for name, creds := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := providerConfiguration(creds)

			for attribute := range object.AttributeTypes {
				if _, ok := cfg[attribute]; !ok {
					t.Errorf("provider configuration does not set attribute %q", attribute)
				}
			}
			for attribute := range cfg {
				if _, ok := object.AttributeTypes[attribute]; !ok {
					t.Errorf("provider configuration sets unknown attribute %q", attribute)
				}
			}

			// This is the conversion upjet performs before configuring the
			// provider, so a type mismatch fails here as well.
			data, err := json.Marshal(cfg)
			if err != nil {
				t.Fatalf("cannot marshal provider configuration: %v", err)
			}
			if _, err := tftypes.ValueFromJSONWithOpts(data, tfType, tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true}); err != nil {
				t.Errorf("cannot convert provider configuration to a Terraform value: %v", err)
			}
		})
	}
}

func TestProviderConfiguration(t *testing.T) {
	cases := map[string]struct {
		creds credentials
		want  map[string]any
	}{
		"UnsetOptionalFieldsAreNil": {
			creds: credentials{Username: testUsername, Password: testPassword},
			want: map[string]any{
				"username":             testUsername,
				"password":             testPassword,
				"host":                 nil,
				"insecure":             (*bool)(nil),
				"retries":              (*int64)(nil),
				"lock_release_timeout": (*int64)(nil),
				"reuse_connection":     false,
				"auto_commit":          true,
				"selected_devices":     nil,
				"devices":              nil,
			},
		},
		"DevicesAndSelection": {
			creds: credentials{
				Username:        testUsername,
				Password:        testPassword,
				Host:            "10.0.0.1",
				Devices:         []device{{Name: testDevice, Host: "10.0.0.1:830", Managed: ptr(true)}},
				SelectedDevices: []string{testDevice},
			},
			want: map[string]any{
				"username":             testUsername,
				"password":             testPassword,
				"host":                 "10.0.0.1",
				"insecure":             (*bool)(nil),
				"retries":              (*int64)(nil),
				"lock_release_timeout": (*int64)(nil),
				"reuse_connection":     false,
				"auto_commit":          true,
				"selected_devices":     []any{testDevice},
				"devices": []any{map[string]any{
					"name":    testDevice,
					"host":    "10.0.0.1:830",
					"managed": ptr(true),
				}},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := providerConfiguration(tc.creds)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("providerConfiguration(...): -want, +got:\n%s", diff)
			}
		})
	}
}
