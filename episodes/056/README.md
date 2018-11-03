# Episode 056 : Heptio Contour and IngressRoute

- Hosted by @jbeda
- Recording date: 20181102

Heptio Contour is an "ingress controller" that uses Kubernetes objects to configure Envoy.
With Contour v0.6 and v0.7, it has introduced a new CRD called the "IngressRoute".
We are going to dig into what this is, how it works and why Contour went this direction.


<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=BSKU6QHOvVE
" target="_blank"><img src="http://img.youtube.com/vi/BSKU6QHOvVE/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review
- 00:19:00 - Contour/Ingress

## Week in Review

- [IBM buys Red Hat](https://newsroom.ibm.com/2018-10-28-IBM-To-Acquire-Red-Hat-Completely-Changing-The-Cloud-Landscape-And-Becoming-Worlds-1-Hybrid-Cloud-Provider)
- [The full interview with IBM Ginni Rometty and Red Hat Jim Whitehurst CEOs ](https://youtu.be/TAWIxYa5hUI)
- [Kubernetes 1.13 code slush is already here](https://groups.google.com/forum/#!topic/kubernetes-dev/qbd_VlVqSeg)
- [Kubernetes LTS Discussion ongoing](https://github.com/kubernetes/community/issues/2720) - WG in the process of being established to investigate possibilities <- lots of great discussion here
- [KubeCon contributor Summit](https://contributor.kubernetes.io/events/2018/12-contributor-summit/)
- [gravitational/gravity](https://github.com/gravitational/gravity) - packaging solution for applications/clusters
- [replicatedhq/ship](https://github.com/replicatedhq/ship) - packaging/deployment solution built on kustomize (can init from helm)

## Swag!

We will pick a random 7 people who submit to https://j.hept.io/tgik-roll-call and send them some fun swag!  You'll need the secret password from watching the video.  We'll take everyone that submits up to 72 hours from airing.

## Show Notes

- 00:22 Load Balancers and Ingress
- Network traffic comes the user, goes to a software load balancer, say ELB.
- ELB routes into a software load balancer in your cluster, say nginx.
    - 00:34 Question from the audience: Is there a performance hit? In AWS the NLB is transparent to the network.
- nginx picks a service in a cluster to route traffic to. If you say foo.com, go to service 1. That service is considered the _upstream_. It's very confusing. "Down is upstream"
- 00:29 Limitations in Ingress protecting namespaces from stepping on each other
- 00:31 Clarifying what we mean by upstream
- 00:34 Installing Contour
    - [Deploying on AWS with NLB](https://github.com/heptio/contour/blob/master/docs/deploy-aws-nlb.md)
    - [Documentation](https://github.com/heptio/contour/tree/master/docs)
- 00:37 Investigating the yaml we use to install Contour
    - Pro tip to new listeners: Joe will typically spend time going through yaml files before deployment to see exactly what something does on your cluster before applying. Good habit to get in to!
- 00:50 We're up and running!


## Reference Links

* [Contour v0.6 release](https://github.com/heptio/contour/releases/tag/v0.6.0)
* [Contour v0.7 release](https://github.com/heptio/contour/releases/tag/v0.7.0)
* Documentation/examples: [IngressRoute doc](https://github.com/heptio/contour/blob/master/docs/ingressroute.md), [IngressRoute examples](https://github.com/heptio/contour/tree/master/deployment/example-workload/ingressroute)
* Blog Posts:
  * [Introducing Heptio Contour 0.6](https://blog.heptio.com/introducing-heptio-contour-0-6-ecaa5ee6a67d)
  * [Improving the multi-team Kubernetes ingress experience with Heptio Contour 0.6](https://blog.heptio.com/improving-the-multi-team-kubernetes-ingress-experience-with-heptio-contour-0-6-55ae0c0cadef)
  * [Flexible software deployment patterns with IngressRoute](https://blog.heptio.com/flexible-software-deployment-patterns-with-ingressroute-a49a43253992)
  * [Heptio Contour 0.7 Release Brings Improved Ingress Control and Request-Prefix Rewriting Support](https://blog.heptio.com/heptio-contour-0-7-release-brings-improved-ingress-control-and-request-prefix-rewriting-support-bce325ba3c4b)
