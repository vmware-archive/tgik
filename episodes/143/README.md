# Episode 143 : Cluster API Update

- Hosted by [@scott_lowe](https://twitter.com/scott_Lowe)
- 01/22/2021

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=AeHfVFepsMg
" target="_blank"><img src="http://img.youtube.com/vi/AeHfVFepsMg/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:03:20 - Week in Review
- 00:11:20 - ClusterResourceSet
- 00:12:05 - `ktx` sidebar
- 00:12:35 - Back to ClusterResourceSet
- 00:33:05 - KubeadmControlPlane
- 00:40:34 - `clusterctl get kubeconfig` sidebar
- 00:42:16 - MachineHealthCheck

## Week in Review

### Core K8s

* New version of Cluster API (0.3.13) [released](https://github.com/kubernetes-sigs/cluster-api/releases/v0.3.13)
* New version of Cluster API Provider for AWS (0.6.4) [released](https://github.com/kubernetes-sigs/cluster-api-provider-aws/releases/v0.6.4)

### K8s and Cloud Native Ecosystem

* [Helm course](https://kube.academy/courses/helm-101) recently released on KubeAcademy
* New book _Production Kubernetes_ recently went to production ([O'Reilly site](https://www.oreilly.com/library/view/production-kubernetes/9781492092292/), [@joshrosso's tweet](https://twitter.com/joshrosso/status/1351932207108136965))
* [Using Harbor as a Docker Hub Proxy](https://tanzu.vmware.com/developer/guides/kubernetes/harbor-as-docker-proxy/): Stop fighting dockerhub rate limits, also save bandwidth for your home lab.
* Blog post: [etcd Deep Dive](https://www.mgasch.com/2021/01/listwatch-part-1/) (Series on ListWatch pattern in Kubernetes)
* CRD Docs Website: [CAPV example](https://doc.crds.dev/github.com/kubernetes-sigs/cluster-api-provider-vsphere/infrastructure.cluster.x-k8s.io/VSphereMachine/v1alpha3@v0.7.4)

## Show Notes

* [Cluster API Book](https://cluster-api.sigs.k8s.io)
* [ClusterResourceSet](https://cluster-api.sigs.k8s.io/tasks/experimental-features/cluster-resource-set.html)
* [CAEP for ClusterResourceSet](https://github.com/kubernetes-sigs/cluster-api/blob/master/docs/proposals/20200220-cluster-resource-set.md)
* [Using ClusterResourceSet to install Calico CNI](https://samperrin.com/posts/cluster-api-vsphere-using-clusterresourceset-to-install-calico-cni/)
* [MachineHealthCheck](https://cluster-api.sigs.k8s.io/tasks/healthcheck.html)
* [`clusterctl` commands](https://cluster-api.sigs.k8s.io/clusterctl/commands/commands.html)
* [`ktx`](https://github.com/heptiolabs/ktx) (CLI tool used to manage Kubeconfig contexts)
