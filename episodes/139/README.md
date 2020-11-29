# Episode 139 : Running Telco workloads on Kubernetes
- Hosted by @opowero and @AbeeV
- 11/21/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=v9GQYTc7TJc
" target="_blank"><img src="http://img.youtube.com/vi/v9GQYTc7TJc/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - [Welcome to TGIK!](https://youtu.be/v9GQYTc7TJc?t=7)
- 00:02:30 - [Week in Review](https://youtu.be/v9GQYTc7TJc?t=150)
- 00:16:20 - [Intro to telco workload session](https://youtu.be/v9GQYTc7TJc?t=980)
- 00:19:40 - [History of telco "generations"](https://youtu.be/v9GQYTc7TJc?t=1180)
- 00:24:00 - [4G network architecture](https://youtu.be/v9GQYTc7TJc?t=1440)
- 00:35:15 - [5G network architecture](https://youtu.be/v9GQYTc7TJc?t=2115)
- 00:39:40 - [Additions to Kubernetes for CNF support](https://youtu.be/v9GQYTc7TJc?t=2380)
- 00:42:03 - [multus (multi interface)](https://youtu.be/v9GQYTc7TJc?t=2523)
- 01:01:12 - [SR-IOV and DPDK (bandwidthl n/s, e/w)](https://youtu.be/v9GQYTc7TJc?t=3672)
- 01:25:00 - [Node Discovery](https://youtu.be/v9GQYTc7TJc?t=5100)
- 01:29:37 - [Questions and Wrap Up](https://youtu.be/v9GQYTc7TJc?t=5377)

## Week in Review

### KubeCon NA Virtual 2020

#### Talks
All talks will be publicy avaiable after 4th december on [CNCF youtube channel](https://www.youtube.com/channel/UCvqbFHwN-nwalWPjPUKpvTA)*
 - Kubeconf Infographic summaries by [@awsgeek Jerry Hargrove](https://www.awsgeek.com/KubeCon-Virtual-2020/)
 - [Infrastructure for Entertainment](https://kccncna20.sched.com/event/ekD3/infrastructure-for-entertainment-justin-garrison-amazon), Justin Garrison 
 - [PKI the Wrong Way: Simple TLS Mistakes and Surprising Consequences](https://kccncna20.sched.com/event/ekES/pki-the-wrong-way-simple-tls-mistakes-and-surprising-consequences-tabitha-sable-datadog), Tabitha Sable
 - Cheryl Hung Keynote: [Are Certifications Worth It? ](https://kccncna20.sched.com/event/eoDp/keynote-are-certifications-worth-it-cheryl-hung-vice-president-ecosystem-cloud-native-computing-foundation)
 - Day0 sig-security CTF [Capture the Flag](https://sched.co/fEjI):
   - [Control-plane.io github repo for the CTF simulator](https://t.co/79rFU6nZy1?amp=1)
   - [Recap slides](https://drive.google.com/file/d/1am_eYqyQjlJO10TWuRtfP8VjLikIDeDX/view)
- Ian Coldwater, Duffie Cooley, Brad Geesaman and Rory McCune Panel [Keynote: SIG-Honk AMA Panel: Hacking and Hardening in the Cloud Native Garden](https://kccncna20.sched.com/event/eoIZ/keynote-sig-honk-ama-panel-hacking-and-hardening-in-the-cloud-native-garden-ian-coldwater-independent-duffie-cooley-independent-brad-geesaman-co-founder-darkbit-rory-mccune-principal-consultant)
- And lastly, this one is just for fun: [Having Cloud Native fun with HonkCTL](https://kccncna20.sched.com/event/ekBS/having-cloud-native-fun-with-honkctl-jeffrey-sica-red-hat)

Thanks to Walid Shaari for sharing his list of favorite talks so far! 

#### Other KubeCon News

- [My key takeaways on Cloud-Native KubeCon Conference (CNCF) 2020 thus far](https://itnext.io/my-key-takeaways-on-cloud-native-kubecon-conference-cncf-2020-thusfar-cb1dfae85a4e) from JPantsjoha

### Core K8s 

- Core devs are KubeConning, watch this space next week!
- [CVE-2020-8569](https://groups.google.com/g/kubernetes-dev/c/_2a2BVACg3s/m/taaiikXRAgAJ)
- Kubernetes [1.20.0beta2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md/#v1200-beta2)
    - Everything is frozen, home stretch!

### K8s and Cloud Native Ecosystem

- [CertManager donated to CNCF](https://www.businesswire.com/news/home/20201117005452/en/Jetstack-Donates-cert-manager-to-Cloud-Native-Computing-Foundation)
- [Certified Kubernetes Security exam now available](https://www.cncf.io/announcements/2020/11/17/kubernetes-security-specialist-certification-now-available/)
   - [Linuxfoundation kubernetes security essentials course lfs260](https://training.linuxfoundation.org/training/kubernetes-security-essentials-lfs260/)
   - TGIKer Walid Shaari's CKS Preperation curated [CKS resources repository](https://github.com/walidshaari/Certified-Kubernetes-Security-Specialist)
   - Killer.sh Udemy [CKS full course and exam simulator](https://www.udemy.com/course/certified-kubernetes-security-specialist/?couponCode=CKS-KILLER-SHELL)  - use promotional code CKS-KILLER-SHELL
- [Cloud Native Buildpacks promoted from incubation to sandbox](https://www.cncf.io/blog/2020/11/18/toc-approves-cloud-native-buildpacks-from-sandbox-to-incubation/)
- Intuit has a 2 part blog series on TurboTax moving to Kubernetes. 
    - [Part 1](https://medium.com/intuit-engineering/turbotax-moves-to-kubernetes-an-intuit-journey-part-1-aa861c061a11)
    - [Part 2](https://medium.com/intuit-engineering/turbotax-moves-to-kubernetes-an-intuit-journey-part-2-f5217772fbb6)
- Datadog's [container survey](https://www.datadoghq.com/container-report/) has been updated for Nov 2020
    - Lots of real world data in here! 

## Show Notes

- [SRIOV blog](https://blog.scottlowe.org/2009/12/02/what-is-sr-iov/)
- [SRIOV github](https://github.com/k8snetworkplumbingwg/sriov-network-device-plugin)
- [DPDK](https://www.dpdk.org/)
- [Multus](https://github.com/intel/multus-cni#quickstart-installation-guide)
- [Topology manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager/)
- [Node Feature Discovery](https://github.com/kubernetes-sigs/node-feature-discovery)
- [3GPP](https://www.3gpp.org/)
- [5GPP](https://5g-ppp.eu/)
- [Istio](https://istio.io/latest/docs/setup/getting-started/)


