# Testing against an IOS-XE device

This provider drives real NETCONF against a real IOS-XE control plane, so
testing anything beyond the unit tests needs a device. This page collects the
options.

> [!NOTE]
> The reachability checks below were made on 2026-08-25. Cisco changes sandbox
> hostnames, credentials and access rules from time to time, so re-verify
> before relying on them.

## There is no IOS-XE container

Cisco ships a containerized network OS only for IOS-XR (XRd). IOS-XE is a
VM-only image. The closest container-like thing is IOL (IOS On Linux), which
runs inside Cisco Modeling Labs and whose image is licensed only for use inside
CML. So the choice is between a Cisco-hosted device and an x86 VM you run
yourself.

## Options at a glance

| Option | Cost | Good for | Catch |
| --- | --- | --- | --- |
| DevNet always-on sandbox | free | smoke-testing a few managed resources | shared, config gets reverted, rate-limited |
| DevNet reservable sandbox | free | real e2e/uptest runs | reserve hours or days ahead |
| Catalyst 8000V / 9000V you host | license or cloud pay-as-you-go | full control, matches upstream CI | x86 only |
| CML-Free | free | topologies, IOL nodes | no C8000V/C9000V in the free refplat |
| netopeer2 + Cisco YANG models | free | CI plumbing regression | no IOS semantics |

## DevNet always-on sandbox

The quickest way to exercise the provider end to end. NETCONF was reachable on
all three hosts:

```text
devnetsandboxiosxec8k.cisco.com:830   open
sandbox-iosxe-latest-1.cisco.com:830  open
sandbox-iosxe-recomm-1.cisco.com:830  open
```

The classic `sandbox-iosxe-latest-1` / `sandbox-iosxe-recomm-1` pair has
historically used `admin` / `C1sco12345`. The newer Catalyst 8000 lab
(`devnetsandboxiosxec8k`) issues per-launch credentials, obtained by launching
its tile in the [DevNet sandbox portal](https://devnetsandbox.cisco.com), so
fetch those rather than assuming a static password.

Sessions give privilege 15 read-write access with NETCONF and RESTCONF enabled.
Do not modify the management interface or AAA: monitoring reverts the device to
a base configuration if connectivity breaks.

Suitable for a handful of resources. Not suitable for a full uptest sweep: the
device is shared, and the singleton resources (`System`, `AAA`, `Logging`, ...)
rewrite global configuration.

### Running the provider against it

The `ProviderConfig` credentials are just the host and login:

```json
{"username":"admin","password":"C1sco12345","host":"sandbox-iosxe-latest-1.cisco.com:830","insecure":true}
```

```console
# in one shell, against a cluster with the CRDs applied
kubectl apply -f package/crds
PROVIDER_MAX_RECONCILE_RATE=1 make run

# in another
kubectl apply -f examples/cluster/providerconfig/   # after filling in secret.yaml.tmpl
kubectl apply -f examples/cluster/routing/vrf.yaml
kubectl get vrf.routing.iosxe.upbound.io -o wide
```

`make run` passes only `--debug`, so other options go through their environment
variables (run the built binary directly for full control).

Keep the reconcile rate at 1 against a shared device. Upjet configures the
embedded provider once per reconciliation, so the Terraform provider's own
serialization of NETCONF operations does not span concurrent reconciliations.

## Hosting Catalyst 8000V or Catalyst 9000V

This is what the upstream Terraform provider tests against: its `GNUmakefile`
runs acceptance tests against `IOSXE_1715_ROUTER_HOST` (a C8000V) and
`IOSXE_1715_SWITCH_HOST` (a C9000V) on IOS-XE 17.15.x with
`IOSXE_PROTOCOL=netconf`. Matching that gives the closest thing to upstream
coverage.

Both images are x86 only. On an Apple Silicon machine that means emulating x86
under QEMU or UTM, which boots but is painfully slow. Better:

- a Catalyst 8000V from the AWS or Azure marketplace, where pay-as-you-go
  sidesteps licensing entirely;
- an x86 Linux box running KVM, or an ESXi host.

Then enable `netconf-yang` and add a privilege 15 local user:

```text
conf t
 netconf-yang
 username crossplane privilege 15 secret <password>
end
write memory
```

## Cisco Modeling Labs

CML-Free is a no-cost single-user tier: up to 5 nodes, no license or Smart
Licensing account, does not expire. The free reference platform ISO ships IOL,
IOL-L2, ASAv and host images only — Catalyst 8000V and 9000V are not included.
IOL is real IOS-XE code, a trimmed build missing some commands, so basic
`netconf-yang` may work, but this is unverified. CML itself runs on x86.

Note the licensing restriction: the VM images shipped with CML or its refplat
ISO are licensed only for use inside CML.

## A fake NETCONF server for CI

If CI needs something with no device behind it, run a `netopeer2` / `sysrepo`
container with the Cisco IOS-XE YANG models from
[YangModels/yang](https://github.com/YangModels/yang) loaded. That exercises the
transport, the path construction and the `edit-config` payloads the provider
generates.

What it does not give you: IOS semantics, defaults, or validation. Loading
`Cisco-IOS-XE-native` and its dependency tree is also heavy. Treat it as a
regression net for provider wiring, not as a substitute for a device.

## Wiring a device into uptest

[cluster/test/setup.sh](../cluster/test/setup.sh) creates the credentials secret
from `UPTEST_CLOUD_CREDENTIALS` — for this provider, the JSON blob above — and
builds both the cluster scoped and the namespaced `ProviderConfig`.

```console
export UPTEST_CLOUD_CREDENTIALS='{"username":"admin","password":"...","host":"...:830","insecure":true}'
export UPTEST_EXAMPLE_LIST="examples/cluster/routing/vrf.yaml"
make e2e
```

`.github/workflows/e2e.yaml` triggers on an issue comment and reads the same
value from the `UPTEST_CLOUD_CREDENTIALS` repository secret, so that secret is
what should eventually point at a reserved sandbox or a self-hosted C8000V.
Start with a short `UPTEST_EXAMPLE_LIST`: a full 138-resource sweep against a
single device serializes badly.

## Open questions to return to

- Which device backs CI: a reserved DevNet sandbox, or a self-hosted
  Catalyst 8000V in a cloud account.
- Whether IOL under CML-Free supports `netconf-yang` well enough to be useful
  for local development.
- Which examples belong in the e2e list, and which singleton resources must be
  excluded from it because they rewrite global device configuration.

## Sources

- [New Always-On DevNet Sandbox for Cisco Catalyst 8000 & Catalyst 9000](https://community.cisco.com/t5/devnet-general-blogs/new-always-on-devnet-sandbox-for-cisco-catalyst-8000-amp/ba-p/5330526)
- [DevNet standard network devices](https://developer.cisco.com/site/standard-network-devices/)
- [Cisco Modeling Labs - Free](https://developer.cisco.com/docs/modeling-labs/cml-free/)
- [VM images for CML labs](https://developer.cisco.com/docs/modeling-labs/vm-images-for-cml-labs/)
- [Catalyst 8000V KVM installation guide](https://www.cisco.com/c/en/us/td/docs/routers/C8000V/Configuration/c8000v-installation-configuration-guide/m_install_cisco_kvm_environments.html)
- [IOS XE 17.15 NETCONF protocol guide](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1715/b_1715_programmability_cg/m_1715_prog_yang_netconf.html)
