# Episode 117 : Grokking Kubernetes : ETCD
- Hosted by @mauilion 
- 05/08/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=bQw3UZgYmOk
" target="_blank"><img src="http://img.youtube.com/vi/bQw3UZgYmOk/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core K8s

- k8s 1.19alpha3, here are the [changes since the last alpha](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#changelog-since-v1190-alpha2)
    - Pro-tip, [always check the API changes/deprecations](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#deprecation), that will save you future pain. 
        - Deprecation checkers (https://github.com/swade1987/deprek8ion, https://github.com/FairwindsOps/pluto)

### K8s and Cloud Native Ecosystem

- [Running Kind inside of a k8s cluster for CI](https://d2iq.com/blog/running-kind-inside-a-kubernetes-cluster-for-continuous-integration)
- An overview of [Pod Security policies](https://developer.squareup.com/blog/kubernetes-pod-security-policies/)
- [Helm and Kustomize working together](https://povilasv.me/helm-kustomze-better-together/)
- [State of Persistant storage](https://itnext.io/state-of-persistent-storage-in-k8s-a-benchmark-77a96bb1ac29) - (with benchmarks)
- Steve's Wade [Mettles journey towards throw-away clusters from Kube Philly meetup](https://youtu.be/3sIKwz0LOI0)
- Steve Wade's [GitOps at Mettle Webinar next Friday](https://t.co/g45KFNG3Xv?amp=1)


## Show Notes
- [x] lab setup!
- [ ] weavework/footloose
- [ ] etcdadm
- [ ] stacked etcd
- [ ] external etcd
    - [ ] why is it simpler?
    - [ ] why is it better? 
- [ ] what about load balancing etcd?
- [ ] what is etcd?
    - [ ] versioning whaaaaaat?
- [ ] etcd metrics!
    - [ ] where is /metrics?
- [ ] troubleshooting etcd
- [ ] how does the apiserver talk to etcd?
- [ ] is there authentication?
- [ ] can I store other stuff? 
- [ ] events?

## Reference Links

- [Footloose](https://github.com/weaveworks/footloose)
- [etcdadm](https://github.com/kubernetes-sigs/etcdadm)