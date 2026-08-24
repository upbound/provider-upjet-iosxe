# Provider Cisco IOS-XE

`provider-upjet-iosxe` manages the configuration of Cisco IOS-XE devices from
Kubernetes. It is generated with [Upjet](https://github.com/crossplane/upjet)
from the [Cisco IOS-XE Terraform provider](https://github.com/CiscoDevNet/terraform-provider-iosxe),
which is linked into the controller and called in-process over the Terraform
Plugin Framework: the controller talks NETCONF to the devices directly, with no
Terraform CLI and no Terraform state involved.

138 managed resources cover interfaces, VLANs and switching, VRFs and static
routing, BGP, OSPF, EIGRP, IS-IS, MPLS, EVPN, multicast, NAT, QoS, access lists,
crypto, AAA, device security and system, management and telemetry settings, plus
a generic YANG object for configuration without a dedicated resource.

Devices are addressed through a `ProviderConfig` that carries the NETCONF
credentials. A single `ProviderConfig` can front several devices, which managed
resources then select by name with `spec.forProvider.device`.

Both Crossplane v2 API flavours are available: namespaced resources in
`<group>.iosxe.m.upbound.io` and cluster scoped resources in
`<group>.iosxe.upbound.io`.
