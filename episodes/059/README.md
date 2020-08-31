# Episode 059 : CNAB

- Hosted by @krisnova
- Recording date: 2018-12-07
- Roll Call: [http://j.hept.io/tgik-roll-call](https://hello.heptio.com/tgik-roll-call/) - Listen in during the show for the secret codeword to enter our raffle. 7 winners chosen per episode, and it helps us support the [Girls Who Code](https://hello.heptio.com/tgik-roll-call/clkn/http/www.blackgirlscode.com/) Foundation.

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=R_Wul_xdHsk
" target="_blank"><img src="http://img.youtube.com/vi/R_Wul_xdHsk/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review
- 00:27:00 - Kubecon/Heptio information
- 00:30:00 - Let's get to CNAB! (See notes below)


## K8s Week in Review

- [Kubernetes v1.13 is out](https://discuss.kubernetes.io/t/kubernetes-v1-13-0-released/3707)
    - [Release Notes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.13.md#kubernetes-113-release-notes)
    - [Announcement](https://kubernetes.io/blog/2018/12/03/kubernetes-1-13-release-announcement/)
        - Kubeadm is now GA! Get all the info from [this blog post](https://kubernetes.io/blog/2018/12/04/production-ready-kubernetes-cluster-creation-with-kubeadm/)
        - CSI is now GA (so is the spec!)
            - [List of CSI Drivers](https://kubernetes-csi.github.io/docs/Drivers.html)
        - CoreDNS is now the default DNS server
- [CVE-2018-1002105](https://discuss.kubernetes.io/t/kubernetes-security-announcement-v1-10-11-v1-11-5-v1-12-3-released-to-address-cve-2018-1002105/3700)
    - v1.10.11, v1.11.5, and v1.12.3 released (all v1.13.x are not exposed)
    - [Tracking Issue](https://github.com/kubernetes/kubernetes/issues/71411)
    - Gravitational made a [test tool](https://github.com/gravitational/cve-2018-1002105) to check your cluster. YMMV, etc, check out the [corresponding blog post](https://gravitational.com/blog/kubernetes-websocket-upgrade-security-vulnerability/) for more info.

- Kubedex's [90 Days of production on EKS](https://kubedex.com/90-days-of-aws-eks-in-production/)
- [Help us rename Heptio Ark](https://github.com/heptio/ark/issues/1122) - submissions due by Sunday 12/16
- The Upbound folks announced [Crossplane](https://blog.upbound.io/introducing-crossplane-open-source-multicloud-control-plane/)


## Show Notes

[CNAB](https://cnab.io/) - "A spec for packaging distributed apps. CNABs facilitate the bundling, installing and managing of container-native apps — and their coupled services."



## Reference Links

* [CNAB Spec](https://github.com/deislabs/cnab-spec)
* [Duffle](https://duffle.sh/)
* [Kubernetes duffle driver](https://github.com/deislabs)(https://github.com/deislabs/duffle/pull/560)
* [Build your first CNAB bundle](https://github.com/deislabs/duffle/blob/master/docs/guides/bundle-guide.md)
* [ansible-playbook-bundle](https://github.com/ansibleplaybookbundle/ansible-playbook-bundle)

## Kubecon Announcements

 * Heptio Gimbal & Heptio Contour office hours with Actapio/Yahoo Japan at Heptio Booth #P8. Thursday 12:15pm - 1:45pm.
 * [Kris Nova - You can't have a cluster[BLEEP] without a cluster - Session](https://kccna18.sched.com/event/GrYV/you-cant-have-a-cluster-bleep-without-a-cluster-kris-nova-heptio)
 * [All Heptonian Talks](https://hello.heptio.com/kubecon-2018/)
 * Contributor Summit talks and sessions will be recorded and published to YouTube, we'll post links as soon as we have them. 
 * 


## Notes from Microsofties

Shoutout to [@ralph_squillace](https://twitter.com/ralph_squillace) and [@michellenoorali](https://twitter.com/michellenoorali)

The goal of the spec is to describe "How" and installation is going to happen with multiple "Components" regardless of runtime / state management stack. 

###### Note: Multiple invocation images

Not currently working in duffle, but working on fleshing this out.
