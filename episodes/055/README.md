# Episode 055 : Creating container images with Kaniko

- Hosted by @krisnova
- Recording date: 2018-10-26

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=rt0fqOekgjI
" target="_blank"><img src="http://img.youtube.com/vi/rt0fqOekgjI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:12 - Week in Review
- 00:28:00 - Getting into Kaniko
- 00:30:00 - Kaniko Build Context
- 00:36:00 - Running Kaniko Locally.
- 00:50:50 - Local Container Built :)
- 00:58:00 - Running Kaniko in Kubernetes

## Stuff from around the Community
[00:05] - [00:25]

- Kubernetes 1.12.2 is out, here's the [changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.12.md#v1122).
- Weaveflux [1.6](https://discuss.kubernetes.io/t/weave-flux-1-6-0-released/2574) and [1.8](https://discuss.kubernetes.io/t/weave-flux-1-8-0-is-out/3230) are out!
- Heptio Ark v.0.9.9 [bugfix release](https://discuss.kubernetes.io/t/heptio-ark-v0-9-9-bug-fix-release/3222) is out!
- The Datadog folks share how they [run k8s clusters the very hard way](https://www.youtube.com/watch?v=2dsCwp_j0yQ) [video]
- Michael Hausenblaus has a great blog entry on [applied Kubernetes security](https://medium.com/@mhausenblas/10-2018-cncf-meetup-on-applied-kubernetes-security-7293f3f9471f) 
- [What the hell is a pod anyways?](https://medium.com/@dominik.tornow/what-the-hell-is-a-pod-anyways-72e5534b892c) by Andrew Chen and Dominik Tornow
    - Related: https://github.com/jamiehannaford/what-happens-when-k8s
- [Moving Canary deployments on AWS using ELB to kubernetes using Traefik](https://tasdikrahman.me/2018/10/25/canary-deployments-on-AWS-and-kubernetes-using-traefik/)
- kubernetes/features is now [kubernetes/enhacements](https://github.com/kubernetes/enhancements)
- [Kubeadm GA](https://github.com/kubernetes/kubeadm/milestone/9)
- [Kris Nova Transgender Picture from All Things Open](https://www.instagram.com/p/BpSRR9IAWaa/)
- [Unprivledged Container Builds](https://kinvolk.io/blog/2018/04/towards-unprivileged-container-builds/)
- [Ignite talk - What is Kubernetes?](https://youtu.be/oXKVZvDRLIc?t=40) - Hand drawn picture of Joe Beda!
- [00:28] - The Kubelets: New podcast from Kris and Amy Chen on navigating the cloud native landscape, stay tuned for more information!

## Show Notes - Getting into Kaniko
[00:28] 

Getting Kaniko: https://github.com/GoogleContainerTools/kaniko

[00:33] - Working with a simple Dockerfile
[00:35] - Multiple contexts
[00:36] - Running Kaniko Locally (in a container)


 
#### Creating a build context

```bash
tar -C ./context -zcvf context.tar.gz .
gsutil cp context.tar.gz gs://krisnova-tgik
```

## Reference Links

 - [secure metrics and monitoring with spiffe](https://blog.scytale.io/secure-application-metrics-distributed-logging-with-spiffe-f54f9f798124)
 - [Pod Lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)
 - [Link to running Kaniko in Kubernetes](https://github.com/GoogleContainerTools/kaniko#running-kaniko-in-a-kubernetes-cluster)
 - [Kaniko vs Others](https://github.com/GoogleContainerTools/kaniko/blob/master/README.md#comparison-with-other-tools)

