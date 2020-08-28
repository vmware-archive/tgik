# Episode 049 : Flannel (CNI)

- Hosted by @krisnova
- Recording date: 2018-08-31
- HackMD: https://hackmd.io/Vd_cF06PRvSovEoMvn2Riw?both

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url
Note the 0.jpg for the thumbnail --->

<a href="http://www.youtube.com/watch?feature=player_embedded&v=2YoK4bBy3CM
" target="_blank"><img src="http://img.youtube.com/vi/2YoK4bBy3CM/0.jpg"
alt="IMAGE ALT TEXT HERE" width="640" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:07:00 - Week in Review
- 00:29:00 - Introduction - Looking at Flannel
- 00:31:15 - How to turn on --pod-network-cidr
- 00:39:57 - Deploying a pod that doesn't work
- 00:42:47 - Installing Flannel
- 00:51:50 - Looking at various types of Flannel backends
- 00:58:28 - Trick to run CNI Linux-compiled binaries on a Mac
- 01:00:07 - Looking at the Flannel source code
- 01:12:25 - Observing iptables and route tables mutations
- 01:16:48 - Looking at how Calico, Flannel and Canal fit together
- 01:25:05 - Back to iptables and route tables mutations
- 01:30:35 - Wrapping up

## Show Notes

Deploy a cluster without Calico using Kubciron

```bash

# From this directory
kubicorn create tgik-flannel-cluster -S kubicorn -p aws -M serverPool.bootstrapScripts[0]=kubicorn/amazon_k8s_ubuntu_16.04_master.sh
kubicorn apply tgik-flannel-cluster -S kubicorn

```

#### Using kubeadm without --pod-network-cidr

```
knova:049 nova$ k logs pod/kube-flannel-ds-amd64-qpv98 -n kube-system -f
I0831 15:41:20.041778       1 main.go:475] Determining IP address of default interface
I0831 15:41:20.042070       1 main.go:488] Using interface with name eth0 and address 10.0.100.70
I0831 15:41:20.042091       1 main.go:505] Defaulting external address to interface address (10.0.100.70)
I0831 15:41:20.053126       1 kube.go:131] Waiting 10m0s for node controller to sync
I0831 15:41:20.053163       1 kube.go:294] Starting kube subnet manager
I0831 15:41:21.053341       1 kube.go:138] Node controller sync successful
I0831 15:41:21.053367       1 main.go:235] Created subnet manager: Kubernetes Subnet Manager - ip-10-0-100-70.us-west-2.compute.internal
I0831 15:41:21.053373       1 main.go:238] Installing signal handlers
I0831 15:41:21.053419       1 main.go:353] Found network config - Backend type: vxlan
I0831 15:41:21.053476       1 vxlan.go:120] VXLAN config: VNI=1 Port=0 GBP=false DirectRouting=false
E0831 15:41:21.053688       1 main.go:280] Error registering network: failed to acquire lease: node "ip-10-0-100-70.us-west-2.compute.internal" pod cidr not assigned
I0831 15:41:21.053722       1 main.go:333] Stopping shutdownHandler...
```

#### Using kubeadm with --pod-network-cidr


Documentation on setting up flannel with `kubeadm`
[Kubeadm --pod-network-cidr](https://coreos.com/flannel/docs/latest/kubernetes.html#the-flannel-cni-plugin)

Documentation on declaring a pod network cidr with `kubeadm`
[Pod Network Add-On](https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/#pod-network)

Turns out it's actually a completely different name in the config file
[Kubeadm API docs](https://godoc.org/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm#MasterConfiguration)
[Actual struct](https://github.com/kubernetes/kubernetes/blob/master/cmd/kubeadm/app/apis/kubeadm/types.go#L197-L204)

```
type Networking struct {
	// ServiceSubnet is the subnet used by k8s services. Defaults to "10.96.0.0/12".
	ServiceSubnet string
	// PodSubnet is the subnet used by pods.
	PodSubnet string
	// DNSDomain is the dns domain used by k8s services. Defaults to "cluster.local".
	DNSDomain string
}
```



## Exploring flannel

Installing flannel

https://github.com/coreos/flannel#deploying-flannel-manually

```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```


Exec into a flannel pod on the master

```
kubectl exec -it <pod> sh
```

Show route table

```
ip route show
```

Build and run flanneld locally

```
docker run -it -v /Users/nova/go/src/github.com/coreos/flannel/:/go/src/github.com/coreos/flannel golang:latest /bin/bash
cd /go/src/github.com/coreos/flannel
go run main.go
# etcd errors
```

Build and run the CNI plugin locally

```
docker run -it -v /Users/nova/go/src/github.com/containernetworking/plugins/:/go/src/github.com/containernetworking/plugins golang:latest /bin/bash
cd /go/src/github.com/containernetworking/plugins
sh build.sh
```


## Reference Links


 - [Go IPTables for Flannel](https://github.com/coreos?utf8=%E2%9C%93&q=iptables&type=&language=)
 - [Found this looking for flannel docs, not the best](https://chrislovecnm.com/kubernetes/cni/choosing-a-cni-provider/)
 - [Fantastic resource](https://github.com/containernetworking/cni#3rd-party-plugins)
 - [Release for Ubuntu Xenial](https://github.com/kubernetes/release/tree/master/debian/xenial)


## Duffie notes

People complain about VXLan because of multicast, but flannel knows the answers from the Node object.
Flannel doesn't need to use multicast.
Flannel only uses VXLan for encapsulation.
All flat network - Think of it like base64
Security might be a concern here.
VXLan TCP Port: 4789

This is really interesting: https://github.com/projectcalico/canal

Canal is basically a wrapper that wires up Calico and Flannel at the same time: https://github.com/projectcalico/canal#policy-based-networking-for-cloud-native-applications

