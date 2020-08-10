# Episode 127 : GitOps with Steve Wade

- Hosted by @mauilion and @swade1987

- 08/07/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=F70wRexHIwg
" target="_blank"><img src="http://img.youtube.com/vi/F70wRexHIwg/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core K8s

- [1.19rc4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#v1190-rc4) is out.
    - Still on track for Aug 25-ish release
    - [componentstatus is deprecated](https://github.com/kubernetes/kubernetes/pull/93570)
    - What's with all the boring stuff the past few weeks? Stephen Augustus wants to let you know: "very large SIG Testing/Release effort to clean up tests, pruning our backlogs; very boring things" 

### Cloud Native Ecosystem

- New open source lightweight and extensible service mesh: [OSM: Open Service Mesh](https://openservicemesh.io/)
- [Happy Sysadmin](https://sysadminday.com/) Day last week! (Did we miss this last week?)
- Emanuel Evans dives into [deconstructing Kubernetes networking](https://eevans.co/blog/deconstructing-kubernetes-networking/)
- David Lorite Solanas walks us through [monitoring etcd](https://sysdig.com/blog/monitor-etcd/)
- [QuakeKube](https://github.com/criticalstack/quake-kube) apparently exists, so we had to mention it because it's Quake.
- How [Lyft improved the scalability of cronjobs](https://eng.lyft.com/improving-kubernetes-cronjobs-at-scale-part-1-cf1479df98d4) by Kevin Yang

## Show Notes
* what the heck is a gitops?
* how does deployment relate to gitops
* directional flows
- [Flux Architectural Diagram](https://github.com/fluxcd/flux/blob/master/docs/_files/flux-cd-diagram.png)
- [Helm Operator Architectural Diagram](https://github.com/fluxcd/helm-operator/blob/master/docs/_files/fluxcd-helm-operator-diagram.png) 
- [How Kustomize works](https://github.com/kubernetes-sigs/kustomize/blob/master/docs/images/base.jpg)
- Dig into [https://github.com/swade1987/gitops-with-kustomize](https://github.com/swade1987/gitops-with-kustomize)
- Dig into [https://github.com/swade1987/gitops-with-secrets](https://github.com/swade1987/gitops-with-secrets)
- Best practices for structuring your git repos


## Reference Links
