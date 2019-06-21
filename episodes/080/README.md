# Episode 080 : Velero 1.0

- Hosted by @jbeda
- 2019-06-21

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=tj5Ey2bHsfM
" target="_blank"><img src="http://img.youtube.com/vi/tj5Ey2bHsfM/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:00 - Week in Review
- 00:19:00 - Velero and why it was renamed from Ark
- 00:21:00 - Let's get started, starting with the docs
- 00:25:00 - Creating an S3 bucket & IAM Account
- 00:34:00 - Run `velero install`
- 00:38:00 - Install nginx example workload
- 00:48:00 - Take a backup using `velero backup create`
- 00:52:00 - Look at contents of backup
- 01:03:00 - Create scheduled backup
- 01:07:00 - Install velero into cluster 2
- 01:09:00 - Migrate nginx app from cluster 1 into cluster 2
- 01:13:00 - Roadmap

## Week in Review

* [Kubernetes 1.15!](https://kubernetes.io/blog/2019/06/19/kubernetes-1-15-release-announcement/)
    * [Future of CRDs](https://kubernetes.io/blog/2019/06/20/crd-structural-schema/)
    * [Volume cloning is Alpha](https://kubernetes.io/blog/2019/06/21/introducing-volume-cloning-alpha-for-kubernetes/)
    *  Reminder [1.16 deprecation timeline](https://groups.google.com/forum/#!searchin/kubernetes-dev/jordan$20liggitt|sort:date/kubernetes-dev/je0rjyfTVyc/gEUw1YcyAQAJ)
* Congrats to the HAProxy folks for landing [HAProxy 2.0](https://www.haproxy.com/blog/haproxy-2-0-and-beyond/) with k8s support
* Bjorn Wenzel is doing a series on a [complete CI/CD setup](https://koudingspawn.de/the-complete-ci-cd-part-1/)
* Daniel Weibel did a post on [Boosting your kubectl productivity](https://itnext.io/boosting-your-kubectl-productivity-b348f7c25712?sk=c5819044a83b83d7d3bf52b5bc12e9a0)
* [bank-vaults](https://github.com/banzaicloud/bank-vaults/) - a set of tools for managing vault by Banzaicloud has a new milestone
* [Contour v0.13](https://github.com/heptio/contour/releases/tag/v0.13.0)

## Show Notes
* https://velero.io/
* https://velero.io/docs/v1.0.0/install-overview/
* Matty put that link here


## Reference Links



