# Episode 123 : Grokking Kubernetes : DNS Part 2

- Hosted by @mauilion 
- Recording date: 2020-07-10


<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=FKvFloyf_Kg
" target="_blank"><img src="http://img.youtube.com/vi/FKvFloyf_Kg/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### K8s Core

- 1.19.0 RC0 is live, keep an eye on the [1.19 changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md) for more info
- We are in CODE FREEZE! Check out [the status](https://groups.google.com/forum/#!topic/kubernetes-dev/-HoZA0mYbw8) from release manager Taylor Dolezal
- Otherwise quiet week, [1.18.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#downloads-for-v1185) is out
    - Expect more point releases next week ~7/15
- [CVE-2020-8558](https://discuss.kubernetes.io/t/security-advisory-cve-2020-8558-kubernetes-node-setting-allows-for-neighboring-hosts-to-bypass-localhost-boundary/11788)

### K8s Ecosystem

- From the "Thank a Kubeadm maintainer" department: Emanuel Evans talks about how he set up a [minimal viable Kubernetes](https://eevans.co/blog/minimum-viable-kubernetes/) by hand. 
- [Atlanta Kubernetes meetup](https://www.meetup.com/Kubernetes-Atlanta-Meetup/)
- [Porter](https://github.com/kubesphere/porter) is a bare metal load balancer for K8s, anyoone tried it?
- [Restricting Flux Permissions](https://itnext.io/restricting-flux-permissions-1f79372c77b5)
- [How SSO bought a cluster to its knees!](https://medium.com/@swade1987/how-sso-bought-a-cluster-to-its-knees-e0e002bff08)
- [https://artifacthub.io/](https://artifacthub.io/)


## Show Notes

It’s not DNS. There is a no way it’s DNS. It was DNS, again.

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
- [x] modifying it via configmap
    - [x] default upstreams
    - [x] modified upstreams
    - [x] stub dns (so cool!) [docs](https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/)
- [x] dns policy in the pod specs
- [x] other dns things like hosts and a custom resolv.conf
    - [x] Important assumptions about resolv.conf! All resolvers need the same view of the world.
    - [x] [docs for resolv](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/)
    - [x] [docs for host aliases](https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/)
- [x] The Chicken and Egg Problem
- [x] service discovery things!
- [x] speaking of services! service type external name!
- [x] Let's get fun with node local dns!
- [x] can I use dns names in network policy!
