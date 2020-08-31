# Episode 089 : Tekton Pipelines

- Hosted by @jbeda
- Recording date: 2019-09-06

<a href="https://www.youtube.com/watch?v=CGywE84CCx8
" target="_blank"><img src="http://img.youtube.com/vi/CGywE84CCx8/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in k8s
- Lots of news out of VMworld!
    - [VMware Tanzu](https://blogs.vmware.com/cloudnative/2019/08/26/vmware-completes-approach-to-modern-applications/)
    - [Tanzu Mission Control](https://blogs.vmware.com/cloudnative/2019/08/26/vmware-tanzu-mission-control/)
    - [Project Pacific intro](https://blogs.vmware.com/cloudnative/2019/08/26/project-pacific-is-kubernetes-to-the-core/)
    - [Project Pacific Overview](https://blogs.vmware.com/vsphere/2019/08/introducing-project-pacific.html)
    - [Pacific Tech Overview](https://blogs.vmware.com/vsphere/2019/08/project-pacific-technical-overview.html)
    - [Welcome Pivotal!](https://blogs.vmware.com/cloudnative/2019/08/22/transforming-software-on-kubernetes/)
    - [Kubernetes Academy!](https://kubernetes.academy)
- 1.16 is coming soon! Remember that [api removal](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/) is coming!
- [KubeCon + CloudNativeCon Schedule is up](https://www.cncf.io/announcement/2019/09/05/cloud-native-computing-foundation-announces-schedule-for-kubecon-cloudnativecon-san-diego/)
    - Also check out the [Rejekts conference](https://cloud-native.rejekts.io/) right before
    - if you are a contributor or would like to be check out the [contributor summit](https://events.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/)
- Announcing [etcd 3.4](https://kubernetes.io/blog/2019/08/30/announcing-etcd-3-4/)! lot's of improvements and some great new features.
- In cluster api news [v1alpha2](https://github.com/kubernetes-sigs/cluster-api/releases/tag/v0.2.1) is out and represents a TON of work!
- [networking internals: endpointSlice API is coming in 1.16](https://github.com/kubernetes/enhancements/blob/master/keps/sig-network/20190603-EndpointSlice-API.md)
- Overview of [cronjobs](https://medium.com/jobteaser-dev-team/kubernetes-cronjob-101-56f0a8ea7ca2) (and kittehs)
- [Great primer on Kubernetes on the edge](https://blog.calsoftinc.com/2019/08/powering-edge-with-kubernetes-a-primer.html)
- 3 key takeways from the [Kubernetes Security Audit](https://www.stackrox.com/post/2019/09/the-kubernetes-security-audit-3-key-takeaways/)
- The Kubernetes Election process [is under way](https://github.com/kubernetes/community/tree/master/events/elections/2019), see [this post](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!topic/kubernetes-dev/DMD-p7u-82k) for more details.
- [Interesting deep dive into how "kubectl exec" works](https://erkanerol.github.io/post/how-kubectl-exec-works/) (via @mauilion)

## Show notes/links

(Listeners, feel free to add notes and links throughout the course of the show)

- https://tekton.dev/
- https://github.com/tektoncd/pipeline
- https://github.com/mgreau/tekton-pipelines-elastic-tutorials
- https://cd.foundation
- https://www.slideshare.net/VictorIglesias6/introduction-to-tekton
- [Jenkins X with Tekton Pipelines](https://jenkins-x.io/architecture/jenkins-x-pipelines/)
- https://github.com/kubernetes-sigs/kind
- https://github.com/k14s/kapp
- Kubecon talk on how Tekton runs tasks sequentially in a pod (and moar!): [Russian Doll: Extending Containers with Nested Processes](https://kccncna19.sched.com/event/Uaco/russian-doll-extending-containers-with-nested-processes-christie-wilson-google)
- SLAML Tshirt ![alt text](https://i.ibb.co/nbnBNfq/saml.jpg "SLAML it")
- https://github.com/tektoncd/pipeline/blob/master/docs/auth.md
- https://github.com/GoogleContainerTools/kaniko#running-kaniko <-- auth with kaniko
- Contact info for Tekton folks https://github.com/tektoncd/community#want-to-get-involved (including slack, etc.)
- [Blog post on Tekton from Aled @ JetStack](https://blog.jetstack.io/blog/exploring-tekton/)
- [Workshop](https://github.com/viglesiasce/tekton-workshop) to get up to speed with Tekton on GKE
