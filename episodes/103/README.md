# Episode 103 : Cilium: A Second Look

- Hosted by @joshrosso
- Recording date: 2020-01-31

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=se4soMN4fDg
" target="_blank"><img src="http://img.youtube.com/vi/se4soMN4fDg/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://youtu.be/se4soMN4fDg)
- [00:04:45 - Week in Review](https://youtu.be/se4soMN4fDg?t=284)
- [00:11:21 - Agenda Overview and Topic Selection](https://youtu.be/se4soMN4fDg?t=677)
- [00:15:08 - Deploying without kube-proxy (and with CRD backend)](https://youtu.be/se4soMN4fDg?t=929)
- [01:00:56 - Deploying and using Hubble](https://youtu.be/se4soMN4fDg?t=3656)
- [01:36:53 - CNI-Chaining](https://youtu.be/se4soMN4fDg?t=5813)

## Week in Review

#### Core

- k8s 1.18 is now in [enhancements freeze]( http://bit.ly/k8s-1-18-enhancements)
    - We're still on [1.18alpha2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.18.md#v1180-alpha2), nothing new to report
    - Keep an eye on the [patch release page](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)
- The schedule for KubeCon + CloudNativeCon [has been published](https://www.cncf.io/announcement/2020/01/29/cloud-native-computing-foundation-announces-schedule-for-kubecon-cloudnativecon-europe-2020/)
    - Don't forget to check out the [Contributor Summit](https://events.linuxfoundation.org/kubernetes-contributor-summit-europe/) if you want to start contributing to k8s, registration is now open and does fill up quickly!
    - Cloud Native Rejekts [is happening for Amsterdam](https://cloud-native.rejekts.io/)
    - Our own Duffie got a talk accepted, see [Seccomp Security Profiles and You: A Practical Guide](https://sched.co/ZetL)
    - Friends of the show Ian Coldwater and Brad Geesaman also have a talk [Advanced Persistence Threats: The Future of Kubernetes Attacks](https://sched.co/ZesN)
- [Latest Jepsen Results against etcd 3.4.3](https://etcd.io/blog/jepsen-343-results/)


#### Ecosystem

- [How a GCP Persistent Disk Incident Snowballed into a 23-Hour Outage -- and Taught Us Some Important Lessons](https://grafana.com/blog/2020/01/23/how-a-gcp-persistent-disk-incident-snowballed-into-a-23-hour-outage-and-taught-us-some-important-lessons/)
- Someone put together a [set of Dad Joke microservices](https://github.com/yesinteractive/dad-jokes_microservice) to test your cluster. O_O
- Alex Ellis gives us a good summary of [building containers without Docker](https://blog.alexellis.io/building-containers-without-docker/)
- [Implementing Container Runtime Shim: Interactive Containers](https://iximiuz.com/en/posts/implementing-container-runtime-shim-3/?utm_medium=reddit&utm_source=r_kubernetes)
    - Editor's note, check the previous posts in this series, there's gold in there! --jorge


## Cilium

**Agenda:**

* [x] Kube-proxy REPLACEMENT!?!
* [x] CRD Backend
* [x] Hubble
* [x] CNI-Chaining

## References

* Hubble: https://cilium.io/blog/2019/11/19/announcing-hubble
* 1.6 Release: https://cilium.io/blog/2019/08/20/cilium-16
* Liberating Kubernetes From Kube-proxy and Iptables: https://www.youtube.com/watch?v=bIRwSIwNHC0
    * Kubecon Talk, Martynas Pumputis
* Kris's Cilium TGIK: https://www.youtube.com/watch?v=I8Tp7jU2oJk
