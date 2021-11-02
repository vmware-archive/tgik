# Episode 158 : Cluster API Tinkerbell
- Hosted by [@randomvariable] with guest [@detiber]
    - 2021/06/25


<a href="https://www.youtube.com/watch?v=Di_AR6nAss0
" target="_blank"><img src="http://img.youtube.com/vi/Di_AR6nAss0/hqdefault.jpg" width="480" height="360" border="10" /></a>

Photo by <a href="https://unsplash.com/@callumlwale?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Callum Wale</a> on <a href="https://unsplash.com/s/photos/sheffield-metal?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>


## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core K8s
- [Kubernetes v1.22 beta]
  - [Deprecated APIs are being removed, take note!][deprecated-apis]
      - `admissionregistration.k8s.io/v1beta1` & `apiextensions.k8s.io/v1beta1` amongst others. Get your manifests and controllers up to date.
- [Interactive mode for exec credential providers][credential-providers]


### K8s and Cloud Native Ecosystem
- New CNCF Sandbox projects
  - [Kubevela]
  - [kube-vip]
  - [kubedl]
  - [meshery] & [service mesh performance]

## Show Notes

### Previously in TGIK...

We've covered Cluster API in previous episodes, so you can take a look at the previous episodes for an intro (but don't go away from the livestream!)

* [Episode 143: Cluster API Update]
* [Episode 110: Cluster API v1alpha3]
* [Episode 108: Cluster API Docker]
* [Episode 53: Cluster API AWS]

#### [Cluster API v0.4.0 v1alpha4 Release][cluster-api-v0.4.0]

### Tinkerbell

https://tinkerbell.org/

#### The "Lab" Environment
* The cupboard, the vSphere environment, the router.
* The [sandbox] repository
    * Stands up Tinkerbell prerequisites
        * PostgreSQL database
        * OCI Compliant Registry
        * Web Server
    * Stands up Tinkerbell components
* detiber's homelab setup: https://gist.github.com/detiber/4dbe8473b422a534c8a28120e3729120

#### Components

* [Boots]
    * Serves DHCP networks
    * TFTP Server
    * Bootstraps [iPXE] environment on hosts configured for PXE booting

* [Tink]
    * Server
        * Hardware
        * Templates
            * Comprised of actions that define what steps to perform
                * [Action Hub]
        * Workflows
            * Link a template to hardware
    * Worker
    * CLI

* [Hook]
    * In memory minimal OS for running Tink Worker

* [Hegel]
    * Metadata Service

* [PBnJ]
    * BMC management
    * Not currently used by cluster-api-provider-tinkerbell
    * Could use with vSphere using [vbmc4vsphere]

* Cluster API Pieces
    * [Cluster API Tinkerbell]
    * [Tilt]
        * [See Ellen Korbes' fun little video][tilting-at-windmills]

#### Further resources

<a href="https://www.youtube.com/watch?v=QxpKnMGywTU
" target="_blank">Introduction to Tinkerbell</a>

<!-- Links -->

[Naadir Jeewa]: https://kubernetes.slack.com/archives/D6SJC6KK9
[@detiber]: https://twitter.com/detiber
[credential-providers]: https://github.com/kubernetes/kubernetes/pull/99310
[Kubevela]: https://kubevela.io/
[kube-vip]: https://kube-vip.io/
[deprecated-apis]: https://kubernetes.io/docs/reference/using-api/deprecation-guide/#v1-22
[kubedl]: https://github.com/alibaba/kubedl
[service mesh performance]: https://github.com/layer5io/service-mesh-performance
[meshery]: https://github.com/layer5io/meshery
[Boots]: https://github.com/tinkerbell/boots
[Hegel]: https://github.com/tinkerbell/hegel
[Hook]: https://github.com/tinkerbell/hook
[iPXE]: https://ipxe.org/
[Tink]: https://github.com/tinkerbell/tink
[PBnJ]: https://github.com/tinkerbell/pbnj
[Action Hub]: https://artifacthub.io/packages/search?kind=4
[cluster-api-v0.4.0]: https://github.com/kubernetes-sigs/cluster-api/releases/tag/v0.4.0
[tilting-at-windmills]: https://twitter.com/ellenkorbes/status/1400139022580826117
[Episode 108: Cluster API Docker]: https://www.youtube.com/watch?v=6pFW6h6AORQ
[Episode 143: Cluster API Update]: https://www.youtube.com/watch?v=AeHfVFepsMg
[Episode 110: Cluster API v1alpha3]: https://www.youtube.com/watch?v=A8dUFWbH3pM
[Episode 53: Cluster API AWS]: https://www.youtube.com/watch?v=e1XCsuTYUa4
[Kubernetes v1.22 beta]: https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.22.md
[sandbox]: https://github.com/tinkerbell/sandbox
[vbmc4vsphere]: https://pypi.org/project/vbmc4vsphere/
[Cluster API Tinkerbell]: https://github.com/tinkerbell/cluster-api-provider-tink
[tilt]: https://twitter.com/ellenkorbes/status/1400139022580826117
