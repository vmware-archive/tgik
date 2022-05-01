# Episode 184 : kpack

- Hosted by [@wiggitywhitney](https://twitter.com/wiggitywhitney) and [@ciberkleid](https://twitter.com/ciberkleid)
- 02/04/2022

<!--- Thumbnailed embed of the video, ShE_j4V_2iQ is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=ShE_j4V_2iQ
" target="_blank"><img src="http://img.youtube.com/vi/ShE_j4V_2iQ/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:06:09 - Week in review (and kpack questions from chat)
- 00:14:37 - Introduction to kpack
- 00:27:08 - Installing kpack
- 00:52:04 - Configuring kpack
- 01:45:23 - Using kpack (and troubleshooting)
- 02:26:02 - Examining our created image
- 02:36:35 - Rebasing images
- 02:44:07 - Questions and wrap up

## Week in Review
- Kubernetes is switching over to [EasyCLA](https://docs.linuxfoundation.org/lfx/easycla). According to this [GitHub issue](https://github.com/kubernetes/org/issues/2778) the migration is scheduled to happen this coming weekend.
- [Argo CD vulnerability](https://apiiro.com/blog/malicious-kubernetes-helm-charts-can-be-used-to-steal-sensitive-information-from-argo-cd-deployments/) discovered... But a [patch](https://github.com/argoproj/argo-cd/security/advisories/GHSA-63qx-x74g-jcr7) is available!
- According to [Last Week in Kubernetes Development](http://lwkd.info/2022/20220201), release Engineering published an [emergency extra update](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.23.md) to 1.23 (1.23.3) last week to fix a regression. Users who [use CRDs that rely on `x-kubernetes-preserve-unknown-fields: true`](https://github.com/kubernetes/kubernetes/pull/107688) should avoid 1.23.0-2, or update to 1.23.3 immediately.
- [More money committed to project Alpha-Omega](https://venturebeat.com/2022/02/01/google-and-microsoft-back-the-alpha-omega-project-to-bolster-software-supply-chain/) to design a secure global open source supply chain. [Webinar](https://openssf.org/blog/2022/02/02/openssf-webinar-introduction-to-project-alpha-omega/) to be held on Feb 16 if you want to learn more.
- The Call For Papers for [CNCF-hosted co-located events at KubeCon + CloudNativeCon Europe 2022](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/colocated-events/) are due soon, many on February 14! That is only 10 days away! Get to work! Well, I mean, get to work after you watch TGIK.
- Not exactly this week's news, but ICYMI and can't decide what to watch on TV tonight, catch the [two-part documentary on Kubernetes](https://twitter.com/CloudNativeFdn/status/1489277509199966208) by Honeypot!


## Show Notes

- What is [kpack](https://github.com/pivotal/kpack)?
    - One of several [tools](https://buildpacks.io/docs/tools/) for using [Cloud Native Buildpacks](https://buildpacks.io/)
    - What problem does it solve?
    - Who is it for?
- Installing kpack
    - From release file
    - Using [Tanzu Community Edition](https://tanzucommunityedition.io/)
- Configuring kpack
    - kp CLI
    - Create [Stack](https://github.com/pivotal/kpack/blob/main/docs/stack.md)
    - Create [Store](https://github.com/pivotal/kpack/blob/main/docs/store.md)
    - Create [ClusterBuilder](https://github.com/pivotal/kpack/blob/main/docs/builders.md)
- Using kpack
    - Creating [Images](https://github.com/pivotal/kpack/blob/main/docs/image.md)
    - Following Progress/Troubleshooting
- Examining created image
    - [Pack CLI](https://buildpacks.io/docs/tools/pack/)
    - [Dive CLI](https://github.com/wagoodman/dive)