# Episode 146 : Crossplane
- Hosted by @evankanderson
- 02/26/2021

<a href="https://www.youtube.com/watch?v=slX2nAFHeK0
" target="_blank"><img src="http://img.youtube.com/vi/slX2nAFHeK0/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:25 - Welcome to TGIK!
- 00:03:36 - Week in Review
- 00:11:11 - Welcome to Crossplane
- 00:13:40 - Brief OAM digression
- 00:18:50 - Crossplane install
- 00:24:20 - Install Crossplane kubectl extension
- 00:29:50 - Exploring Crossplane RBAC-manager
- 00:34:45 - Installing the AWS provider
- 00:43:30 - Diving into Crossplane packaging
- 00:46:00 - Mention of [dependency confusion attacks](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610)
- 00:49:25 - Using [dive](https://github.com/wagoodman/dive) to view package contents
- 00:50:20 - Creating a PostgreSQL instance
- 00:56:20 - Diving in to how a PostgreSQL db backed by RDS is represented in Crossplane
- 01:02:40 - Resource parameters, patches and resource sizing
- 01:09:00 - Trying to figure out how to say "database with 8GB RAM" at different cloud providers
- 01:13:00 - Investigating patch options
- 01:16:10 - Got the RDS instance running!
- 01:19:20 - Crossplane resource category list
- 01:23:00 - Discussion about cluster-level Crossplane Providers and Compositions (vs namespace level)
- 01:25:10 - Dependencies / building your own Compositions
- 

## Week in Review

### Core K8s

- Kubernetes 1.21 Code Freeze is [March 9](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.21#timeline).

### K8s and Cloud Native Ecosystem

- [Klustered](https://www.youtube.com/playlist?list=PLz0t90fOInA5IyhoT96WhycPV8Km-WICj): Watch Kubernetes experts attempt to fix Kubernetes clusters, broken by members of the community.
- [Contour 1.13](https://projectcontour.io/contour_v1130/): Big new features include new Gateway APIs and global rate limiting.
- Knative 0.21 release: [Serving](https://github.com/knative/serving/releases), [Eventing](https://github.com/knative/eventing/releases). Adds DomainMapping, event delivery fixes.
- [Harbor 2.2 release](https://github.com/goharbor/harbor/releases/tag/v2.2.0): Metrics, auth, proxy cache improvements.
- [Network Policy Editor](https://cilium.io/blog/2021/02/10/network-policy-editor): Awesome editor for creating and visualizing network policies from the cilium folks (even if you aren't using Cilium).
- [Nice intro to Flux v2](https://blog.sldk.de/2021/02/introduction-to-gitops-on-kubernetes-with-flux-v2/)
- [DAPR release 1.0](https://github.com/dapr/dapr/releases)
- [Security Groups for pods walkthrough](https://swade1987.medium.com/eks-security-groups-for-pods-38bd0fed14a6)
- [Rate Limiting in controller-runtime and client-go](https://danielmangum.com/posts/controller-runtime-client-go-rate-limiting/)
- [New CNCF radar](https://radar.cncf.io/)
- [How to deep triage a sig-network bug thanks to aojea for showing me this](https://jayunit100.blogspot.com/2021/02/how-to-triage-sig-network-bug.html)

## Show Notes

