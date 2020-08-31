# Episode 094 : SPIFFE and SPIRE

- Hosted by @jbeda
- Recording date: 2019-10-18

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=cx__8khtih4" target="_blank"><img src="http://img.youtube.com/vi/cx__8khtih4/hqdefault.jpg" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

- Kubernetes
  - [CVE-2019-11253](https://discuss.kubernetes.io/t/announce-cve-2019-11253-denial-of-service-vulnerability-from-malicious-yaml-or-json-payloads/8349): denial of service vulnerability from malicious YAML or JSON payloads
      - k8s 1.13.12, 1.14.8, 1.15.5, and 1.16.2 released which included CVE fixes. 
      - While we're here [heads up on golang CVE-2019-17596](https://groups.google.com/forum/#!msg/golang-announce/lVEm7llp0w0/VbafyRkgCgAJ)
  - Good status update from [SIG API Machinery this week](https://docs.google.com/presentation/d/17Nc5jmoIYyCIvKhfXlP6jESb28Q4E7bL-EA-zMEgB00/edit#slide=id.g401c104a3c_0_0)
  - 1.17 enhancements are now frozen
  - Now would be a good time to point out the [schedule for point releases](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md) <- handy!
- Microsoft announced 2 new projects
    - [Open Application Model](https://openappmodel.io)
    - [Dapr](https://dapr.io) - An event-driven, portable runtime for building microservices on cloud and edge.
- [Goldilocks](https://github.com/FairwindsOps/goldilocks/) -  This tool creates a Vertical Pod Autoscaler for each deployment in a namespace and then queries them for information.
- [Keiko](https://github.com/keikoproj) - Self Managing Clusters at Scale -[Presentation](https://github.com/keikoproj/keiko/blob/master/presentations/Keiko.pdf)
- [k-rail](https://github.com/cruise-automation/k-rail) - k-rail is a workload policy enforcement tool for Kubernetes. 
- Non-k8s: [Charity's blog post on why I didn't hire you](https://charity.wtf/2019/10/18/the-real-11-reasons-i-dont-hire-you/)
- Kubernetes Forums 2020 [Bengaluru](https://events19.linuxfoundation.org/events/kubernetes-forum-bengaluru-2019/program/cfp/) Feb 17 - 18 and [Delhi](https://events19.linuxfoundation.org/events/kubernetes-forum-delhi-2019/program/cfp/) Feb 20 - 21 CFP open

## Show Notes

- [Main SPIFFE site](https://spiffe.io/)
- [Original SPIFFE design doc](https://docs.google.com/document/d/1GjurNK2ROw4rXz-k-l68JtpGRkGj2fZcWqP6gksEriQ/edit)
- Link to TGIK repo issue: https://github.com/vmware-tanzu/tgik/issues/230
- [Getting started with SPIRE on k8s](https://spiffe.io/spire/getting-started-k8s/)
- [Correct tutorial install yaml files](https://github.com/spiffe/spire-tutorials)
- [An issue we are working around](https://github.com/spiffe/spire/issues/1081)
- [Using SPIFFE with Envoy](https://blog.envoyproxy.io/using-spire-to-automatically-deliver-tls-certificates-to-envoy-for-stronger-authentication-be5606ac9c75)


