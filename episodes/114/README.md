# Episode 114 : Grokking Kubernetes : Scheduling Part 2

- Hosted by @mauilion
- Recording date: 2020-04-17

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=o_gT_LfFn4s
" target="_blank"><img src="http://img.youtube.com/vi/o_gT_LfFn4s/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review


### Core Kubernetes

- Releases
    - Current: [1.18.2](https://git.k8s.io/kubernetes/CHANGELOG/CHANGELOG-1.18.md/#v1182) 
    - Older series with new point releases: [1.17.5](https://git.k8s.io/kubernetes/CHANGELOG/CHANGELOG-1.17.md/#v1175), and [1.16.9](https://git.k8s.io/kubernetes/CHANGELOG/CHANGELOG-1.16.md/#v1169).
- [SIG PM is being retired](https://groups.google.com/forum/#!topic/kubernetes-dev/RU8iINF758Y)


### Ecosystem

- [@embano1 finds a neat feature of etcd](https://twitter.com/embano1/status/1250540206010310659?s=20)
- [Argocd 1.4.3/1.5.2 released fixes Critical git CVE](https://argoproj.github.io/argo-cd/security_considerations/#CVE-2020-5260-possible-git-credential-leak)
- [kubeweekly.io](kubeweekly.io)
- [Pluto](https://github.com/FairwindsOps/pluto) -  A cli tool to help discover deprecated apiVersions in Kubernetes from Fairwinds
- [kubernetes-node-local-dns-cache](https://povilasv.me/kubernetes-node-local-dns-cache/)
## Show Notes
- [x] kubectl explain pod.spec.nodeName
- [x] kubectl explain pod.spec.nodeSelector
- [x] kubectl explain pod.spec.tolerations
- [x] kubectl explain node.spec.taints
- [x] kubectl explain pod.spec.affinity
- [ ] kubectl explain pod.spec.topologySpreadConstraints
- [ ] [assigning pods to nodes](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector)
- [ ] [admission controller: PodNodeSelector](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#podnodeselector)
- [ ] [kube-scheduler implementation](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/#kube-scheduler-implementation)
- [ ] What's filter?
- [ ] [topologySpreadConstraints Kep](https://github.com/kubernetes/enhancements/blob/master/keps/sig-scheduling/20190221-pod-topology-spread.md)

## Reference Links

 - [Checks and Balances with Security in Kubernetes](https://github.com/setns/live)
    - Nova is doing live streams at 11am Pacific before TGIK. Come join the party! 
    - twitch.tv/setns
 - TGIK 96: [Grokking Kubernetes : kube-scheduler](https://www.youtube.com/watch?v=XxVHNWoZO_c)
 - [Kube academy](https://kube.academy/)
 - [Keep you kubernetes cluster balanced the secret to high availability ](https://itnext.io/keep-you-kubernetes-cluster-balanced-the-secret-to-high-availability-17edf60d9cb7)

