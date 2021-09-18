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

