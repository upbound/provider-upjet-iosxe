package resources

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// SanitizeSensitiveFields clears the Terraform "sensitive" flag of the fields
// whose type cannot be represented as a secret reference.
//
// Upjet turns a sensitive field into a reference to a key of a Kubernetes
// secret, which it only supports for string typed fields and which fails code
// generation for anything else. IOS-XE marks a few booleans as sensitive, for
// example the "password_secret" flag of iosxe_vtp, which only records whether
// the VTP password is stored in encrypted form and is not a secret itself.
// Those flags are kept as regular fields of the managed resource.
func SanitizeSensitiveFields() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		if r.TerraformResource == nil {
			return
		}
		clearUnsupportedSensitive(r.TerraformResource.Schema)
	}
}

func clearUnsupportedSensitive(s map[string]*schema.Schema) {
	for _, f := range s {
		if f == nil {
			continue
		}
		if f.Sensitive && !secretRefSupported(f) {
			f.Sensitive = false
		}
		if el, ok := f.Elem.(*schema.Resource); ok && el != nil {
			clearUnsupportedSensitive(el.Schema)
		}
	}
}

// secretRefSupported reports whether upjet can generate a secret reference for
// the given field, i.e. whether it is a string, a list or set of strings, or a
// map of strings.
func secretRefSupported(f *schema.Schema) bool {
	switch f.Type {
	case schema.TypeString:
		return true
	case schema.TypeList, schema.TypeSet, schema.TypeMap:
		el, ok := f.Elem.(*schema.Schema)
		return ok && el != nil && el.Type == schema.TypeString
	case schema.TypeInvalid, schema.TypeBool, schema.TypeInt, schema.TypeFloat:
		return false
	default:
		return false
	}
}
