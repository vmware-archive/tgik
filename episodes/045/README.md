# Episode 045 : Calico (CNI)

- Hosted by @kris-nova
- Recording date

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url
Note the 0.jpg for the thumbnail --->

<a href="http://www.youtube.com/watch?feature=player_embedded&v=VKtudjLc1oc
" target="_blank"><img src="http://img.youtube.com/vi/VKtudjLc1oc/0.jpg" 
alt="IMAGE ALT TEXT HERE" width="640" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Show Notes

Deploy a cluster without Calico using Kubciron

```bash

# From this directory
kubicorn create tgik-calico-cluster -S kubicorn -p aws -M serverPool.bootstrapScripts[0]=kubicorn/amazon_k8s_ubuntu_16.04_master.sh
kubicorn apply tgik-calico-cluster -S kubicorn

```

## Hacks for the episode

Start the episode off trying to deploy a pod, so we can start debugging

```bash
kubectl run nginx --image nginx
```

List all iptables rules

```bash
iptables -vL -t filter
iptables -vL -t nat
iptables -vL -t mangle
iptables -vL -t raw
iptables -vL -t security

# or

alias iptables-list-all="iptables -vL -t filter && iptables -vL -t nat && iptables -vL -t mangle && iptables -vL -t raw && iptables -vL -t security"
```

Show the route table

```bash
route
netstat -tlpn
ip route list
```

Show the network interfaces

```bash
ifconfig
```

Show cgroups

```bash
cat /proc/cgroups
```

Show the kubelet journal

```bash
journalctl -fu kubelet
```

Inspect docker containers

```bash
docker ps -a
docker inspect <id> | jq
```

Cat kubelet unit files

```bash
cat /etc/systemd/system/kubelet.service.d/*
```


## Reference Links

 - Kubernetes CNI Plugin documentation https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/#cni
 - Kubicorn bootstrap scripts https://github.com/kubicorn/bootstrap
 - Examples of configuring a cluster bootstrap script https://github.com/kubicorn/kubicorn/tree/master/examples/pipeline
 - Network policies and isolation https://kubernetes.io/docs/concepts/services-networking/network-policies/
 - Configuring calicoctl https://docs.projectcalico.org/v3.1/usage/calicoctl/configure/
 - Running calicoctl in a pod https://docs.projectcalico.org/v3.1/usage/calicoctl/install