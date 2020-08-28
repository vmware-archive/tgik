# Episode 107 : pod logging and fluent-bit

- Hosted by @mauilion
- Recording date: 2020-02-28

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=BmTCtLeuROU
" target="_blank"><img src="http://img.youtube.com/vi/BmTCtLeuROU/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review
- A word from our sponsors  [app modernization](https://www.vmware.com/app-modernization.html)
- A new podcast by @richburroughs [kubecuddle](https://kubecuddle.transistor.fm/)
- kubecon eu is right around the corner! [details](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/?gclid=EAIaIQobChMIovni96fy5wIVC77ACh1YLA5XEAAYASAAEgIVffD_BwE)
- schedules for the day 0 events are up! [schedules](https://www.cncf.io/blog/2020/02/24/schedules-announced-for-cloud-native-security-day-serverless-practitioners-summit-servicemeshcon)
- Google has dropped a application manager for GKE [application-manager](https://cloud.google.com/blog/products/containers-kubernetes/announcing-application-manager-for-google-kubernetes-engine)
- Kubernetes is extensible! Check out [pangolin](https://github.com/dpeckett/pangolin)
- Great article from Neuvector! on [i/o intensive containers](https://neuvector.com/container-security/optimize-i-o-intensive-containers/)
- a recent review of the security model of helm 3 [cncf webinar](https://www.cncf.io/webinars/helm-security-a-look-below-deck/)
- Fluentd is part of the cncf here is an [overview](https://www.cncf.io/blog/2020/02/26/cncf-tools-overview-fluentd-unified-logging-layer/)


## Show Notes
- [x] Logging for containers
- [x] Logging for pods
- [x] These things are ephemeral and jump around a lot!
- [x] Where do pod logs go?
- [x] How long do they stay around?
- [x] What is kubectl logs -p :) 
- [x] How do we get those logs off the cluster and into your choice of aggregator
- [x] History:
    * fluentd
    * fluent-bit
- [x] deploy elastic and kibana locally
- [x] A bit about our environment!
- [x] What's a [fluentbit](https://docs.fluentbit.io/manual/)
- [ ] A helm [chart!](https://hub.kubeapps.com/charts/stable/fluent-bit)
- [ ] get some things sending data to them!
- [ ] tune and adjust things
- [ ] profit

Shout out to Eduardo Silva @edsiper

- Fluent Bit is deployed millions of times every month
- Fluent Bit v1.4 is around the corner:
  - rewrite tag filter
  - aws metadata filter
  - extended stream processing capabilities
  - granular logging level per plugins
  - much more

## Reference Links
- [kube-eventer](https://github.com/AliyunContainerService/kube-eventer)
