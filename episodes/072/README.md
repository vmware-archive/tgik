# Episode 072 : Kustomize and friends

- Hosted by @jbeda
- April 12, 2019
- Live notes: https://hackmd.io/dn0KfrC3R3awmHpEgFMQSw

Episode issue: https://github.com/heptio/tgik/issues/111

<!--- Thumbnailed embed of the video, NFnpUlt0IuM is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=NFnpUlt0IuM
" target="_blank"><img src="http://img.youtube.com/vi/NFnpUlt0IuM/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:00 - Week in Review
- 00:15:00 - Kustomize
-

## Week in Review

- [Blue/Green Deployments with Contour IngressRoute](https://blogs.vmware.com/cloudnative/2019/04/12/blue-green-deployments-contours-ingressroute/)
- [CRI-O is now in CNCF Incubation](https://www.cncf.io/blog/2019/04/08/cncf-to-host-cri-o/)
- An update to Alexis Ducastel's [CNI benchmarks](https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-updated-april-2019-4a9886efe9c4)
- [Network Service Mesh](https://networkservicemesh.io/) is now a CNCF sandbox project
- Lots of news from Google at GCP Next but unclear how much of it is open vs. Google specific services.
- Some k8s upstream news - SIG Cluster Lifecycle had a meeting on [how to get started contributing](https://youtu.be/Bof9aveB3rA) and SIG Windows is running [a survey on usecases](https://pollev.com/michaelmicha980) and would like your input.
- Good looking article from Bonsai Cloud on [PodSecurityPolicy](https://banzaicloud.com/blog/pod-security-policy/)
- [Security update for Istio/Envoy](https://istio.io/blog/2019/announcing-1.1.2/)
- [Where you can find](https://k8s.vmware.com/kubecon-europe-2019/) us at Kubecon EU 2019 - has all of our sessions and talks, including the live TGIK at Kubecon!
- Some new projects in this space: [ytt](https://get-ytt.io/) and [kapp](https://get-kapp.io/)


## Show Notes

Issue from Phil with some of the snags we hit during the episode: https://github.com/kubernetes-sigs/kustomize/issues/973

## Reference Links

* [Kustomize website](https://kustomize.io/)
* [Github Repo](https://github.com/kubernetes-sigs/kustomize)
    * [Repo Docs](https://github.com/kubernetes-sigs/kustomize/tree/master/docs).  Good stuff here including a glossary and an annotated `kustomize.yaml`
    * [Repo Examples](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)
* [Upstream Kubectl docs](https://kubectl.docs.kubernetes.io/pages/app_management/apply.html)
* kubectl + kustomize + KEP drama
    * [Article with blow by blow](https://gravitational.com/blog/kubernetes-kustomize-kep-kerfuffle/)
    * [Kustomize subcommand KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cli/kustomize-subcommand-integration.md)
* [Article from JetStack](https://blog.jetstack.io/blog/kustomize-cert-manager/) on kustomize+cert-manager
* [Helm vs Kustomize](https://codeengineered.com/blog/2018/helm-kustomize-complexity/)