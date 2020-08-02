# Episode 126 : Vertical Pod Autoscaling

- Hosted by @jbeda
- 07/31/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=p6zCzyBNhLE
" target="_blank"><img src="http://img.youtube.com/vi/p6zCzyBNhLE/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:02:50 - Week in Review
- 00:20:17 - Autoscaling overview
- 00:31:33 - Metrics Server
- 00:32:08 - Aside on Java and VPC
- 00:33:25 - Back to Metrics Server
- 00:43:40 - Fixing Metrics Server on Kind
- 00:45:40 - Aside on kapp
- 00:48:27 - Back to Metrics Server
- 00:52:03 - Installing Vertical Pod Autoscaler
- 01:01:58 - Activating VPA for a deployment
- 01:14:22 - Description of automatic update flow
- 01:25:48 - Exploring OOM reaction
- 01:41:02 - Installing/exploring Goldilocks
- 01:55:22 - Wrapping up! Thank you!


## Week in Review

### Core k8s

- k8s [1.19rc3 is the latest](https://github.com/kubernetes/kubernetes/releases) of the 1.19 series.
    - [Changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md/#v1190-rc3)
    - 1.19 will be built on [go 1.15](https://groups.google.com/forum/#!topic/kubernetes-dev/uN-hX5zUMFA)
- SIG Multicluster [has settled](https://groups.google.com/forum/#!topic/kubernetes-sig-multicluster/X9IhpnujX7o) on `ClusterSet` as the proper name for a set of clusters. Other popular nominations included  ClusterPool, ClusterFleet, and ClusterArray. 
- Code of Conduct Committee [election process is underway](https://groups.google.com/forum/#!topic/kubernetes-dev/AUHqIC3B0jQ)


### Things from around the ecosystem

- Monica Bhartiya has put together [a summary of changes to the CKA exam](https://monicabhartiya.com/posts/kubernetes-syllabus-update) - great diagram there!
- [I made a Kubernetes game where you explore your cluster and destroy pods](https://www.shogan.co.uk/kubernetes/i-made-a-kubernetes-game-where-you-explore-your-cluster-and-destroy-pods/)
- Jeff Geerling's final thoughts on the [Turing Pi](https://www.youtube.com/watch?v=aApByQWqnV0) with a Kubernetes Cluster. 
- [kVDI](https://github.com/tinyzimmer/kvdi) - A Kubernetes-native Virtual Desktop Infrastructure.
- Chaos Mesh is now a [CNCF Sandbox project](https://chaos-mesh.org/blog/chaos-mesh-join-cncf-sandbox-project/)
- Ryan Miller goes over [Writing a CRD and Controller with Kubebuilder](https://rvmiller.com/2020/07/04/tutorial-writing-a-kubernetes-crd-and-controller-with-kubebuilder/)
- Bryant Hagadorn talks about [Two Quick Ways to Apply Zero Trust in Kubernetes](https://itnext.io/two-quick-ways-to-apply-zero-trust-to-kubernetes-79764dd420bf)


## Show Notes

* Main [VPA github page](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)
* Kubernetes [Metrics Server](https://github.com/kubernetes-sigs/metrics-server)
    * [Hints from Duffie](https://gist.github.com/mauilion/cd10204e924600c15c5c14c97c2ca66b#file-components-yaml-L91-L92) on running Metrics Server with kind. Hostnames of the nodes aren't resolvable from the metrics-server pod.
and because the serving certs for the kubelet are self signed.
* Cool dashboard for recommendations: [Goldilocks from Fairwinds](https://github.com/FairwindsOps/goldilocks)
* Kubernetes [office hours segment on VPA](https://www.youtube.com/watch?v=5kJ6tJXq-qU&feature=youtu.be&t=2558)
* Google Autopilot
    * [Recent paper](https://dl.acm.org/doi/pdf/10.1145/3342195.3387524)
    * [Analysis](https://twitter.com/embano1/status/1257779994664861696) from Michael Gasch
    * [Some highlights](https://twitter.com/copyconstruct/status/1253966160678096896) from Cindy Sridharan
* SIG Autoscaling
    * [Meeting and community information](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md)
    * [Introduction to Autoscaling - Marcin Wielgus, Google & Vivek Bagade, Google](https://www.youtube.com/watch?v=S_VlztrihzU) - KubeCon + CloudNativeCon talk
* [Definitions used in recommender](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler#recommendedcontainerresources_v1_autoscalingk8sio)
* [k14s](https://k14s.io/) along with [kapp](https://get-kapp.io/). TGIK on YTT and KAP: [tgik.io/079](https://tgik.io/079)
* Client side Kubernetes dashboard: [Octant](https://octant.dev/). TGIK on Octant: [tgik.io/087](https://tgik.io/087)
    

