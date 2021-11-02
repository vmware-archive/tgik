# Episode 159: Exploring Kuberhealthy

- Hosted by @jbeda
- Recording date: 2021-07-09

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=4w62d0kL7Mw" target="_blank"><img src="https://i.ytimg.com/vi/4w62d0kL7Mw/maxresdefault.jpg" border="10" /></a>

## Table of Contents

- 00:04:47 - Welcome to TGIK!
- 00:07:20 - Week in Review
- 00:25:28 - Looking at Kuberhealthy
- 00:28:59 - Installing Kuberhealthy
- 00:43:27 - Connecting to console
- 00:50:26 - Finding prom metrics (`/metrics')
- 00:55:15 - Exploring external checks
- 01:01:11 - Creating your own checks
- 01:03:28 - Continuing to debug cross-namespace checks
- 01:10:12 - Looking at HTTP Check
- 01:13:16 - Debugging and filing bug
- 01:48:22 - Wrapping up

## Week in Review

* [K8s 1.22 code freeze](https://groups.google.com/g/kubernetes-dev/c/YmeQrrrqmqM/m/FpMInN1yAwAJ)
* [Contour 1.17 released](https://twitter.com/projectcontour/status/1412685559491973120)
* [Peritus.ai](https://peritus.ai/) on [discuss.kubernetes.io](https://discuss.kubernetes.io/).
* Blog post on [Kubernetes Schema validation](https://opensource.com/article/21/7/kubernetes-schema-validation).
* Cool looking time based scaler project: [another-scheduler](https://github.com/dignajar/another-scheduler).
  * Similar to [Keda cron schedule](https://keda.sh/docs/2.3/scalers/cron/).

## Main Topic Notes
* [Talk from KubeCon NA 2019](https://www.youtube.com/watch?v=aAJlWhBtzqY)
* Blog post: [K8s KPIs with Kuberhealthy](https://kubernetes.io/blog/2020/05/29/k8s-kpis-with-kuberhealthy/)
* https://github.com/kuberhealthy/kuberhealthy

* Bug we found: https://github.com/kuberhealthy/kuberhealthy/issues/975. Looks like it is already fixed but not released yet.