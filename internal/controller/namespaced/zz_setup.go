// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aaa "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/aaa"
	accounting "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/accounting"
	authentication "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/authentication"
	authorization "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/authorization"
	radius "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/radius"
	radiusserver "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/radiusserver"
	tacacs "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/tacacs"
	tacacsserver "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/tacacsserver"
	username "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/aaa/username"
	extended "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/acl/extended"
	ipv6 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/acl/ipv6"
	objectgroup "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/acl/objectgroup"
	rolebased "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/acl/rolebased"
	standard "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/acl/standard"
	bfd "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bfd/bfd"
	templatemultihop "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bfd/templatemultihop"
	templatesinglehop "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bfd/templatesinglehop"
	addressfamilyipv4 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyipv4"
	addressfamilyipv4mvpn "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyipv4mvpn"
	addressfamilyipv4vrf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyipv4vrf"
	addressfamilyipv6 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyipv6"
	addressfamilyipv6vrf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyipv6vrf"
	addressfamilyl2vpn "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyl2vpn"
	addressfamilyvpnv4 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyvpnv4"
	addressfamilyvpnv6 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/addressfamilyvpnv6"
	bgp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/bgp"
	bmpserver "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/bmpserver"
	ipv4mvpnneighbor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/ipv4mvpnneighbor"
	ipv4unicastneighbor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/ipv4unicastneighbor"
	ipv4unicastvrfneighbor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/ipv4unicastvrfneighbor"
	ipv6unicastneighbor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/ipv6unicastneighbor"
	l2vpnevpnneighbor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/l2vpnevpnneighbor"
	neighbor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/neighbor"
	peerpolicytemplate "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/peerpolicytemplate"
	peersessiontemplate "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/bgp/peersessiontemplate"
	crypto "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/crypto"
	ikev2 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ikev2"
	ikev2keyring "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ikev2keyring"
	ikev2policy "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ikev2policy"
	ikev2profile "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ikev2profile"
	ikev2proposal "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ikev2proposal"
	ipsecprofile "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ipsecprofile"
	ipsectransformset "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/ipsectransformset"
	pki "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/crypto/pki"
	eigrp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/eigrp/eigrp"
	vrf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/eigrp/vrf"
	ethernetsegment "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/evpn/ethernetsegment"
	evpn "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/evpn/evpn"
	instance "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/evpn/instance"
	l2vfi "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/evpn/l2vfi"
	exporter "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/flow/exporter"
	monitor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/flow/monitor"
	record "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/flow/record"
	bdi "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/bdi"
	ethernet "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/ethernet"
	isis "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/isis"
	loopback "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/loopback"
	mpls "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/mpls"
	nve "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/nve"
	ospf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/ospf"
	ospfv3 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/ospfv3"
	pim "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/pim"
	pimipv6 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/pimipv6"
	portchannel "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/portchannel"
	portchannelsubinterface "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/portchannelsubinterface"
	stackwisevirtual "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/stackwisevirtual"
	switchport "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/switchport"
	tunnel "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/tunnel"
	vlan "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/vlan"
	vrrpv2 "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/interface/vrrpv2"
	isisisis "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/isis/isis"
	mplsmpls "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/mpls/mpls"
	msdp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/multicast/msdp"
	multicast "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/multicast/multicast"
	pimmulticast "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/multicast/pim"
	pimipv6multicast "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/multicast/pimipv6"
	nat "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/nat/nat"
	ospfospf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/ospf/ospf"
	ospfv3ospf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/ospf/ospfv3"
	ospfv3addressfamilyipv4vrf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/ospf/ospfv3addressfamilyipv4vrf"
	ospfv3addressfamilyipv6vrf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/ospf/ospfv3addressfamilyipv6vrf"
	vrfospf "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/ospf/vrf"
	providerconfig "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/providerconfig"
	classmap "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/qos/classmap"
	policymap "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/qos/policymap"
	policymapevent "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/qos/policymapevent"
	qos "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/qos/qos"
	arp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/arp"
	aspathaccesslist "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/aspathaccesslist"
	communitylistexpanded "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/communitylistexpanded"
	communityliststandard "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/communityliststandard"
	ipv6localpool "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/ipv6localpool"
	ipv6prefixlist "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/ipv6prefixlist"
	largecommunitylistexpanded "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/largecommunitylistexpanded"
	prefixlist "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/prefixlist"
	routemap "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/routemap"
	staticroute "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/staticroute"
	vrfrouting "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/vrf"
	vrfstaticroutes "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/routing/vrfstaticroutes"
	cts "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/cts"
	devicesensor "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/devicesensor"
	devicetracking "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/devicetracking"
	dot1x "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/dot1x"
	keychain "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/keychain"
	parametermap "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/parametermap"
	servicetemplate "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/servicetemplate"
	zone "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/zone"
	zonepair "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/security/zonepair"
	bridgedomain "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/bridgedomain"
	errdisable "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/errdisable"
	spanningtree "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/spanningtree"
	stackwisevirtualswitching "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/stackwisevirtual"
	switchprovision "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/switchprovision"
	udld "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/udld"
	vlanswitching "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/vlan"
	vlanaccessmap "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/vlanaccessmap"
	vlanconfiguration "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/vlanconfiguration"
	vlanfilter "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/vlanfilter"
	vlangroup "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/vlangroup"
	vtp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/switching/vtp"
	banner "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/banner"
	cdp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/cdp"
	clock "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/clock"
	dhcp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/dhcp"
	eem "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/eem"
	license "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/license"
	line "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/line"
	lldp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/lldp"
	logging "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/logging"
	mdtsubscription "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/mdtsubscription"
	monitorsession "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/monitorsession"
	ntp "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/ntp"
	platform "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/platform"
	service "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/service"
	sla "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/sla"
	snmpserver "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/snmpserver"
	system "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/system"
	template "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/system/template"
	object "github.com/upbound/provider-upjet-iosxe/internal/controller/namespaced/yang/object"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aaa.Setup,
		accounting.Setup,
		authentication.Setup,
		authorization.Setup,
		radius.Setup,
		radiusserver.Setup,
		tacacs.Setup,
		tacacsserver.Setup,
		username.Setup,
		extended.Setup,
		ipv6.Setup,
		objectgroup.Setup,
		rolebased.Setup,
		standard.Setup,
		bfd.Setup,
		templatemultihop.Setup,
		templatesinglehop.Setup,
		addressfamilyipv4.Setup,
		addressfamilyipv4mvpn.Setup,
		addressfamilyipv4vrf.Setup,
		addressfamilyipv6.Setup,
		addressfamilyipv6vrf.Setup,
		addressfamilyl2vpn.Setup,
		addressfamilyvpnv4.Setup,
		addressfamilyvpnv6.Setup,
		bgp.Setup,
		bmpserver.Setup,
		ipv4mvpnneighbor.Setup,
		ipv4unicastneighbor.Setup,
		ipv4unicastvrfneighbor.Setup,
		ipv6unicastneighbor.Setup,
		l2vpnevpnneighbor.Setup,
		neighbor.Setup,
		peerpolicytemplate.Setup,
		peersessiontemplate.Setup,
		crypto.Setup,
		ikev2.Setup,
		ikev2keyring.Setup,
		ikev2policy.Setup,
		ikev2profile.Setup,
		ikev2proposal.Setup,
		ipsecprofile.Setup,
		ipsectransformset.Setup,
		pki.Setup,
		eigrp.Setup,
		vrf.Setup,
		ethernetsegment.Setup,
		evpn.Setup,
		instance.Setup,
		l2vfi.Setup,
		exporter.Setup,
		monitor.Setup,
		record.Setup,
		bdi.Setup,
		ethernet.Setup,
		isis.Setup,
		loopback.Setup,
		mpls.Setup,
		nve.Setup,
		ospf.Setup,
		ospfv3.Setup,
		pim.Setup,
		pimipv6.Setup,
		portchannel.Setup,
		portchannelsubinterface.Setup,
		stackwisevirtual.Setup,
		switchport.Setup,
		tunnel.Setup,
		vlan.Setup,
		vrrpv2.Setup,
		isisisis.Setup,
		mplsmpls.Setup,
		msdp.Setup,
		multicast.Setup,
		pimmulticast.Setup,
		pimipv6multicast.Setup,
		nat.Setup,
		ospfospf.Setup,
		ospfv3ospf.Setup,
		ospfv3addressfamilyipv4vrf.Setup,
		ospfv3addressfamilyipv6vrf.Setup,
		vrfospf.Setup,
		providerconfig.Setup,
		classmap.Setup,
		policymap.Setup,
		policymapevent.Setup,
		qos.Setup,
		arp.Setup,
		aspathaccesslist.Setup,
		communitylistexpanded.Setup,
		communityliststandard.Setup,
		ipv6localpool.Setup,
		ipv6prefixlist.Setup,
		largecommunitylistexpanded.Setup,
		prefixlist.Setup,
		routemap.Setup,
		staticroute.Setup,
		vrfrouting.Setup,
		vrfstaticroutes.Setup,
		cts.Setup,
		devicesensor.Setup,
		devicetracking.Setup,
		dot1x.Setup,
		keychain.Setup,
		parametermap.Setup,
		servicetemplate.Setup,
		zone.Setup,
		zonepair.Setup,
		bridgedomain.Setup,
		errdisable.Setup,
		spanningtree.Setup,
		stackwisevirtualswitching.Setup,
		switchprovision.Setup,
		udld.Setup,
		vlanswitching.Setup,
		vlanaccessmap.Setup,
		vlanconfiguration.Setup,
		vlanfilter.Setup,
		vlangroup.Setup,
		vtp.Setup,
		banner.Setup,
		cdp.Setup,
		clock.Setup,
		dhcp.Setup,
		eem.Setup,
		license.Setup,
		line.Setup,
		lldp.Setup,
		logging.Setup,
		mdtsubscription.Setup,
		monitorsession.Setup,
		ntp.Setup,
		platform.Setup,
		service.Setup,
		sla.Setup,
		snmpserver.Setup,
		system.Setup,
		template.Setup,
		object.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aaa.SetupGated,
		accounting.SetupGated,
		authentication.SetupGated,
		authorization.SetupGated,
		radius.SetupGated,
		radiusserver.SetupGated,
		tacacs.SetupGated,
		tacacsserver.SetupGated,
		username.SetupGated,
		extended.SetupGated,
		ipv6.SetupGated,
		objectgroup.SetupGated,
		rolebased.SetupGated,
		standard.SetupGated,
		bfd.SetupGated,
		templatemultihop.SetupGated,
		templatesinglehop.SetupGated,
		addressfamilyipv4.SetupGated,
		addressfamilyipv4mvpn.SetupGated,
		addressfamilyipv4vrf.SetupGated,
		addressfamilyipv6.SetupGated,
		addressfamilyipv6vrf.SetupGated,
		addressfamilyl2vpn.SetupGated,
		addressfamilyvpnv4.SetupGated,
		addressfamilyvpnv6.SetupGated,
		bgp.SetupGated,
		bmpserver.SetupGated,
		ipv4mvpnneighbor.SetupGated,
		ipv4unicastneighbor.SetupGated,
		ipv4unicastvrfneighbor.SetupGated,
		ipv6unicastneighbor.SetupGated,
		l2vpnevpnneighbor.SetupGated,
		neighbor.SetupGated,
		peerpolicytemplate.SetupGated,
		peersessiontemplate.SetupGated,
		crypto.SetupGated,
		ikev2.SetupGated,
		ikev2keyring.SetupGated,
		ikev2policy.SetupGated,
		ikev2profile.SetupGated,
		ikev2proposal.SetupGated,
		ipsecprofile.SetupGated,
		ipsectransformset.SetupGated,
		pki.SetupGated,
		eigrp.SetupGated,
		vrf.SetupGated,
		ethernetsegment.SetupGated,
		evpn.SetupGated,
		instance.SetupGated,
		l2vfi.SetupGated,
		exporter.SetupGated,
		monitor.SetupGated,
		record.SetupGated,
		bdi.SetupGated,
		ethernet.SetupGated,
		isis.SetupGated,
		loopback.SetupGated,
		mpls.SetupGated,
		nve.SetupGated,
		ospf.SetupGated,
		ospfv3.SetupGated,
		pim.SetupGated,
		pimipv6.SetupGated,
		portchannel.SetupGated,
		portchannelsubinterface.SetupGated,
		stackwisevirtual.SetupGated,
		switchport.SetupGated,
		tunnel.SetupGated,
		vlan.SetupGated,
		vrrpv2.SetupGated,
		isisisis.SetupGated,
		mplsmpls.SetupGated,
		msdp.SetupGated,
		multicast.SetupGated,
		pimmulticast.SetupGated,
		pimipv6multicast.SetupGated,
		nat.SetupGated,
		ospfospf.SetupGated,
		ospfv3ospf.SetupGated,
		ospfv3addressfamilyipv4vrf.SetupGated,
		ospfv3addressfamilyipv6vrf.SetupGated,
		vrfospf.SetupGated,
		providerconfig.SetupGated,
		classmap.SetupGated,
		policymap.SetupGated,
		policymapevent.SetupGated,
		qos.SetupGated,
		arp.SetupGated,
		aspathaccesslist.SetupGated,
		communitylistexpanded.SetupGated,
		communityliststandard.SetupGated,
		ipv6localpool.SetupGated,
		ipv6prefixlist.SetupGated,
		largecommunitylistexpanded.SetupGated,
		prefixlist.SetupGated,
		routemap.SetupGated,
		staticroute.SetupGated,
		vrfrouting.SetupGated,
		vrfstaticroutes.SetupGated,
		cts.SetupGated,
		devicesensor.SetupGated,
		devicetracking.SetupGated,
		dot1x.SetupGated,
		keychain.SetupGated,
		parametermap.SetupGated,
		servicetemplate.SetupGated,
		zone.SetupGated,
		zonepair.SetupGated,
		bridgedomain.SetupGated,
		errdisable.SetupGated,
		spanningtree.SetupGated,
		stackwisevirtualswitching.SetupGated,
		switchprovision.SetupGated,
		udld.SetupGated,
		vlanswitching.SetupGated,
		vlanaccessmap.SetupGated,
		vlanconfiguration.SetupGated,
		vlanfilter.SetupGated,
		vlangroup.SetupGated,
		vtp.SetupGated,
		banner.SetupGated,
		cdp.SetupGated,
		clock.SetupGated,
		dhcp.SetupGated,
		eem.SetupGated,
		license.SetupGated,
		line.SetupGated,
		lldp.SetupGated,
		logging.SetupGated,
		mdtsubscription.SetupGated,
		monitorsession.SetupGated,
		ntp.SetupGated,
		platform.SetupGated,
		service.SetupGated,
		sla.SetupGated,
		snmpserver.SetupGated,
		system.SetupGated,
		template.SetupGated,
		object.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		aaa.SetupWebhookWithManager,
		accounting.SetupWebhookWithManager,
		authentication.SetupWebhookWithManager,
		authorization.SetupWebhookWithManager,
		radius.SetupWebhookWithManager,
		radiusserver.SetupWebhookWithManager,
		tacacs.SetupWebhookWithManager,
		tacacsserver.SetupWebhookWithManager,
		username.SetupWebhookWithManager,
		extended.SetupWebhookWithManager,
		ipv6.SetupWebhookWithManager,
		objectgroup.SetupWebhookWithManager,
		rolebased.SetupWebhookWithManager,
		standard.SetupWebhookWithManager,
		bfd.SetupWebhookWithManager,
		templatemultihop.SetupWebhookWithManager,
		templatesinglehop.SetupWebhookWithManager,
		addressfamilyipv4.SetupWebhookWithManager,
		addressfamilyipv4mvpn.SetupWebhookWithManager,
		addressfamilyipv4vrf.SetupWebhookWithManager,
		addressfamilyipv6.SetupWebhookWithManager,
		addressfamilyipv6vrf.SetupWebhookWithManager,
		addressfamilyl2vpn.SetupWebhookWithManager,
		addressfamilyvpnv4.SetupWebhookWithManager,
		addressfamilyvpnv6.SetupWebhookWithManager,
		bgp.SetupWebhookWithManager,
		bmpserver.SetupWebhookWithManager,
		ipv4mvpnneighbor.SetupWebhookWithManager,
		ipv4unicastneighbor.SetupWebhookWithManager,
		ipv4unicastvrfneighbor.SetupWebhookWithManager,
		ipv6unicastneighbor.SetupWebhookWithManager,
		l2vpnevpnneighbor.SetupWebhookWithManager,
		neighbor.SetupWebhookWithManager,
		peerpolicytemplate.SetupWebhookWithManager,
		peersessiontemplate.SetupWebhookWithManager,
		crypto.SetupWebhookWithManager,
		ikev2.SetupWebhookWithManager,
		ikev2keyring.SetupWebhookWithManager,
		ikev2policy.SetupWebhookWithManager,
		ikev2profile.SetupWebhookWithManager,
		ikev2proposal.SetupWebhookWithManager,
		ipsecprofile.SetupWebhookWithManager,
		ipsectransformset.SetupWebhookWithManager,
		pki.SetupWebhookWithManager,
		eigrp.SetupWebhookWithManager,
		vrf.SetupWebhookWithManager,
		ethernetsegment.SetupWebhookWithManager,
		evpn.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		l2vfi.SetupWebhookWithManager,
		exporter.SetupWebhookWithManager,
		monitor.SetupWebhookWithManager,
		record.SetupWebhookWithManager,
		bdi.SetupWebhookWithManager,
		ethernet.SetupWebhookWithManager,
		isis.SetupWebhookWithManager,
		loopback.SetupWebhookWithManager,
		mpls.SetupWebhookWithManager,
		nve.SetupWebhookWithManager,
		ospf.SetupWebhookWithManager,
		ospfv3.SetupWebhookWithManager,
		pim.SetupWebhookWithManager,
		pimipv6.SetupWebhookWithManager,
		portchannel.SetupWebhookWithManager,
		portchannelsubinterface.SetupWebhookWithManager,
		stackwisevirtual.SetupWebhookWithManager,
		switchport.SetupWebhookWithManager,
		tunnel.SetupWebhookWithManager,
		vlan.SetupWebhookWithManager,
		vrrpv2.SetupWebhookWithManager,
		isisisis.SetupWebhookWithManager,
		mplsmpls.SetupWebhookWithManager,
		msdp.SetupWebhookWithManager,
		multicast.SetupWebhookWithManager,
		pimmulticast.SetupWebhookWithManager,
		pimipv6multicast.SetupWebhookWithManager,
		nat.SetupWebhookWithManager,
		ospfospf.SetupWebhookWithManager,
		ospfv3ospf.SetupWebhookWithManager,
		ospfv3addressfamilyipv4vrf.SetupWebhookWithManager,
		ospfv3addressfamilyipv6vrf.SetupWebhookWithManager,
		vrfospf.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		classmap.SetupWebhookWithManager,
		policymap.SetupWebhookWithManager,
		policymapevent.SetupWebhookWithManager,
		qos.SetupWebhookWithManager,
		arp.SetupWebhookWithManager,
		aspathaccesslist.SetupWebhookWithManager,
		communitylistexpanded.SetupWebhookWithManager,
		communityliststandard.SetupWebhookWithManager,
		ipv6localpool.SetupWebhookWithManager,
		ipv6prefixlist.SetupWebhookWithManager,
		largecommunitylistexpanded.SetupWebhookWithManager,
		prefixlist.SetupWebhookWithManager,
		routemap.SetupWebhookWithManager,
		staticroute.SetupWebhookWithManager,
		vrfrouting.SetupWebhookWithManager,
		vrfstaticroutes.SetupWebhookWithManager,
		cts.SetupWebhookWithManager,
		devicesensor.SetupWebhookWithManager,
		devicetracking.SetupWebhookWithManager,
		dot1x.SetupWebhookWithManager,
		keychain.SetupWebhookWithManager,
		parametermap.SetupWebhookWithManager,
		servicetemplate.SetupWebhookWithManager,
		zone.SetupWebhookWithManager,
		zonepair.SetupWebhookWithManager,
		bridgedomain.SetupWebhookWithManager,
		errdisable.SetupWebhookWithManager,
		spanningtree.SetupWebhookWithManager,
		stackwisevirtualswitching.SetupWebhookWithManager,
		switchprovision.SetupWebhookWithManager,
		udld.SetupWebhookWithManager,
		vlanswitching.SetupWebhookWithManager,
		vlanaccessmap.SetupWebhookWithManager,
		vlanconfiguration.SetupWebhookWithManager,
		vlanfilter.SetupWebhookWithManager,
		vlangroup.SetupWebhookWithManager,
		vtp.SetupWebhookWithManager,
		banner.SetupWebhookWithManager,
		cdp.SetupWebhookWithManager,
		clock.SetupWebhookWithManager,
		dhcp.SetupWebhookWithManager,
		eem.SetupWebhookWithManager,
		license.SetupWebhookWithManager,
		line.SetupWebhookWithManager,
		lldp.SetupWebhookWithManager,
		logging.SetupWebhookWithManager,
		mdtsubscription.SetupWebhookWithManager,
		monitorsession.SetupWebhookWithManager,
		ntp.SetupWebhookWithManager,
		platform.SetupWebhookWithManager,
		service.SetupWebhookWithManager,
		sla.SetupWebhookWithManager,
		snmpserver.SetupWebhookWithManager,
		system.SetupWebhookWithManager,
		template.SetupWebhookWithManager,
		object.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
