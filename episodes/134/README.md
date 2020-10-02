# Episode 134 : CSI
- Hosted by @joshrosso
- 10/02/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=llzH_we5guI
" target="_blank"><img src="http://img.youtube.com/vi/llzH_we5guI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://youtu.be/llzH_we5guI?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa)
- [00:04:19 - Week in Review](https://youtu.be/llzH_we5guI?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=259)
- [00:19:57 - Local Provisioning and Primitives](https://youtu.be/llzH_we5guI?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1197)
- [01:00:11 - Intro to CSI](https://youtu.be/llzH_we5guI?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=3611)
- [01:15:22 - Installing and Using a CSI Driver](https://youtu.be/llzH_we5guI?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=4522)
- [01:51:21 - Goodbye!](https://youtu.be/llzH_we5guI?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=6681)

## Agenda

* [x] Primitives via local provisioning
    * [x] PV/PVC
    * [x] local provisioner
* [x] CSI Introduction
    * [x] Architecture
    * [x] Deployment
* [x] Self-service storage
    * [x] StorageClasses
    * [x] Dynamic provisioning

## Week in Review

### Core K8s

- The [KubeCon + CloudNative Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/schedule/) is now live!
    - Congrats to all the speakers, and if you didn't get accepted good luck next time, keep trying, we've all been there!
- Another "mostly boring" week in core ahead of KubeCon!
- [Rootless mode enhancement](https://github.com/kubernetes/enhancements/issues/2033).  Great to see some motion on this! 
- [K8s 1.20 alpha 1 is out](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#v1200-alpha1)

### K8s and Cloud Native Ecosystem

* [KubeAcademy Pro: In-Depth Kubernetes Training, Totally Free](https://tanzu.vmware.com/content/blog/introducing-kubeacademy-pro-in-depth-kubernetes-training-totally-free)
    * [Joe Beda kicking off event #1 with an AMA](https://www.brighttalk.com/webcast/14883/441914)

- Your CTO might have just gotten back from VMWorld, [stay safe out there folks](https://i.redd.it/s1suyqft2kp51.png)!
- Techworld with Nana is doing a [video tutorial series on Prometheus](https://www.youtube.com/playlist?list=PLy7NrYWoggjxCF3av5JKwyG7FFF9eLeL4)
- Ibrahim Jelliti is [colleting CKSS exam preparation materials](https://github.com/ijelliti/CKSS-Certified-Kubernetes-Security-Specialist) - good luck to you pursuing this!
- Jacob Beasely investigaed [Deploying SQL Databases on Kubernetes](https://www.intelletive.com/2020/08/12/blog-best-practices-when-deploying-sql-databases-on-kubernetes/)
- Didier Durand has set up a nice [microservices on K8s demo](https://github.com/didier-durand/microservices-on-cloud-kubernetes) - "It is composed of 10 polyglot microservices behind a nice-looking web frontend calling them to serve client requests. A load-generator - part of the package - will generate traffic while the application is running to make use of tools (Prometheus, OpenTelemetry, etc.) more attractive." This looks great for demos!
- [Percona k8s Operator with OpenEBS](https://www.percona.com/blog/2020/10/01/deploying-percona-kubernetes-operators-with-openebs-local-storage)

## Show Notes

- [Using Kubernetes Local Storage for Scale-Out Storage Services - Michelle Au & Ian Chakeres](https://youtu.be/eqkgiPppZN8)


