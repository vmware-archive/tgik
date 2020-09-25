# Episode 133 : Hierarchical Namespaces
- Hosted by @joshrosso
- 07/25/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=zAtaJ0x-ZwY 
" target="_blank"><img src="http://img.youtube.com/vi/zAtaJ0x-ZwY/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa)
- [00:04:37 - Week in Review](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=277)
- [00:15:11 - Hierarchical Namespace Overview](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=911)
- [00:30:05 - Installing HNC and the Plugin](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1805)
- [00:40:53 - Creating Namespaces](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=2453)
- [00:54:38 - Synchronizing Objects](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=3288)
- [01:15:32 - Experiments with NS Ownership and more!](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=4532)
- [01:31:53 - Advanced Resource Quota Propogation Ideas](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=5513)
- [01:37:05 - Wrap Up](https://youtu.be/zAtaJ0x-ZwY?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=5825)

## Week in Review

### Core K8s

* K8s 1.20.0 [call for enhancements](https://groups.google.com/g/kubernetes-dev/c/J9KWovW9XM0)
  * Tuesday, Oct 6th: Week 4 - Enhancements Freeze
  * Thursday, Nov 12th: Week 9 - Code Freeze
  * Don't forget to [Vote](https://groups.google.com/g/kubernetes-dev/c/acchxwYbjQg) on the [Steering Committee](https://github.com/kubernetes/steering/blob/master/charter.md) Election.
    * [Eligibility](https://github.com/kubernetes/community/tree/master/events/elections/2020#eligibility)

### K8s and Cloud Native Ecosystem

* [Open Service Mesh (OSM) accepted into CNCF as a sandbox project](https://openservicemesh.io/blog/open-service-mesh-osm-accepted-into-cncf-as-a-sandbox-project)
* [Helm 3.3.2](https://github.com/helm/helm/releases/tag/v3.3.2) [security](https://github.com/helm/helm/security/advisories) (patch) release
  * Users are strongly recommended to update to this release. Most of the issues were discovered by Trail of Bits during their CNCF-sponsored audit of the Helm codebase. We are grateful for Trail of Bits' detailed and thorough analysis of the Helm codebase. In addition, a Helm core maintainer identified one more issue.
* KubeEdge accepted as [incubating](https://www.cncf.io/blog/2020/09/16/toc-approves-kubeedge-as-incubating-project/)
* Stakrox "[Kubernetes and Container Security and Adoption Trends](https://www.stackrox.com/kubernetes-adoption-security-and-market-share-for-containers/)" report - some infographics suggest everyone uses kubernetes now.
* [CF summit schedule](https://www.cloudfoundry.org/events/summit/europe-2020/schedule) announced, article highlights there is lots of K8s content happening.
* [The Forrester Wave™: Multicloud Container Development Platforms, Q3 2020](https://info.rancher.com/forrester-new-wave-enterprise-container-platform)

## Show Notes

- [Kubernetes Blog: Introducing Hierarchical Namespaces](https://kubernetes.io/blog/2020/08/14/introducing-hierarchical-namespaces)
- [GitHub: HN Controller](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/hnc)
