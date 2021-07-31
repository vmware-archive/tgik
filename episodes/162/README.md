# Episode 162 : CNI code surf !!!

- Hosted by @jayunit100
- Recording date: 2021-07-30

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url

<a href=https://www.youtube.com/watch?v=l3TWbrWkVzY"></aref>
-->

This show we'll look at how the codebases for calico, antrea are structures, and how they are designed to work/run in different environments.

Show notes: https://hackmd.io/ECpQYlYCQX6h5w9yB02RHw
Scrib: https://scribbletogether.com/whiteboard/B6549CE8-5E7A-4A11-8F14-6A9920837131


## Week in Review

- [kapp-controller: a new package manager for Kubernetes](https://carvel.dev/blog/introduction-to-carvel-package-manager-for-kubernetes/)
- [Contour v1.18.0 release](https://github.com/projectcontour/contour/releases/tag/v1.18.0)
- [Kpng iptables implementation merging ](https://github.com/kubernetes-sigs/kpng/pull/44/files)
- [Ingress NGINX v1.0.0-beta.1 release](https://github.com/kubernetes/ingress-nginx/releases/tag/controller-v1.0.0-beta.1) - WARNING! BREAKING CHANGES
- Linkerd is now a graduated CNCF project - https://linkerd.io/2021/07/28/announcing-cncf-graduation/
- Kubernetes Release team [shadow](https://github.com/kubernetes/sig-release/blob/master/release-team/shadows.md) application is open - [Form](https://forms.gle/7As7hacvMhxBQaox8)

## Table of Contents

- utils https://github.com/jayunit100/k8sprototypes/
  - kind/kind-local-up.sh script to create cnis
  - `CLUSTER=calico CONFIG=kind-conf.yaml ./kind-local-up.sh` example usage
  - `CLUSTER=antrea CONFIG=kind-conf.yaml ./kind-local-up.sh`
- libcalico.go
- kube-controllers: watches apiserver, converts netpol objects into native calico objects as needed, node-controller, endpoint-controller, ....
- calico controller independent of the agent
- operator:
    - images + yaml match ups
    - typha management
- kovacs: route reflector logic
- calico 3.20 is on the way out !!!
- security holes in k8s
    - hostnetwork pods
    - services
-
## Show Notes

- diagram
- cillium ? kernel 4.17 not 4.15

- 1:00 - introduction
- 4:00 - kpng,
- 22:00 - oops need helm3 for cilium
- 25:00 - cilium installation, ebpf and xdp
- 30:00 - cilium ~ needs Kernel 4.17 !
- 34:00 - tigera operator, controllers
- 38:00 - calico controller vs cillium controller
- 40:00 - calico node and felix
- 44:00 - calico apiserver and lib-calico
- 50:00 - conversion.go and lib-calico
- 52:00 - casey davenport arrives !
- 53:00 - the calico controller - non-critical
- 54:00 - calico 3.20 on the way, iptables + ebpf modes, and how ipam is managed
- 1:10:00 - trying out multus
- 1:13:00 - /etc/cni and containerd CNI Add
- 1:20:00 - trying to install flannel / antrea as secondary CNIs
- 1:30:00 - intro to antrea
- 1:34:00 - calico-node vs antrea node on windows
