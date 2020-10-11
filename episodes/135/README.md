# Episode 135: Antrea CNI

- Hosted by @jbeda
- Recording date: 2020-10-09 

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=py6gfeDuA30" target="_blank"><img src="https://i.ytimg.com/vi/py6gfeDuA30/maxresdefault.jpg" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:46 - Week in Review
- 00:17:06 - Starting with Antrea
- 00:18:26 - CNI vs. Service Mesh
- 00:23:04 - Why Antrea?
- 00:28:01 - Launching a cluster
- 00:41:00 - Installing Antrea
- 00:49:59 - Installing a sample app
- 00:54:38 - Feature gates
- 01:04:27 - Debugging bad config
- 01:10:25 - Setting up Traceflow
- 01:27:02 - Prometheus exports
- 01:35:30 - Antrea Architecture
- 01:41:34 - `route -n` in a node
- 01:48:28 - Wrap up and Thank You!


## Week in review

### Core

Not much going on in core this week, these are some great KEPs for 1.20: 
- [kubectl debug](https://github.com/kubernetes/enhancements/issues/1441)
- [Graceful node shutdown](https://github.com/kubernetes/enhancements/issues/2000)
- [Container Resource based Pod Autoscaling](https://github.com/kubernetes/enhancements/issues/1610)
- [Deprecate and remove SelfLink](https://github.com/kubernetes/enhancements/issues/1164)
- [Rename the kubeadm master node-role label and taint](https://github.com/kubernetes/enhancements/issues/2067)

### Cloud Native Ecosystem
* A bunch of cloud engineering stuff happened at the [cloud engineering summit hosted by pulumi!](https://cloudengineering.heysummit.com/) the videos are free to watch check out some of these talks!
* [@mauilion is presenting a kubecon AMA keynote with Ian Coldwater, Brad Geesaman and Rory McCune this year and need your questions!](https://discuss.kubernetes.io/t/call-for-questions-sig-honk-ama-kubecon-na-keynote-panel/13159) 
* [Rook graduates the cncf!](https://www.cncf.io/announcements/2020/10/07/cloud-native-computing-foundation-announces-rook-graduation/) It's the 13th project and the first block,file or object store to graduate.
* [Kube DOOM](https://github.com/storax/kubedoom)
* Heads up Helm users, important [changes to the Helm Repos](https://www.cncf.io/blog/2020/10/07/important-reminder-for-all-helm-users-stable-incubator-repos-are-deprecated-and-all-images-are-changing-location/)!
* [Contour 1.9 is released](https://projectcontour.io/contour_v190/)! Lots of great features.

## Show Notes

### Build AWS cluster

This is a simple test cluster. AWS specific cloud provider is not used so we won't have ELB or EBS integration.

Pre built AMIs: 
* https://github.com/kubernetes-sigs/cluster-api-provider-aws/blob/master/docs/amis.md
* Built via https://github.com/kubernetes-sigs/image-builder

Created a set of 3 nodes in EC2 using a launch template
* Single security group that allows external SSH and all network access on VPC
* t3a.large
* 30GB EBS root volume, gp2

Name each node in console and copy/paste IPs to `config.sh`.

Launch on control-plane
```
POD_NETWORK=10.20.0.0/16
kubeadm init --pod-network-cidr=$POD_NETWORK
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

Join on each client node. Something like this.
```
kubeadm join 172.31.43.148:6443 --token tuc8cn.4kg46l7bpybcs2fz \
    --discovery-token-ca-cert-hash sha256:0559388cee70283b8d483620d4dfb2d5bb704ce566de146d92c9f82d05c43e9f
```

### Install Antrea. 
[Directions here](https://github.com/vmware-tanzu/antrea/blob/master/docs/getting-started.md#installation)
```
kubectl apply -f https://github.com/vmware-tanzu/antrea/releases/download/v0.10.1/antrea.yml
```

TODO: Explore [feature gates](https://github.com/vmware-tanzu/antrea/blob/master/docs/feature-gates.md)

Download/use [antctl](https://github.com/vmware-tanzu/antrea/blob/master/docs/antctl.md)

Install [Octant Plugin](https://github.com/vmware-tanzu/antrea/blob/master/docs/octant-plugin-installation.md)

### Demo App
Demo app: [Emojivoto](https://github.com/BuoyantIO/emojivoto) from Boyant folks 
```
kubectl apply -k github.com/BuoyantIO/emojivoto/kustomize/deployment
```

### Explore
Details in https://github.com/vmware-tanzu/tgik/issues/326.

From @jayunit100:
* Try out `route -n` on node to see the "gateways" that get installed.
* Set up prometheus here. `kubectl apply -f build/yamls/antrea-prometheus.yml`.  Need to enable prom merics though.
* Set up network policies and see them with antctl.
* CNI debugging doc: https://github.com/jayunit100/k8sprototypes/blob/master/2020kubecon/cni_debugging.md
* Metrics intro: https://github.com/jayunit100/k8sprototypes/tree/master/antrea_metrics

From @antoninbas:
* Use Traceflow (packet tracing in the Pod network) along with NetworkPolicies to show traffic being accepted or dropped. antctl query endpoint can be demo'd simultaneously to show policies applied to specific Pods and aggregated NetworkPolicy metrics can be show using vanilla kubectl.
* Enable flow export and show flow information using Kibana dashboards (https://github.com/vmware-tanzu/antrea/blob/master/docs/network-flow-visibility.md). This requires a "beefier" cluster to run the ELK stack, but we provide the manifests.

From @jianjuns:
* Traceflow guide: https://github.com/vmware-tanzu/antrea/blob/master/docs/traceflow-guide.md
* If we want to go further to dump OVS config, flows, and Antrea internal state, these two docs list some useful commands:
https://github.com/vmware-tanzu/antrea/blob/master/docs/antctl.md
https://github.com/vmware-tanzu/antrea/blob/master/docs/troubleshooting.md

## Reference Links
* https://antrea.io/
* [Antrea Docs](https://antrea.io/docs/)
* [Antrea Arch Doc](https://github.com/vmware-tanzu/antrea/blob/master/docs/architecture.md)
* [Antrea Github repo](https://github.com/vmware-tanzu/antrea)
* [Antrea on kind](https://github.com/vmware-tanzu/antrea/blob/master/docs/kind.md)