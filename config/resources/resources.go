// Package resources holds the mapping between the Cisco IOS-XE Terraform
// resources supported by this provider and the Crossplane API groups and kinds
// they are generated into.
package resources

import (
	"sort"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// The API groups of the provider. The full API group of a resource is the
// short group joined with the root group, e.g. "bgp.iosxe.upbound.io".
const (
	groupAAA       = "aaa"
	groupACL       = "acl"
	groupBFD       = "bfd"
	groupBGP       = "bgp"
	groupCrypto    = "crypto"
	groupEIGRP     = "eigrp"
	groupEVPN      = "evpn"
	groupFlow      = "flow"
	groupInterface = "interface"
	groupISIS      = "isis"
	groupMPLS      = "mpls"
	groupMulticast = "multicast"
	groupNAT       = "nat"
	groupOSPF      = "ospf"
	groupQOS       = "qos"
	groupRouting   = "routing"
	groupSecurity  = "security"
	groupSwitching = "switching"
	groupSystem    = "system"
	groupYANG      = "yang"
)

// resourceConfig is the Crossplane API surface of a Terraform resource.
type resourceConfig struct {
	// shortGroup is the API group without the provider suffix, e.g. "bgp"
	// results in the API group "bgp.iosxe.upbound.io".
	shortGroup string
	// kind is the Kubernetes kind of the managed resource.
	kind string
}

// resourceConfigs is the set of supported Terraform resources, keyed by their
// Terraform resource name.
//
// Upjet derives the API group and the kind from the Terraform resource name by
// default, which does not work well for IOS-XE: the group would become the
// first word after the "iosxe_" prefix, producing groups such as "access" for
// iosxe_access_list_standard or "as" for iosxe_as_path_access_list. Every
// supported resource is therefore mapped explicitly here, which also makes
// this table the include list of the provider: a resource that is not listed
// is not generated.
//
// The following Terraform resources are intentionally not exposed as managed
// resources, because they are imperative one-shot operations rather than a
// declarative piece of device configuration that can be observed and
// reconciled: iosxe_cli, iosxe_commit and iosxe_save_config.
var resourceConfigs = map[string]resourceConfig{
	// AAA, RADIUS, TACACS and local users.
	"iosxe_aaa":                {shortGroup: groupAAA, kind: "AAA"},
	"iosxe_aaa_accounting":     {shortGroup: groupAAA, kind: "Accounting"},
	"iosxe_aaa_authentication": {shortGroup: groupAAA, kind: "Authentication"},
	"iosxe_aaa_authorization":  {shortGroup: groupAAA, kind: "Authorization"},
	"iosxe_radius":             {shortGroup: groupAAA, kind: "RADIUS"},
	"iosxe_radius_server":      {shortGroup: groupAAA, kind: "RADIUSServer"},
	"iosxe_tacacs":             {shortGroup: groupAAA, kind: "TACACS"},
	"iosxe_tacacs_server":      {shortGroup: groupAAA, kind: "TACACSServer"},
	"iosxe_username":           {shortGroup: groupAAA, kind: "Username"},

	// Access lists and object groups.
	"iosxe_access_list_extended":   {shortGroup: groupACL, kind: "Extended"},
	"iosxe_access_list_ipv6":       {shortGroup: groupACL, kind: "IPv6"},
	"iosxe_access_list_role_based": {shortGroup: groupACL, kind: "RoleBased"},
	"iosxe_access_list_standard":   {shortGroup: groupACL, kind: "Standard"},
	"iosxe_object_group":           {shortGroup: groupACL, kind: "ObjectGroup"},

	// Bidirectional Forwarding Detection.
	"iosxe_bfd":                     {shortGroup: groupBFD, kind: "BFD"},
	"iosxe_bfd_template_multi_hop":  {shortGroup: groupBFD, kind: "TemplateMultiHop"},
	"iosxe_bfd_template_single_hop": {shortGroup: groupBFD, kind: "TemplateSingleHop"},

	// BGP.
	"iosxe_bgp":                           {shortGroup: groupBGP, kind: "BGP"},
	"iosxe_bgp_address_family_ipv4":       {shortGroup: groupBGP, kind: "AddressFamilyIPv4"},
	"iosxe_bgp_address_family_ipv4_mvpn":  {shortGroup: groupBGP, kind: "AddressFamilyIPv4MVPN"},
	"iosxe_bgp_address_family_ipv4_vrf":   {shortGroup: groupBGP, kind: "AddressFamilyIPv4VRF"},
	"iosxe_bgp_address_family_ipv6":       {shortGroup: groupBGP, kind: "AddressFamilyIPv6"},
	"iosxe_bgp_address_family_ipv6_vrf":   {shortGroup: groupBGP, kind: "AddressFamilyIPv6VRF"},
	"iosxe_bgp_address_family_l2vpn":      {shortGroup: groupBGP, kind: "AddressFamilyL2VPN"},
	"iosxe_bgp_address_family_vpnv4":      {shortGroup: groupBGP, kind: "AddressFamilyVPNv4"},
	"iosxe_bgp_address_family_vpnv6":      {shortGroup: groupBGP, kind: "AddressFamilyVPNv6"},
	"iosxe_bgp_bmp_server":                {shortGroup: groupBGP, kind: "BMPServer"},
	"iosxe_bgp_ipv4_mvpn_neighbor":        {shortGroup: groupBGP, kind: "IPv4MVPNNeighbor"},
	"iosxe_bgp_ipv4_unicast_neighbor":     {shortGroup: groupBGP, kind: "IPv4UnicastNeighbor"},
	"iosxe_bgp_ipv4_unicast_vrf_neighbor": {shortGroup: groupBGP, kind: "IPv4UnicastVRFNeighbor"},
	"iosxe_bgp_ipv6_unicast_neighbor":     {shortGroup: groupBGP, kind: "IPv6UnicastNeighbor"},
	"iosxe_bgp_l2vpn_evpn_neighbor":       {shortGroup: groupBGP, kind: "L2VPNEVPNNeighbor"},
	"iosxe_bgp_neighbor":                  {shortGroup: groupBGP, kind: "Neighbor"},
	"iosxe_bgp_peer_policy_template":      {shortGroup: groupBGP, kind: "PeerPolicyTemplate"},
	"iosxe_bgp_peer_session_template":     {shortGroup: groupBGP, kind: "PeerSessionTemplate"},

	// IKEv2, IPsec and PKI.
	"iosxe_crypto":                     {shortGroup: groupCrypto, kind: "Crypto"},
	"iosxe_crypto_ikev2":               {shortGroup: groupCrypto, kind: "IKEv2"},
	"iosxe_crypto_ikev2_keyring":       {shortGroup: groupCrypto, kind: "IKEv2Keyring"},
	"iosxe_crypto_ikev2_policy":        {shortGroup: groupCrypto, kind: "IKEv2Policy"},
	"iosxe_crypto_ikev2_profile":       {shortGroup: groupCrypto, kind: "IKEv2Profile"},
	"iosxe_crypto_ikev2_proposal":      {shortGroup: groupCrypto, kind: "IKEv2Proposal"},
	"iosxe_crypto_ipsec_profile":       {shortGroup: groupCrypto, kind: "IPSecProfile"},
	"iosxe_crypto_ipsec_transform_set": {shortGroup: groupCrypto, kind: "IPSecTransformSet"},
	"iosxe_crypto_pki":                 {shortGroup: groupCrypto, kind: "PKI"},

	// EIGRP.
	"iosxe_eigrp":     {shortGroup: groupEIGRP, kind: "EIGRP"},
	"iosxe_eigrp_vrf": {shortGroup: groupEIGRP, kind: "VRF"},

	// EVPN and L2 VPN.
	"iosxe_evpn":                  {shortGroup: groupEVPN, kind: "EVPN"},
	"iosxe_evpn_ethernet_segment": {shortGroup: groupEVPN, kind: "EthernetSegment"},
	"iosxe_evpn_instance":         {shortGroup: groupEVPN, kind: "Instance"},
	"iosxe_l2_vfi":                {shortGroup: groupEVPN, kind: "L2VFI"},

	// Flexible NetFlow.
	"iosxe_flow_exporter": {shortGroup: groupFlow, kind: "Exporter"},
	"iosxe_flow_monitor":  {shortGroup: groupFlow, kind: "Monitor"},
	"iosxe_flow_record":   {shortGroup: groupFlow, kind: "Record"},

	// Interfaces.
	"iosxe_interface_bdi":                       {shortGroup: groupInterface, kind: "BDI"},
	"iosxe_interface_ethernet":                  {shortGroup: groupInterface, kind: "Ethernet"},
	"iosxe_interface_isis":                      {shortGroup: groupInterface, kind: "ISIS"},
	"iosxe_interface_loopback":                  {shortGroup: groupInterface, kind: "Loopback"},
	"iosxe_interface_mpls":                      {shortGroup: groupInterface, kind: "MPLS"},
	"iosxe_interface_nve":                       {shortGroup: groupInterface, kind: "NVE"},
	"iosxe_interface_ospf":                      {shortGroup: groupInterface, kind: "OSPF"},
	"iosxe_interface_ospfv3":                    {shortGroup: groupInterface, kind: "OSPFv3"},
	"iosxe_interface_pim":                       {shortGroup: groupInterface, kind: "PIM"},
	"iosxe_interface_pim_ipv6":                  {shortGroup: groupInterface, kind: "PIMIPv6"},
	"iosxe_interface_port_channel":              {shortGroup: groupInterface, kind: "PortChannel"},
	"iosxe_interface_port_channel_subinterface": {shortGroup: groupInterface, kind: "PortChannelSubinterface"},
	"iosxe_interface_stackwise_virtual":         {shortGroup: groupInterface, kind: "StackwiseVirtual"},
	"iosxe_interface_switchport":                {shortGroup: groupInterface, kind: "Switchport"},
	"iosxe_interface_tunnel":                    {shortGroup: groupInterface, kind: "Tunnel"},
	"iosxe_interface_vlan":                      {shortGroup: groupInterface, kind: "VLAN"},
	"iosxe_interface_vrrp_v2":                   {shortGroup: groupInterface, kind: "VRRPv2"},

	// IS-IS.
	"iosxe_isis": {shortGroup: groupISIS, kind: "ISIS"},

	// MPLS.
	"iosxe_mpls": {shortGroup: groupMPLS, kind: "MPLS"},

	// Multicast routing.
	"iosxe_msdp":      {shortGroup: groupMulticast, kind: "MSDP"},
	"iosxe_multicast": {shortGroup: groupMulticast, kind: "Multicast"},
	"iosxe_pim":       {shortGroup: groupMulticast, kind: "PIM"},
	"iosxe_pim_ipv6":  {shortGroup: groupMulticast, kind: "PIMIPv6"},

	// NAT.
	"iosxe_nat": {shortGroup: groupNAT, kind: "NAT"},

	// OSPF and OSPFv3.
	"iosxe_ospf":                           {shortGroup: groupOSPF, kind: "OSPF"},
	"iosxe_ospf_vrf":                       {shortGroup: groupOSPF, kind: "VRF"},
	"iosxe_ospfv3":                         {shortGroup: groupOSPF, kind: "OSPFv3"},
	"iosxe_ospfv3_address_family_ipv4_vrf": {shortGroup: groupOSPF, kind: "OSPFv3AddressFamilyIPv4VRF"},
	"iosxe_ospfv3_address_family_ipv6_vrf": {shortGroup: groupOSPF, kind: "OSPFv3AddressFamilyIPv6VRF"},

	// Quality of service.
	"iosxe_class_map":        {shortGroup: groupQOS, kind: "ClassMap"},
	"iosxe_policy_map":       {shortGroup: groupQOS, kind: "PolicyMap"},
	"iosxe_policy_map_event": {shortGroup: groupQOS, kind: "PolicyMapEvent"},
	"iosxe_qos":              {shortGroup: groupQOS, kind: "QOS"},

	// Routing, VRFs and routing policy.
	"iosxe_arp":                           {shortGroup: groupRouting, kind: "ARP"},
	"iosxe_as_path_access_list":           {shortGroup: groupRouting, kind: "ASPathAccessList"},
	"iosxe_community_list_expanded":       {shortGroup: groupRouting, kind: "CommunityListExpanded"},
	"iosxe_community_list_standard":       {shortGroup: groupRouting, kind: "CommunityListStandard"},
	"iosxe_ipv6_local_pool":               {shortGroup: groupRouting, kind: "IPv6LocalPool"},
	"iosxe_ipv6_prefix_list":              {shortGroup: groupRouting, kind: "IPv6PrefixList"},
	"iosxe_large_community_list_expanded": {shortGroup: groupRouting, kind: "LargeCommunityListExpanded"},
	"iosxe_prefix_list":                   {shortGroup: groupRouting, kind: "PrefixList"},
	"iosxe_route_map":                     {shortGroup: groupRouting, kind: "RouteMap"},
	"iosxe_static_route":                  {shortGroup: groupRouting, kind: "StaticRoute"},
	"iosxe_static_routes_vrf":             {shortGroup: groupRouting, kind: "VRFStaticRoutes"},
	"iosxe_vrf":                           {shortGroup: groupRouting, kind: "VRF"},

	// Device security and identity.
	"iosxe_cts":                {shortGroup: groupSecurity, kind: "CTS"},
	"iosxe_device_sensor":      {shortGroup: groupSecurity, kind: "DeviceSensor"},
	"iosxe_device_tracking":    {shortGroup: groupSecurity, kind: "DeviceTracking"},
	"iosxe_dot1x":              {shortGroup: groupSecurity, kind: "Dot1x"},
	"iosxe_key_chain":          {shortGroup: groupSecurity, kind: "KeyChain"},
	"iosxe_parameter_map":      {shortGroup: groupSecurity, kind: "ParameterMap"},
	"iosxe_service_template":   {shortGroup: groupSecurity, kind: "ServiceTemplate"},
	"iosxe_zone_pair_security": {shortGroup: groupSecurity, kind: "ZonePair"},
	"iosxe_zone_security":      {shortGroup: groupSecurity, kind: "Zone"},

	// Layer 2 switching.
	"iosxe_bridge_domain":     {shortGroup: groupSwitching, kind: "BridgeDomain"},
	"iosxe_errdisable":        {shortGroup: groupSwitching, kind: "Errdisable"},
	"iosxe_spanning_tree":     {shortGroup: groupSwitching, kind: "SpanningTree"},
	"iosxe_stackwise_virtual": {shortGroup: groupSwitching, kind: "StackwiseVirtual"},
	// The kind of iosxe_switch cannot be "Switch": upjet names the generated
	// controller package after the kind, and "switch" is a Go keyword. The
	// resource configures the provisioning of a switch in a stack.
	"iosxe_switch":             {shortGroup: groupSwitching, kind: "SwitchProvision"},
	"iosxe_udld":               {shortGroup: groupSwitching, kind: "UDLD"},
	"iosxe_vlan":               {shortGroup: groupSwitching, kind: "VLAN"},
	"iosxe_vlan_access_map":    {shortGroup: groupSwitching, kind: "VLANAccessMap"},
	"iosxe_vlan_configuration": {shortGroup: groupSwitching, kind: "VLANConfiguration"},
	"iosxe_vlan_filter":        {shortGroup: groupSwitching, kind: "VLANFilter"},
	"iosxe_vlan_group":         {shortGroup: groupSwitching, kind: "VLANGroup"},
	"iosxe_vtp":                {shortGroup: groupSwitching, kind: "VTP"},

	// Device wide, management and telemetry configuration.
	"iosxe_banner":           {shortGroup: groupSystem, kind: "Banner"},
	"iosxe_cdp":              {shortGroup: groupSystem, kind: "CDP"},
	"iosxe_clock":            {shortGroup: groupSystem, kind: "Clock"},
	"iosxe_dhcp":             {shortGroup: groupSystem, kind: "DHCP"},
	"iosxe_eem":              {shortGroup: groupSystem, kind: "EEM"},
	"iosxe_license":          {shortGroup: groupSystem, kind: "License"},
	"iosxe_line":             {shortGroup: groupSystem, kind: "Line"},
	"iosxe_lldp":             {shortGroup: groupSystem, kind: "LLDP"},
	"iosxe_logging":          {shortGroup: groupSystem, kind: "Logging"},
	"iosxe_mdt_subscription": {shortGroup: groupSystem, kind: "MDTSubscription"},
	"iosxe_monitor_session":  {shortGroup: groupSystem, kind: "MonitorSession"},
	"iosxe_ntp":              {shortGroup: groupSystem, kind: "NTP"},
	"iosxe_platform":         {shortGroup: groupSystem, kind: "Platform"},
	"iosxe_service":          {shortGroup: groupSystem, kind: "Service"},
	"iosxe_sla":              {shortGroup: groupSystem, kind: "SLA"},
	"iosxe_snmp_server":      {shortGroup: groupSystem, kind: "SNMPServer"},
	"iosxe_system":           {shortGroup: groupSystem, kind: "System"},
	"iosxe_template":         {shortGroup: groupSystem, kind: "Template"},

	// Generic YANG object, an escape hatch for configuration that is not
	// covered by a dedicated resource.
	"iosxe_yang": {shortGroup: groupYANG, kind: "Object"},
}

// IncludeList returns the exact match regular expressions of the Terraform
// resources supported by this provider.
func IncludeList() []string {
	l := make([]string, 0, len(resourceConfigs))
	for name := range resourceConfigs {
		// "$" is added to match the exact resource name since the entries are
		// evaluated as regular expressions.
		l = append(l, name+"$")
	}
	// The generated code must not depend on map iteration order.
	sort.Strings(l)
	return l
}

// Configure registers the API group and the kind of every supported resource.
// The cluster-scoped and the namespaced providers share the same resource
// configuration.
func Configure(p *ujconfig.Provider) {
	for name, rc := range resourceConfigs {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.ShortGroup = rc.shortGroup
			r.Kind = rc.kind
			configureReferences(name, r)
		})
	}
}
