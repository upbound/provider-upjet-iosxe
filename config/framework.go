package config

import (
	"github.com/CiscoDevNet/terraform-provider-iosxe/xpprovider"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
)

// TerraformProviderVersion is the version of the Cisco IOS-XE Terraform
// provider embedded in this Crossplane provider. Keep it in sync with
// TERRAFORM_PROVIDER_VERSION in the Makefile.
const TerraformProviderVersion = "1.0.0"

// FrameworkProvider returns a Terraform Plugin Framework provider instance of
// the Cisco IOS-XE Terraform provider. It is used both at code generation time
// (to read the resource schemas) and at runtime (to reconcile resources
// in-process, without forking the Terraform CLI).
func FrameworkProvider() fwprovider.Provider {
	return xpprovider.GetProvider(TerraformProviderVersion)
}
