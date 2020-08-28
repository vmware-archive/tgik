# Episode 121 : Starboard

- Hosted by @mauilion
- Recording date: 2020-06-19

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=P_QZtZJAZyU
" target="_blank"><img src="http://img.youtube.com/vi/P_QZtZJAZyU/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents 

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core K8s
- [Kubernetes v1.16.11 release](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md/#v11611)
- [Kubernetes v1.17.7 release](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md/#v1177)
- [Kubernetes v1.18.4 release](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#v1184)

- [k8s.gcr.io is moving](https://groups.google.com/forum/#!topic/kubernetes-dev/MkXnkTSJ_vs) - About two weeks until this happens, heads up if it affects you!
- Notes from the monthly k8s community meeting:
    - SIG Updates:
      - SIG Release [Slides](https://docs.google.com/presentation/d/1VuS06RJWccA8ceSlu5a3hIIsVb2mge6Zz5YDpxHiBFY/edit#slide=id.g401c104a3c_0_0)
      - SIG Apps  [Slides](https://docs.google.com/presentation/d/18UcJQs3ThW6Vdgl_mdc1984uU16GwuIhuD0Pujl42xU/edit?usp=sharing)
      - Product Security Committee [Slides](https://docs.google.com/presentation/d/1TJQevF8wLRsjVRNQuUiRgwI5EO2mdt03RfifYwDoGrc/preview)
      - SIG Architecture [Slides](https://docs.google.com/presentation/d/1NytMrpVYKzFo7rLcEEHnFl8zOx05fnjs3xBSZXVE0nI/edit#slide=id.g401c104a3c_0_0)
    - [Full Notes](https://discuss.kubernetes.io/t/kubernetes-community-meeting-notes/35/78)

### Community
- [kubernetes community values](https://github.com/kubernetes/community/blob/master/values.md)
- kubecon NA deadline [extended!](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/cfp/)

### Cloud Native Ecosystem
- great writeup on a recent [cve](https://medium.com/@BreizhZeroDayHunters/when-its-not-only-about-a-kubernetes-cve-8f6b448eafa8)
- [Kubernetes Goat](https://github.com/madhuakula/kubernetes-goat) - Insecure cluster by default to help learn security
- [Argocd 1.16 release](https://blog.argoproj.io/argo-cd-v1-6-democratizing-gitops-with-gitops-engine-5a17cfc87d62)
- [Starboard Announced!](https://blog.aquasec.com/starboard-kubernetes-tools)

### Repositories:
- [https://github.com/swade1987/gitops-with-kustomize](https://github.com/swade1987/gitops-with-kustomize)
- [https://github.com/swade1987/gitops-with-secrets](https://github.com/swade1987/gitops-with-secrets)

### Media: 
- [The Popcast!](https://twitter.com/PopcastPop)
    - Kubernetes turns 6! [episode 1](https://www.youtube.com/watch?v=EnVJ2lwaADw) [episode 2](https://www.youtube.com/watch?v=BdziTqhrGeo)
- [Mettles journey towards throw-away clusters - Cloud Native Nordics (Tech Talks)](https://www.youtube.com/watch?v=zSKCHZ3wpFM)


### Blog Posts:
- [CNCF stands with the Black Lives Matter movement](https://www.cncf.io/blog/2020/06/11/statement-from-cncf-general-manager-priyanka-sharma-on-the-black-lives-matter-movement/)
- Adnand Rashid has published a PDF of his [CKA Notes](https://drive.google.com/file/d/1RhPULD1IAVgCo1KD857iCoaNKuJjQKa1/view)
- Speaking of the CKA, it's [getting updated](https://training.linuxfoundation.org/cka-program-changes-2020/), so update your notes!



## Show Notes
-

## Reference Links
- [Starboard](https://github.com/aquasecurity/starboard) 
- Integrations so far:
  - [Trivy vulnerability scanner](https://github.com/aquasecurity/trivy)
  - [Polaris config audit](https://github.com/fairwindsops/polaris)
  - [kube-bench CIS benchmark tests](https://github.com/aquasecurity/kube-bench)
  - [kube-hunter pen-testing](https://github.com/aquasecurity/kube-hunter)
  - Plugs in to [Octant](https://octant.dev)
