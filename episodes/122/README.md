# Episode 122 : Grokking Kubernetes : DNS

- Hosted by @mauilion 
- Recording date: 2020-06-26

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=SeNwPNo7gxw
" target="_blank"><img src="http://img.youtube.com/vi/SeNwPNo7gxw/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### K8s Core

- 1.19 is still in beta2, no changes the last 2 weeks.
- Release team has [lots of point release goodies this week](https://groups.google.com/u/1/g/kubernetes-dev/c/j6qzDkMh-mE)
    - "As a general note, hyperkube is deprecated and will NOT be present in the v1.19.0 release or any releases > v1.19.0. Please, if you're using hyperkube for older releases (v1.18.z --> v1.16.z), make plans to migrate away from it as a deployment strategy. We have extremely limited capabilities to test hyperkube and support for it should be considered to be "best-effort"."
    
- Goodbye !!! [kubectl --export](https://github.com/kubernetes/kubernetes/pull/88649)
- [Kubernetes 1.16.12 Released](https://git.k8s.io/kubernetes/CHANGELOG/CHANGELOG-1.16.md/#v11612)

### K8s Ecosystem

- [Harbor CNCF Graduation](https://www.cncf.io/announcement/2020/06/23/cloud-native-computing-foundation-announces-harbor-graduation/)
- [Dex is now a CNCF sandbox project](https://github.com/cncf/toc/issues/472) 
- [Dex github](https://github.com/dexidp/dex/)
- [New policy project Konstraint](https://github.com/plexsystems/konstraint)
- The folks at learnk8s have whipped up a visualization on [how to quarantine a pod in Kubernetes](https://i.redd.it/k60qqcn032751.png)
- The folks at Flant also have an article on [how ConfigMaps work](https://medium.com/flant-com/configmaps-in-kubernetes-f9f6d0081dcb) with some tips
- Alex Ellis has [5 good tips on troubleshooting applications on k8s](https://levelup.gitconnected.com/5-tips-for-troubleshooting-apps-on-kubernetes-835b6b539c24)
- Weird Science! [Gocker a mini docker implemented in GO](https://unixism.net/2020/06/containers-the-hard-way-gocker-a-mini-docker-written-in-go/)
- GKE and Bayer have hit [15k node cluster sizes](https://cloud.google.com/blog/products/containers-kubernetes/google-kubernetes-engine-clusters-can-have-up-to-15000-nodes)

- [Some useful Dex blog posts (1/3)](https://medium.com/@swade1987/part-1-single-sign-on-for-kubernetes-c292fd9ee563)
- [Some useful Dex blog posts (2/3)](https://medium.com/@swade1987/part-2-sso-for-kubernetes-cli-43a518af0de8)
- [Some useful Dex blog posts (3/3)](https://medium.com/@swade1987/part-3-sso-for-kubernetes-uis-be118284abb2)

## Show Notes


## Reference Links

- [x] What's DNS
    - [x] libc vs muscl
    - [x] /etc/nsswitch
    - [x] /etc/gai.conf
    - [x] the tls connection :) 
    - [X] what are [ndots](https://pracucci.com/kubernetes-dns-resolution-ndots-options-and-why-it-may-affect-application-performances.html)
- [x] what's coredns?
- [x] Standard kube-dns deployment
- [x] metrics!!
- [ ] modifying it via configmap
    - [x] default upstreams
    - [ ] modified upstreams
    - [ ] stub dns (so cool!)
- [x] dns policy in the pod spec!
- [ ] other dns things like hosts!
- [x] The Chicken and Egg Problem
- [ ] service discovery things!
- [ ] speaking of services! service type external name!
- [ ] Let's get fun with node local dns!
- [ ] can I use dns names in network policy!
