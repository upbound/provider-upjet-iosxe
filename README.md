# provider-upjet-iosxe

`provider-upjet-iosxe` is a [Crossplane](https://crossplane.io/) provider for
Cisco IOS-XE devices. It is built with [Upjet](https://github.com/crossplane/upjet)
from the [Cisco IOS-XE Terraform provider](https://github.com/CiscoDevNet/terraform-provider-iosxe)
and exposes IOS-XE configuration as Kubernetes managed resources.

The provider runs in Upjet's **no-fork** mode: the Terraform provider is linked
into the controller binary and called in-process through the
[Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework).
No Terraform CLI, no provider plugin process and no Terraform state files are
involved at runtime; the controller talks NETCONF to the devices through the
embedded provider.

Both API flavours of Crossplane v2 are generated:

| Scope | API groups | ProviderConfig |
| --- | --- | --- |
| Cluster scoped (legacy) | `<group>.iosxe.upbound.io` | `iosxe.upbound.io/v1beta1` |
| Namespaced | `<group>.iosxe.m.upbound.io` | `iosxe.m.upbound.io/v1beta1` |

## API groups

138 of the 141 Terraform resources are exposed as managed resources, grouped by
configuration domain:

| Group | Contents |
| --- | --- |
| `aaa` | AAA, accounting, authentication, authorization, RADIUS, TACACS, local users |
| `acl` | Standard, extended, IPv6 and role based access lists, object groups |
| `bfd` | BFD and BFD templates |
| `bgp` | BGP, address families, neighbors, peer templates, BMP servers |
| `crypto` | IKEv2, IPsec, PKI |
| `eigrp` | EIGRP, EIGRP VRFs |
| `evpn` | EVPN, ethernet segments, EVPN instances, L2 VFIs |
| `flow` | Flexible NetFlow exporters, monitors, records |
| `interface` | Ethernet, loopback, port channel, tunnel, VLAN and other interfaces |
| `isis` | IS-IS |
| `mpls` | MPLS |
| `multicast` | Multicast, PIM, MSDP |
| `nat` | NAT |
| `ospf` | OSPF, OSPFv3 and their VRF address families |
| `qos` | Class maps, policy maps, QoS |
| `routing` | Static routes, VRFs, route maps, prefix lists, community lists, ARP |
| `security` | 802.1X, CTS, device tracking, zone based firewall, key chains |
| `switching` | VLANs, spanning tree, VTP, UDLD, bridge domains, stacking |
| `system` | Hostname and system settings, logging, NTP, SNMP, LLDP, CDP, EEM, telemetry |
| `yang` | Generic YANG object, an escape hatch for unsupported configuration |

`iosxe_cli`, `iosxe_commit` and `iosxe_save_config` are deliberately not
exposed: they are imperative one-shot operations rather than declarative
configuration that can be observed and reconciled.

## Getting started

Install the provider:

```console
kubectl apply -f examples/install.yaml
```

Create a secret with the device credentials and a `ProviderConfig` that
references it:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "username": "admin",
      "password": "t0ps3cr3t11",
      "host": "10.0.0.1:830",
      "insecure": true
    }
---
apiVersion: iosxe.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
```

The supported credential keys are:

| Key | Description |
| --- | --- |
| `username`, `password` | Device credentials. Required. |
| `host` | Hostname or IP address, optionally `host:port`. The NETCONF default port is 830. |
| `insecure` | Skip SSH host key verification. Defaults to `true` in the Terraform provider. |
| `retries` | Number of retries for NETCONF calls. |
| `lockReleaseTimeout` | Seconds to wait for the device configuration lock to be released. |
| `devices` | List of `{"name", "host", "managed"}` objects to manage several devices with one `ProviderConfig`. |
| `selectedDevices` | Restricts reconciliation to a subset of `devices`. |

Then create managed resources:

```yaml
apiVersion: routing.iosxe.upbound.io/v1alpha1
kind: VRF
metadata:
  name: example
spec:
  forProvider:
    name: VRF22
    description: Managed by Crossplane
    rd: "22:22"
    addressFamilyIpv4: true
  providerConfigRef:
    name: default
```

`examples-generated/` contains a generated example for every resource, in both
the cluster scoped and the namespaced flavour.

### Selecting a device

Every managed resource has an optional `spec.forProvider.device` field that
selects a device by name from the `devices` list of the `ProviderConfig`.
Resources that leave it empty are reconciled against the device configured in
the top level `host`.

### External names and importing

Each IOS-XE resource manages a YANG object and is identified by the path of
that object, for example
`Cisco-IOS-XE-native:native/vrf/definition=VRF22`. That path is what the
Terraform provider reports as the resource id, so it is also this provider's
external name: it is set in the `crossplane.io/external-name` annotation after
creation, and setting it on a new managed resource adopts the existing device
configuration instead of creating it.

### Deletion behaviour

The `spec.forProvider.deleteMode` field of most resources controls what happens
on deletion: `all` (the default) removes the whole YANG object, while
`attributes` only removes the attributes that are explicitly configured and
leaves the rest of the object in place.

## Developing

### Terraform provider dependency

The IOS-XE Terraform provider keeps its Plugin Framework implementation in an
`internal/` package, which Go does not allow other modules to import. Running
it in-process therefore requires a fork of the Terraform provider that exports
it. The fork adds a single package:

```go
// xpprovider/xpprovider.go
package xpprovider

func GetProvider(version string) provider.Provider {
	return iosxeprovider.New(version)()
}
```

and is wired in through a `replace` directive in `go.mod`:

```text
replace github.com/CiscoDevNet/terraform-provider-iosxe => ../upbound/terraform-provider-iosxe
```

> [!IMPORTANT]
> The replace directive currently points at a local checkout. Before CI can
> build this repository, push the fork (for example to
> `github.com/upbound/terraform-provider-iosxe`) and change the directive to
> that module and a pinned pseudo-version.

### Code generation

```console
make submodules          # first time only
make generate
```

`make generate` fetches the Terraform provider schema with the Terraform CLI
(`config/schema.json`), scrapes the provider documentation into
`config/provider-metadata.yaml`, and runs the Upjet pipeline. Note that
Terraform is only used at generation time.

### Adding a resource

Add an entry to `resourceConfigs` in
[config/resources/resources.go](config/resources/resources.go) with the API
group and kind it should be generated into, then run `make generate`. That
table is also the include list of the provider, so a resource that is not
listed is not generated. External names do not need to be configured per
resource: every IOS-XE resource is identified by the YANG path the Terraform
provider computes, so `config.IdentifierFromProvider` applies to all of them.

Two constraints are worth remembering when picking a kind: kinds must be unique
within a group, and a kind must not be a Go keyword, because Upjet names the
generated controller package after it.

### Running locally

```console
make run
```

### Build

```console
make build      # binary and provider package
```

## Operational notes

- Upjet configures the embedded Terraform provider once per reconciliation, so
  NETCONF connection reuse is disabled and the provider opens and closes a
  session per operation. Automatic commit is enabled, which means every
  operation commits the configuration it changes.
- Because each reconciliation gets its own provider instance, the in-provider
  serialization of NETCONF operations does not span concurrent reconciliations.
  Devices that are sensitive to concurrent configuration sessions should be
  managed with a low `--max-reconcile-rate`.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/provider-upjet-iosxe/issues).
