# Episode 124 : KUDO

- Hosted by @jbeda
- Recording date: 2020-07-17

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=XNi8uO0PwPo
" target="_blank"><img src="http://img.youtube.com/vi/XNi8uO0PwPo/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:03:50 - Week in Review
- 00:23:21 - Start talking about KUDO
- 00:24.50 - Comparison to other projects
- 00:26:13 - Operators defined and KUDO overview
- 00:33:06 - Limitations of KUDO and some future plans
- 00:35:16 - Using KUDO with Zookeeper operator
- 00:35:44 - Install KUDO commandline
- 00:36:07 - kubectl kudo init
- 00:46:24 - Install zookeeper
- 00:50:37 - Run validation task (and hit a bug!)
- 01:00:17 - Noodling on how backup/recovery would integrate
- 01:02:25 - Operator index and publishing operators
- 01:07:30 - Looking at zookeeper operator
- 01:13:15 - Compared to Helm, templating options, gitops?
- 01:17:28 - Validation as Job
- 01:19:03 - Exploring other Operators
- 01:22:32 - Different Task types
- 01:26:33 - Testing with Kuttl
- 01:31:05 - Final thoughts and goodbye

## Core:
- CVEs:
    - issues.k8s.io/93032 CVE-2020-8557 Node disk DOS by writing to container /etc/hosts
    - issues.k8s.io/90259 CVE-2020-8558 net.ipv4.conf.all.route_localnet=1 opens security issue
    - issues.k8s.io/92914 CVE-2020-8559 Privilege escalation from compromised node to cluster

- New Releases:
    - [Kubernetes v1.16.12](https://groups.google.com/g/kubernetes-announce/c/IALpWZDLOYI)
    - [Kubernetes v1.17.8](https://groups.google.com/g/kubernetes-announce/c/tEZjvXRNA0Q)
    - [Kubernetes v1.18.5](https://groups.google.com/g/kubernetes-announce/c/XtgeDsAp91s)

## Week in Review:
- [Contour accepted in CNCF](https://www.cncf.io/blog/2020/07/07/toc-accepts-contour-as-incubating-project/)
- [Octant v0.14](https://github.com/vmware-tanzu/octant/releases/tag/v0.14.0)
- [CNCF Ambassador spotlight series! with Alex Ellis](https://www.cncf.io/blog/2020/07/16/cncf-ambassador-spotlight-alex-ellis/)
- [A wild new security cert coming in November! CKS](https://www.cncf.io/blog/2020/07/15/certified-kubernetes-security-specialist-cks-coming-in-november/)
- [Netflix has moved from Mesos to Kubernetes!](https://twitter.com/aspyker/status/1283836267646431234?s=20)
- [fluent-bit v1.5](https://www.cncf.io/blog/2020/07/14/fluent-bit-v1-5-lightweight-and-high-performance-log-processor/)
- [Kubecon EU virtual complimentary pass options!](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/)
- [A webinar on Antrea and cni data plane acceleration](https://www.cncf.io/webinars/securing-and-accelerating-the-kubernetes-cni-data-plane-with-project-antrea-and-nvidia-mellanox-connectx-smartnics/)
- [Tim Hockin sharing knowledge on services!](https://twitter.com/ahmetb/status/1283825893102612480?s=20)
- [A great "what happens when"](https://github.com/jamiehannaford/what-happens-when-k8s/blob/master/README.md)
- [Istio and Open Usage Commons](https://istio.io/latest/blog/2020/open-usage/)
- [Scalability, with Wojciech Tyczynski](https://kubernetespodcast.com/episode/111-scalability/)

## Show Notes

* [KUDO](https://kudo.dev/)
* [Main github repo](https://github.com/kudobuilder/kudo)
* [Operators repo](https://github.com/kudobuilder/operators)
* Cool test tool that isn't KUDO specific: [kuttl](https://kuttl.dev/)

