# Episode 108 : Cluster-API for Docker!

- Hosted by @mauilion
- 03/06/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=6pFW6h6AORQ
" target="_blank"><img src="http://img.youtube.com/vi/6pFW6h6AORQ/hqdefault.jpg" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core Kubernetes

- We are in [1.18 beta1](https://github.com/kubernetes/kubernetes/blob/v1.18.0-beta.1/CHANGELOG/CHANGELOG-1.18.md#v1180-beta1)
- Sidecars [bumped to 1.19](https://github.com/kubernetes/enhancements/issues/753)
    - Here's the [KEP for Sidecars](https://github.com/kubernetes/enhancements/blob/e0fc0e7eab51078b5b7bd0730cf48f550bc91d1e/keps/sig-apps/sidecarcontainers.md)
- Getting [rid of bazel](https://github.com/kubernetes/kubernetes/issues/88553)?
- Moving the [official images to a community controlled resource](https://github.com/kubernetes/release/issues/911) is on track for 4/1
    - [Full checklist for k8s-infra here](https://github.com/kubernetes/release/issues/911)
- KubeCon + CloudNativeCon [has been postponed](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/attend/novel-coronavirus-update/)
    - Contributor Summit [also postponed](https://kubernetes.io/blog/2020/03/04/contributor-summit-delayed/)
    - Keep an eye out on blog.k8s.io for the latest updates

### Ecosystem

- [The Kui Framework for Graphical Terminals](https://github.com/IBM/kui)
- [Everyone might be a cluster-admin in your Kubernetes cluster](https://www.jeffgeerling.com/blog/2020/everyone-might-be-cluster-admin-your-kubernetes-cluster)
- Kubeflow has [gone 1.0](https://medium.com/kubeflow/kubeflow-1-0-cloud-native-ml-for-everyone-a3950202751) - congrats to that team!
    - Speaking of, Spotify has [OSSed their terraform modules](https://github.com/spotify/terraform-gke-kubeflow-cluster) for setting up a kubeflow cluster in GKE
- [Vault replication across multiple datacenters on Kubernetes ](https://banzaicloud.com/blog/vault-multi-datacenter/)
- Techworld with Nana explains [namespaces in 15 minutes](https://www.youtube.com/watch?v=K3jNo4z5Jx8&feature=youtu.be). [Video]
- Have CRDs [killed the free control plane](https://caleblloyd.com/software/crds-killed-free-kubernetes-control-plane/)?
- All the Tanzu tshirts for the episode 100 contest have been delivered, except for Sergey Yagodkin's because our vendor can't ship to Russia. :(
    - Anyone have a good idea to show Sergey our gratitude? Leave a comment in chat!
- Google GKE is now [charging to use their control plane](https://cloud.google.com/kubernetes-engine/pricing)


## Show Notes

- [Cluster API Upgrade Tool](https://github.com/vmware/cluster-api-upgrade-tool)

## Reference Links

- [CAPI Glossary](https://cluster-api.sigs.k8s.io/reference/glossary.html)
- Duffie being interviewed on [kubecuddle](https://twitter.com/richburroughs/status/1235977388812627968?s=20)