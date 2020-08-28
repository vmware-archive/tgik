# Episode 40 : Kubebuilder

- Hosted by @krisnova
- Recording date: 2018-06-22

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url
Note the 0.jpg for the thumbnail --->

<a href="http://www.youtube.com/watch?feature=player_embedded&v=N-lTSk1bGAg
" target="_blank"><img src="http://img.youtube.com/vi/N-lTSk1bGAg/0.jpg" 
alt="Episode 40 - Kubebuilder" width="640" height="360" border="10" /></a>

## Table of Contents

- [00:00:14](https://youtu.be/N-lTSk1bGAg?t=14s) - Welcome to TGIK!
- [00:02:56](https://youtu.be/N-lTSk1bGAg?t=2m56s) - Week in Review
- [00:10:50](https://youtu.be/N-lTSk1bGAg?t=10m50s) - Kubebuilder
- [00:13:48](https://youtu.be/N-lTSk1bGAg?t=13m48s) - Installing Kubebuilder
- [00:19:40](https://youtu.be/N-lTSk1bGAg?t=19m40s) - Running Kubebuilder
- [00:24:45](https://youtu.be/N-lTSk1bGAg?t=24m45s) - Creating our first API
- [00:26:10](https://youtu.be/N-lTSk1bGAg?t=26m10s) - Our first `kubebuilder init`
- [00:34:30](https://youtu.be/N-lTSk1bGAg?t=34m30s) - Looking at the files Kubebuilder just created (IDE time)
- [00:37:35](https://youtu.be/N-lTSk1bGAg?t=37m35s) - Why do Kubernetes projects have these hack directories? 
- [00:43:56](https://youtu.be/N-lTSk1bGAg?t=43m56s) - Adding a Schema and Business Logic
- [00:54:56](https://youtu.be/N-lTSk1bGAg?t=54m56s) - Integration Tests and Pushing to Kubernetes
- [01:09:09](https://youtu.be/N-lTSk1bGAg?t=1h9m9s) - Watching our status
- [01:15:00](https://youtu.be/N-lTSk1bGAg?t=1h15m) - Generating our documentation

Timestamps will be filled out a working day after the broadcast as we have to wait for YouTube to render the stream. Timestamps will also be published to the YouTube video description. 

## Show Notes

### Welcome to TGIK!

This week Kris will be looking at Kubebuilder, a framework for building Kubernetes APIs using [custom resource definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/).

### The Week in Review

- Clusterctl [release](https://github.com/kubernetes-sigs/cluster-api) - part of SIG Cluster Lifecycle. Example of how folks are thinking of deploying clusters in a standardized way. 
    - https://github.com/kubernetes-sigs/cluster-api/issues/383
- [GPUs go GA on GKE](https://cloudplatform.googleblog.com/2018/06/GPUs-as-a-service-with-Kubernetes-Engine-are-now-generally-available.html)
    - Check out Kris' talk with Dave Aronchik showing off [Kubeflow](https://www.youtube.com/watch?v=givpqZ2IchI)
- The [History of Kubernetes](https://blog.risingstack.com/the-history-of-kubernetes/)
- [Scaling Kubernetes to 25m users](https://medium.com/@brendanrius/scaling-kubernetes-for-25m-users-a7937e3536a0)
- [AWS Cost Savings by Utilizing Kubernetes Ingress with Classic ELB](https://akomljen.com/aws-cost-savings-by-utilizing-kubernetes-ingress-with-classic-elb/)
- [KubeIP](https://kubeip.com/)

## Kubebuilder

Part of SIG Cluster lifecycle / Cluster API. 

Starting with [main.go](https://github.com/kubernetes-sigs/kubebuilder/blob/master/cmd/kubebuilder/main.go). 

### Installing Kubebuilder

Do not use `go get`, grab the tarball for your system. 

- There's a [book](http://book.kubebuilder.io) with getting started documentation. We're following the quickstart section. 
- You'll end up with it in `/usr/local/kubebuilder/bin`
- Using GOOS and GOARCH (@krisnova mentions a video from Rob Pike on this, anyone have a link?)

## Reference Links

- https://github.com/kubernetes-sigs/kubebuilder
- GitHub Repo: [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)
- Slack channel: [#kubebuilder](http://slack.k8s.io/#kubebuilder)
- Google Group: [kubebuilder@googlegroups.com](https://groups.google.com/forum/#!forum/kubebuilder)
- Planned Work: [Sprint Dashboard](https://github.com/kubernetes-sigs/kubebuilder/projects/1)

- https://www.reddit.com/r/kubernetes/comments/8ien90/if_i_were_to_build_an_operator_what_should_i_use/
- https://itnext.io/under-the-hood-of-kubebuilder-framework-ff6b38c10796

## Documentation Used in this Episode 

- GitBook: [book.kubebuilder.com](http://book.kubebuilder.io)
