# Episode 097 : Cluster Autoscaler

- Hosted by @joshrosso
- Recording date: 2019-11-08

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=NY7pyRNrHzE
" target="_blank"><img src="http://img.youtube.com/vi/NY7pyRNrHzE/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:04 - Week in Review
- 00:18:40 - Autoscaling Poll
- 00:23:13 - Autoscaling Overview
- 00:42:00 - Deployment Resource and Limits
- 01:00:00 - Horizontal Pod Autoscaler Overview and Configuration
- 01:12:49 - Cluster Autoscaler Overview and Deployment
- 01:38:00 - Autoscaler Implications for Cluster API
- 01:46:00 - Quick Overview of Vertical Pod Autoscaler

## Week in Review

- k8s 1.17 news
    - Beta [release is live](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.17.md/#v1170-beta1)
    - Code freeze November 14!

Cool k8s Links from around the web
    - [ChaosKube](https://github.com/linki/chaoskube)
    - [godaddy has OSSed their gated deployments tool](https://github.com/godaddy/kubernetes-gated-deployments)
    - [Java in a Container Resource Guidelines](https://www.ccampo.me/java/docker/containers/kubernetes/2019/10/31/java-in-a-container.html)
    - [Contour's gone 1.0](https://projectcontour.io/announcing-contour-1.0/)
    - [skaffold is now GA](https://cloud.google.com/blog/products/application-development/kubernetes-development-simplified-skaffold-is-now-ga)
    - [Handle Namespace deletion more gracefully in built-in controllers](https://github.com/kubernetes/kubernetes/pull/84123)
    

## Show Notes

* [x] The world of "scaling"
* [x] Resource scheduling
* [x] Horizontal Pod Autoscaler (HPA)
* [x] Vertical Pod Autoscaler (VPA)
* [x] Cluster Autoscaler (CA)
* [x] Bringing it together

## Reference Links

* [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)
* [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
* [Vertical Pod Autosclaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)
