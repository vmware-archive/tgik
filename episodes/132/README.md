# Episode 132 : Sealed Secrets

- Hosted by @jbeda
- 09/18/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?x-cDk8DIvwE
" target="_blank"><img src="http://img.youtube.com/vi/x-cDk8DIvwE/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:31 - Week in Review
- 00:23:15 - Sealed Secret Overview & Install
- 00:37:21 - Sealing a Secret
- 00:52:49 - Retrieving the Public Key
- 00:55:28 - Exploring SealedSecret Scope
- 01:00:00 - Rotation: Secrets and Keys
- 01:15:37 - Multi-Cluster Sealed Secrets
- 01:30:28 - Wrap-up

## Week in Review

### Core

- 1.20 has begun, check out this [status update from release manager Jeremy Rickard](https://groups.google.com/g/kubernetes-dev/c/-ErnMdUrDIE) for all the important dates.
    - Here's the [release page](https://www.kubernetes.dev/resources/release/) on k8s.dev will all the info 
- [CVE 2020-14386](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-14386)
    - [Sysdig has detection/mitigation tips](https://sysdig.com/blog/cve-2020-14386-falco/)
    - [GKE's security bulletin](https://cloud.google.com/kubernetes-engine/docs/security-bulletins#gcp-2020-012)
- [1.19.2 point release](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#changelog-since-v1191) 
- [Warnings are a thing now in K8s](https://kubernetes.io/blog/2020/09/03/warnings/)
    - [KEP 1693](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1693-warnings)
    - Jordan Liggit's [live demo](https://www.youtube.com/watch?v=6Wn_dIEg0E8)

### Cloud Native Ecosystem

- [What gitlab.com learned from running k8s for a year](https://about.gitlab.com/blog/2020/09/16/year-of-kubernetes/)
- [Kruster](https://cruster.io/learn/demo/) - clusters on rPi, I know there's tons of these but your author has a soft spot for all tools using kubeadm. :D
- [Helm 3.3.2/3.3.3 Release](https://github.com/helm/helm/releases/tag/v3.3.2)
    - [Upgrading to Helm 3 with Flux](https://itnext.io/upgrading-to-helm-3-with-flux-cd-6b7014223a51)
- [Kubernetes CNI Benchmark (Updated: August 2020)](https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-updated-august-2020-6e1b757b9e49)


### Sealed Secrets Notes

- [GitOps for secrets using Sealed Secrets](https://github.com/swade1987/gitops-with-secrets)

## Show Notes

* [Sealed Secrets Github](https://github.com/bitnami-labs/sealed-secrets)
* [Theory 11 cards](store.theory11.com) - I *love* the star wars cards.

## Alternative Secret Management

- [Managing Kubernetes Secrets Securely with Gitops (with Flux and SOPS)](https://itnext.io/managing-kubernetes-secrets-securely-with-gitops-b8174b4f4d30)