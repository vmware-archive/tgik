# Episode 92 : Continuing Minecraft Controller

- Hosted by @jbeda
- 10/04/2019

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=pzXyZXcFa7o
" target="_blank"><img src="http://img.youtube.com/vi/pzXyZXcFa7o/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK! Including news of the week.
- 00:23:43 - Starting in on coding the controller

## Week in Review

- Steering Committee Election [Results](https://kubernetes.io/blog/2019/10/03/2019-steering-committee-election-results/)
    - Check out [Nikhita's talk at KubeCon](https://www.youtube.com/watch?v=Bho4miiByP0)
- Kubevirt [joins](https://blog.openshift.com/kubevirt-joins-cloud-native-computing-foundation/) CNCF
- Yep, 1.17.0 alpha is [already here](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.0-alpha.1)
- Kenta Iso did a talk on Kubernetes Controllers(https://speakerdeck.com/govargo/inside-of-kubernetes-controller)
- New CVE for Kubernetes. [CVE-2019-11253: Kube Apiserver vulnerable to Billion Laughs variant](https://github.com/kubernetes/kubernetes/issues/83253) Shout out to Rory a frequent viewer of tgik!
- [Kontena Lens](https://blog.kontena.io/kontena-lens-desktop-app/)
- Really neat loadbalancer solution for k8s from [Alex Ellis and inlets](https://blog.alexellis.io/ingress-for-your-local-kubernetes-cluster/)
- Great [Meet Our Contributors](https://www.youtube.com/watch?v=rZ4iBBZu3Z4) this week with Davanum Srinivas (Dims), Nikhita Raghunath, Paris Pittman, and Kiran Oliver
    - This is a monthly show from the project where k8s contributors share their tips, check it out if you're interested in how the sausage gets made.
- If you are or are interested in becoming a contributor check out the [contributor summit!](https://kubernetes.io/blog/2019/09/24/san-diego-contributor-summit/)
    - New user track is on a waiting list but signing up never hurts!
- [Hacktoberfest](http://hacktoberfest.com/) started this week! A few Kubernetes projects are participating with labeled GitHub issues. Check out [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/issues?q=is%3Aissue+is%3Aopen+label%3Ahacktoberfest).
- [opengovernance.dev](https://opengovernance.dev) - A good definition around what OSS Open Governance looks like.


## Show Notes

This is a continuation of [TGIK 083](https://github.com/heptio/tgik/tree/master/episodes/083).

## Reference Links
Code for the project is checked in to https://github.com/jbeda/kinecraft.

- [Programming Kubernetes](https://learning.oreilly.com/library/view/programming-kubernetes/9781492047094/)
- [Kubebuilder Docs](https://book.kubebuilder.io/quick-start.html)
- [Minecraft Image](https://github.com/itzg/dockerfiles/tree/master/minecraft-server)
- [rcon-cli](https://github.com/itzg/rcon-cli) uses [rcon library](https://github.com/james4k/rcon)
- [Minecraft protocol](https://wiki.vg/Protocol#Handshake)
