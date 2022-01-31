# TGI Kubernetes Episode 170 : Testing with the new E2E Framework
- Hosted by Vladimir Vivien [@vladimirvivien](https://twitter.com/VladimirVivien)  and [John Schnake](https://www.github.com/johnschnake)
- 10/08/2021

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=TRVO-wCptao
" target="_blank"><img src="http://img.youtube.com/vi/TRVO-wCptao/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:02:30 - Week in Review
- 00:22:52 - Intro to E2E Framework
- 00:28:40 - Intro to Sonobuoy
- 00:47:00 - Dive into E2E Framework and Examples

## Week in Review
- TCE: https://tanzucommunityedition.io/ ([roadmap](https://github.com/vmware-tanzu/community-edition/blob/main/ROADMAP.md)) and new Tanzu Framework: https://github.com/vmware-tanzu/tanzu-framework
  - Amanda Katona and Josh Rosso gave a DevOps Loop [talk](https://devopsloop.io/sessions/tanzu-community-edition) about Tanzu Community Edition 
- [DevOps Loop 2021 agenda](https://devopsloop.io/agenda) and the [YouTube playlist](https://www.youtube.com/watch?v=jKZ-_zYLLKg&list=PLAdzTan_eSPSe5sm-KpWHfiLft7pV2v5c). The talks I watched:
  - [Evolving DevOps in the Age of Cloud Native](https://devopsloop.io/sessions/evolving-devops-in-the-age-of-cloud-native) with Joe Beda
  - [Rethinking the SDLC](https://devopsloop.io/sessions/rethinking-the-sdlc) with Emily Freeman
  - [VMware Tanzu Community Edition: a First Look](https://devopsloop.io/sessions/tanzu-community-edition) with Amanda and Josh
  - [The History of CI/CD](https://devopsloop.io/sessions/history-of-ci-cd) with Kat Cosgrove
  - [The Business Benefits of GitOps](https://devopsloop.io/sessions/business-benefits-of-gitops) with Cornelia Davis
- [Cartographer](https://cartographer.sh/posts/building-paths-to-prod/) is a super interesting project. It flips the script on supply chains by being unified glue between the tools you already use.  It brings choreography to supply chains in a Kubernetes native way. https://twitter.com/OssCartographer/status/1445378516674035722
- [ChainGuard](https://chainguard.dev) Applying Zero-Trust principles to supply chain security to make the software lifecycle secure by default.
- Kubernetes Cluster API 
  - CAPI reaches production readiness with the [release of version v1.0](https://www.cncf.io/blog/2021/10/06/kubernetes-cluster-api-reaches-production-readiness-with-version-1-0/)
  - [Introducing Cluster API's ClusterClass and Managed Topologies](https://kubernetes.io/blog/2021/10/08/capi-clusterclass-and-managed-topologies/)
- New release (v1.7.0) of Velero https://github.com/vmware-tanzu/velero/releases/tag/v1.7.0
- Project Open Service Mesh now supports Contour https://github.com/openservicemesh/osm/releases/tag/v0.10.0
- Kubernetes SIGs news
  - SIG K8s Infra is looking for contributors to help migrate resources from Google owned infrastructure to community owned, mostly to help with prow jobs. For more information visit: [sig-k8s-infra](https://github.com/kubernetes/community/tree/master/sig-k8s-infra)
  - SIG Node is looking for contributors to help to track their CI dashboard and look upon the failing and flaky tests. See the [CI Dashboard](https://testgrid.k8s.io/sig-node). For more information see [sig-node] (https://github.com/kubernetes/community/tree/master/sig-node)
  - SIG Auth is looking for contributors to help them with, KMS-Plugin: https://bit.ly/3ld2Y5f Audit logging: https://bit.ly/3a4HQHQ For more information see [sig-auth](https://github.com/kubernetes/community/tree/master/sig-auth)
- [Kubernetes Steering Committee](https://github.com/kubernetes/community/tree/master/committee-steering) elections are coming up, the [list of nominees](https://github.com/kubernetes/community/issues) is getting longer
- Oct 13th, the Linux Foundation is holding [OSPOlogy](https://community.linuxfoundation.org/events/details/lfhq-todo-group-presents-ospology-the-state-of-ospos-2021/) to assess the current state of open source programs.




## Show Notes

Projects mentioned

* [kubernetes-sigs/E2E-Framework](https://github.com/kubernetes-sigs/e2e-framework)
  * John found an issue - Let's discuss
  * Usage in the wild - https://github.com/K8sbykeshed/k8s-service-lb-validator
* [Sonobuoy](https://github.com/vmware-tanzu/sonobuoy)
    * [Plugins](https://github.com/vmware-tanzu/sonobuoy-plugins)
        * kube-hunter
        * e2e, windows-e2e
        * Build your own!
* [Kubernetes test](https://github.com/kubernetes/kubernetes/tree/master/test)
* [KubeTest2](https://github.com/kubernetes-sigs/kubetest2)
* Created two issues for e2e-test-framework:
    * [Namespace examples](https://github.com/kubernetes-sigs/e2e-framework/issues/61)
    * [Dry-run](https://github.com/kubernetes-sigs/e2e-framework/issues/62)

