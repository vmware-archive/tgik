# Episode 106: Flagger - Advanced Application Rollout

- Hosted by [@joshrosso](https://twitter.com/joshrosso)
- Recording 02.21.2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=HQX6wImrLTw
" target="_blank"><img src="http://img.youtube.com/vi/HQX6wImrLTw/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://youtu.be/se4soMN4fDg)
- [00:02:05 - Week in Review](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=127)
- [00:19:00 - Flagger Introduction](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1140)
- [00:24:32 - Contour Deploy](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1472)
- [00:28:26 - Flagger Deployment](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1706)
- [00:53:06 - Canary Test (success and failure)](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=3186)
- [01:16:12 - A/B Test](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=4572)
- [01:29:59 - Summary and Wrap-up](https://youtu.be/HQX6wImrLTw?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=5399)

## Week in Review

#### Core

- 1.18 News
  - Currently we are tracking 50 enhancements 
    - Alpha: 18
    - Beta: 16
    - Stable: 16
    -   v1.18.0-alpha.5 - same as last week, no change! 
    -   Code freeze scheduled for Thursday, March 05, 2020 [sig-release/releases/release-1.18](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.18)
    -   The Plan: Tuesday, March 24: Week 12 - Kubernetes v1.18.0 released
    -   [Patch Release Updates](https://git.k8s.io/sig-release/releases/patch-releases.md)
        - 1.15.10, 1.16.7, and 1.17.3 were released on 2020-02-11
- Two [Office Hours](https://github.com/kubernetes/community/blob/master/events/office-hours.md) sessions covering [many topics](https://discuss.kubernetes.io/t/kubernetes-office-hours-for-feb-2020-new-us-west-session-too/9689?u=castrojo)

#### Ecosystem

- [Architecting Kubernetes clusters — how many should you have?](https://learnk8s.io/how-many-clusters)
- [CPU limits and aggressive throttling in Kubernetes](https://medium.com/omio-engineering/cpu-limits-and-aggressive-throttling-in-kubernetes-c5b20bd8a718)
- [Zalando slides regarding Optimizing kubernetes  resource request, limits for cost efficeny and lateny highload](https://www.slideshare.net/try_except_/optimizing-kubernetes-resource-requestslimits-for-costefficiency-and-latency-highload)
- [Complete guide to deploy Spark on Kubernetes](https://developer.sh/posts/spark-kubernetes-guide)


## Flagger

**Agenda:**

* [x] Install / Configure
* [x] Traffic Shifting
    * [x] Canary
    * [x] A/B
    * [x] Blue Green  
          ( Covered at a conceptual level )

## References

* Flagger: https://flagger.app
* TGIK Argo: https://www.youtube.com/watch?v=M_rxPPLG8pU
* TGIK Flux: https://www.youtube.com/watch?v=aQz3H9bIH8Y
* Zalando Kube optimization (CPU Limits / Requests): https://www.slideshare.net/try_except_/optimizing-kubernetes-resource-requestslimits-for-costefficiency-and-latency-highload
