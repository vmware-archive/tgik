# Episode 149 : GitOps with Flux v2
- Hosted by @e_k_anderson
- 03/19/2021

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=v47-NTb6Qw8
" target="_blank"><img src="http://img.youtube.com/vi/v47-NTb6Qw8/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review
- 00:18:38 - What is GitOps
- 00:20:20 - Flux Architecture
- 00:23:30 - Installing & Setting Flux on Windows
- 00:31:57 - Review of CDRs and RBAC
- 00:34:30 - Git add & git commit
- 00:35:10 - Apply manifest
- 00:36:16 - Create git source
- 00:44:15 - Create kustomization object
- 00:48:30 - Switch back to github
- 01:00:00 - Bootstrap done!
- 01:01:00 - Digression about long-lived vs short-term clusters
- 01:09:00 - Flux bootstraps itself to run under GitOps!
- 01:14:00 - Battling with character encodings (UTF-16 file)
- 01:22:00 - Blame powershell redirects for making the file UTF-16
- 01:25:00 - How to stop reconciliation for controllers if needed
- 01:28:00 - Migrating systems into "gitops mode"
- 01:29:00 - Self-healing when Evan monkeys with flux components
- 01:32:00 - Exploring the status of Flux resources
- 01:39:00 - Exploring `flux export`
- 01:45:00 - Grafana dashboards

## Week in Review

### Core K8s

- Kubernetes patch releases went live:
  - [Kubernetes 1.18.17](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md) built with Golang 1.13.15
  - [Kubernetes 1.19.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md) built with Golang 1.15.8
  - [Kubernetes 1.20.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md) built with Golang 1.15.8
- [Windows Conformance is preparing a KEP](https://github.com/kubernetes/enhancements/issues/2578)

### K8s and Cloud Native Ecosystem

- [Flux is a CNCF incubating project](https://fluxcd.io/blog/2021/03/flux-is-a-cncf-incubation-project/)
- [GitOps Application Manager](https://github.com/redhat-developer/kam)
- [Sonobuoy v0.50 is now released - distroless!](https://sonobuoy.io/sonobuoy-v0-50/)
- [GitOpsCon EU 2021 (3rd May) - Call For Proposals
](https://docs.google.com/forms/d/e/1FAIpQLSeNahDbiEolx6WZmtxx4L65qmq7pZTX86nQAltq2uC12tCQYg/viewform)
- [skarp/runj - Experimental OCI runtime for BSD Jails :)](https://github.com/samuelkarp/runj)
- [GitOps Days 2021 CFP is open (event is 9-10 June)](https://www.gitopsdays.com/)

## Show Notes

- [Gitea](https://docs.gitea.io/en-us/install-on-kubernetes/)
- [Flux V2](https://toolkit.fluxcd.io/)
