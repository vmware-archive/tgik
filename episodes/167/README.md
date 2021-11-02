# Episode 167: Policy++, the bleeding edge of Pod, Network, and other emerging K8s APIs

https://unsplash.com/photos/MzKTG9cQ2zo
 
## Week in Review 

- TKG 1.4 is out 
- https://kubernetes.io/blog/2021/09/13/read-write-once-pod-access-mode-alpha/
- (not this week but... go 1.17 is out time to update)

## Show Notes!

Policies in general:
- AWS SecurityGroups
- NSX DFW
- GCE Security APIs?
- Selinux ? does it have networking primitives?
...

### Gke Networking

- calico/cillium
- pod networking on GCE ipam
- svc mesh [istio/gke]
- multicluster vs svc mesh

services:
- need grouping 
- ingress to monolithic -> new gateway APIs totally modular
- ExternalIPs are confusing, original primitive
    - loadbalancerIngree

### PSP.next

(liggitt or others if any interesting notes highlights feel free to add here !)
https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/

https://github.com/kubernetes/enhancements/issues/2579 

```
from @liggitt ... 
- biggest update is the plan for what to do with OS-specific fields

- plan is to let pods optionally identify their OS, and skip Pod Security checks that don't apply to that OS - details in https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2802-identify-windows-pods-apiserver-admission#changes-to-podsecurity-standards

- so a pod could say "I'm a windows pod" and not be required to add linux-specific fields in order to pass the restricted Pod Security standard profile check
```

### NetworkPolicy++

#### Port range

#### Namespace by name
    TODO: (jayunit) add doc links

#### CNP 
    google doc : https://docs.google.com/document/d/1NkEJFf07CdFY9X7FmMzE3RqZOWgWshRRR2xpnpKBUSM/edit

#### Netpol Status (help!)
    https://github.com/kubernetes/enhancements/pull/2947

### PSPs -> Adm controller KEP 

https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2579-psp-replacement 



https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2579-psp-replacement#windows-support

- 00:01:00 Intro
- 00:02:58 News: VMWare Tanzu 1.4 is out!
- 00:06:11 RWM RWO storage volumes to Single Pod Access
- 00:11:00 Tim Hockins (Google) surprise guest ! 
- 00:15:00 trying to patch my golang
- 00:18:00 NetworkPolicy WG what's up
- 00:19:30 quick GKE overview of enhanced networks
- 00:25:00 there's lots of problems with Services
- 00:30:00 "proto loadbalancer" ExternalIPs
- 00:32:00 svc type LoadBalancers evolution
- 00:35:00 EndPort fields
- 00:38:00 Are EndPorts still alpha?
- 00:41:00 Kind allows you to easily declare FeatureGates for Kubernetes!
- 00:55:00 bypassing namespace restrictions
- 00:56:00 Tenants and future ClusterScoped NetworkPolicies
- 01:00:00 Empowerment vs Priority based policies for the future
- 01:07:00 more on tenants and namespaces
- 01:13:00 delegating to lower level network policies
- 01:20:00 PSPs, OPA, and the future of PSPs
- 01:25:00 how does the Kubelet deal with PSps, how will it in the future?
