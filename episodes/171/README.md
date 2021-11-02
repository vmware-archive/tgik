# Episode 171 : Pod Security Problems
- Hosted by @e_k_anderson and @pudijoglekar
- 10/22/2021

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=Vk1ARLbAcVc
" target="_blank"><img src="http://img.youtube.com/vi/Vk1ARLbAcVc/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:01:20 - Week in Review
- 00:12:00 - Intro to Pod Security Policy
- 00:20:00 - Enabling PSP on a cluster
- 00:22:00 - Digression about how Windows is still a big developer thing
- 00:24:00 - PSP causes `apiserver` pod to disappear in pod list
- 00:28:00 - Figuring out how to fix cluster, talking about API design of open vs closed by default
- 00:30:00 - Going over PodSecurityPolicy API deficiencies
- 00:43:00 - Introducing the replacement
- 00:44:00 - Shout-out to Kube Security Admisison testing tooling!
- 00:46:00 - Attempting to enable PodSecurityAdmission feature gate
- 00:48:00 - Debugging `apiserver` failing to come up after flag change
           * This resolved after the episode; every second server start worked, so changing some other pod value got it working
- 00:59:00 - Trying again with `kind`
- 01:02:00 - `kind` times out
- 01:02:00 - Pod Security Standards -- how they came to be, what's in them
- 01:10:00 - Audit and Warn settings
- 01:16:00 - Versioning Pod Security Standards
- 01:20:00 - Discussion on how to [move towards secure-by-default](https://docs.google.com/document/d/1Y1OCgto48Woc0UsZHq7zHuYOFGYI_DTpC2o7k7KYeaA/edit#)
- 01:22:00 - `dry-run` on increasing security levels
- 01:30:00 - Wrapping up


## Week in Review

- First Hybrid Kubecon happened!
    - CNCF twitter moments
        - [Day1](https://mobile.twitter.com/i/events/1448410029460111360)
        - [Day2](https://mobile.twitter.com/i/events/1448769340694667267)
        - [Day3](https://mobile.twitter.com/i/events/1449110714572886018)
    - Link any talks you found especially interesting!
        - Neat talk from SolarWinds about Supply Chain Security
        - Great talk about burnout experiences on Wednesday: Julia Simon
        - Ian and Brad had a good talk "seven of nine kubernetes security secrets"
        - I liked the simple demos in "Back to the Drawing Board: Building Containers with SBoMs" by Nisha Kumar
        - As usual, Tabitha and Ellen's talk was amazing: "Beyond Kubernetes Security - Ellen Körbes, Tilt & Tabitha Sable, Datadog"
        - (mentioned by PJ) - "PodSecurityPolicy Replacement: Past, Present, and Future - Tabitha Sable, Datadog & Tim Allclair, Apple"
    - 

### Core K8s

- Code freeze is coming soon for 1.23! (November 16): https://github.com/kubernetes/sig-release/tree/master/releases/release-1.23
- Cluster API v1: https://www.cncf.io/blog/2021/10/06/kubernetes-cluster-api-reaches-production-readiness-with-version-1-0/
- ingress-nginx high CVE: https://github.com/kubernetes/ingress-nginx/issues/7837

### K8s and Cloud Native Ecosystem

- [Knative's next release will be 1.0](https://knative.dev/blog/articles/announcing-knative-1.0/)


## Show Notes

- https://kubernetes.io/docs/concepts/policy/pod-security-policy/
- https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/
- https://medium.com/@LachlanEvenson/hands-on-with-kubernetes-pod-security-admission-b6cac495cd11
- https://github.com/appscodelabs/tasty-kube/tree/master/kind/psp
- [Kubecon NA 2021 reveiw blog by Loft/Rich Burroughs](https://loft.sh/blog/kubecon-2021-los-angeles-wrapup/)
- https://kubernetes.io/docs/concepts/security/pod-security-admission/
- https://docs.google.com/document/d/1Y1OCgto48Woc0UsZHq7zHuYOFGYI_DTpC2o7k7KYeaA/edit#
- [Guidelines for secure-by-default Cloud Native software](https://docs.google.com/document/d/1Y1OCgto48Woc0UsZHq7zHuYOFGYI_DTpC2o7k7KYeaA/edit#)


