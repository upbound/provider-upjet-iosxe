package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// externalNameOverrides holds the external name configuration of the resources
// that deviate from the provider-wide default.
//
// Every Cisco IOS-XE resource manages a YANG object and is identified by the
// path of that object, e.g.
// "Cisco-IOS-XE-native:native/vrf/definition=VRF1". The Terraform provider
// computes that path from the resource arguments and reports it in the "id"
// attribute, which is also what "terraform import" consumes. That makes
// config.IdentifierFromProvider the correct configuration for all of them, and
// this map is expected to stay empty unless a resource starts deviating from
// that convention.
var externalNameOverrides = map[string]config.ExternalName{}

// ExternalNameConfigurations returns a ResourceOption that sets the external
// name configuration of every supported resource.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := externalNameOverrides[r.Name]; ok {
			r.ExternalName = e
			return
		}
		r.ExternalName = config.IdentifierFromProvider
	}
}
