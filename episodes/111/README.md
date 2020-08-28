# Episode 111 : CloudEvents and Argo Events

- Hosted by @jbeda
- Recording date: 2020-03-27

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=LQbBgQnUs_k
" target="_blank"><img src="http://img.youtube.com/vi/LQbBgQnUs_k/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 8:00 Sysdig Kubernetes 1.18 blog post
- 10:51 CNCF CNCF Webinar kubernetes-1,18 from the release team!
- 13:00 Digitial Ocean Kubernetes tutorials in diffferent languages
- 18:35 The start of the CloudEvents and Argo Events discussion

## Week in Review
### Community
- Kubernetes 1.18 [fit and finish](https://kubernetes.io/blog/2020/03/25/kubernetes-1-18-release-announcement/)
- Sysdig on [what's in 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/)
- [CNCF Webinar](https://www.cncf.io/webinars/kubernetes-1-18/) from the release team!
- [Big announcement of Tanzu!](https://blogs.vmware.com/cloudnative/2020/03/10/become-a-modern-software-organization-with-vmware-tanzu/)
- This week on [Kubernetes Podcast from Google](https://kubernetespodcast.com/episode/095-etcd/) etcd with Xiang Li
- [Migrating from Helm 2 to Helm 3](https://geeksocket.in/posts/helm-2-3-migration/)
    - [Exploreing the security of helm](https://engineering.bitnami.com/articles/helm-security.html)
- From Digital Ocean [Kubernetes Tutorials: 13 in Spanish, 20 in Portuguese and 14 in Russian](https://www.digitalocean.com/community/tutorials?language=es&q=kubernetes)


### Core
- New versions of kubernetes are out with fixes for:
- CVE-2020-8551 [Kubelet DoS via API](https://github.com/kubernetes/kubernetes/issues/89377)
- CVE-2020-8552 [apiserver DoS oom](https://github.com/kubernetes/kubernetes/issues/89378)
- [Very good summary of the above by @ehashman](https://twitter.com/ehashdn/status/1242184473535377408?s=20)


## Show Notes


## Reference Links
* Cloud Events cloudevents.io
    * Primer - https://github.com/cloudevents/spec/blob/master/primer.md
    * Conformance tool - https://github.com/cloudevents/conformance
    * Cloud event viewer/sync - https://github.com/n3wscott/sockeye
* Argo Events links from @VaibhavPage
    * Installation - https://argoproj.github.io/argo-events/installation/
    * Intro - https://argoproj.github.io/argo-events/tutorials/01-introduction/
    * Demo - https://argoproj.github.io/argo-events/demo/notebooks/
    * Concepts - https://argoproj.github.io/argo-events/concepts/gateway/

## Getting macOS set up with k8s 1.18.0 with Kind
```
brew update
brew upgrade kind kubectl octant
kind create cluster --image kindest/node:v1.18.0
kubectl get version
```
