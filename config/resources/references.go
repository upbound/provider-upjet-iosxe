package resources

import (
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// extractName extracts the "name" parameter of the referenced managed
// resource. IOS-XE external names are YANG paths, so a reference has to point
// at the parameter that the referring resource expects, not at the external
// name.
const extractName = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",false)`

// extractASN extracts the "asn" parameter of the referenced managed resource.
const extractASN = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("asn",false)`

// The fields naming the VRF a resource is configured in.
const (
	fieldVRF           = "vrf"
	fieldVRFForwarding = "vrf_forwarding"
)

// vrfReferences maps the resources that are configured inside a VRF to the
// field naming that VRF.
var vrfReferences = map[string]string{
	"iosxe_bgp_ipv4_unicast_vrf_neighbor":       fieldVRF,
	"iosxe_eigrp_vrf":                           fieldVRF,
	"iosxe_interface_ethernet":                  fieldVRFForwarding,
	"iosxe_interface_loopback":                  fieldVRFForwarding,
	"iosxe_interface_port_channel":              fieldVRFForwarding,
	"iosxe_interface_port_channel_subinterface": fieldVRFForwarding,
	"iosxe_interface_tunnel":                    fieldVRFForwarding,
	"iosxe_interface_vlan":                      fieldVRFForwarding,
	"iosxe_ospf_vrf":                            fieldVRF,
	"iosxe_ospfv3_address_family_ipv4_vrf":      fieldVRF,
	"iosxe_ospfv3_address_family_ipv6_vrf":      fieldVRF,
	"iosxe_static_routes_vrf":                   fieldVRF,
}

// configureReferences adds the cross-resource references of a resource, so
// that a managed resource can point at the VRF or the BGP process it belongs
// to instead of repeating its name, and so that Crossplane waits for that
// parent configuration to exist first.
func configureReferences(name string, r *ujconfig.Resource) {
	if field, ok := vrfReferences[name]; ok {
		r.References[field] = ujconfig.Reference{
			TerraformName: "iosxe_vrf",
			Extractor:     extractName,
		}
	}

	// Every BGP resource other than the BGP process itself configures a part
	// of that process and identifies it by its autonomous system number.
	if strings.HasPrefix(name, "iosxe_bgp_") {
		r.References["asn"] = ujconfig.Reference{
			TerraformName: "iosxe_bgp",
			Extractor:     extractASN,
		}
	}
}
