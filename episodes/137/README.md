# Episode 137 : Waypoint
- Hosted by @joshrosso
- 10/22/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=YkOOCyK6Yak
" target="_blank"><img src="http://img.youtube.com/vi/YkOOCyK6Yak/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://www.youtube.com/watch?v=YkOOCyK6Yak)
- [00:03:53 - Week in Review](https://youtu.be/YkOOCyK6Yak?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=233)
- [00:18:30 - Waypoint & k8s Abstractions](https://youtu.be/YkOOCyK6Yak?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1110)
- [00:35:01 - Installing Waypoint and Doing a Deployment](https://youtu.be/YkOOCyK6Yak?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=2101)
- [01:01:21 - Installing Metal LB to enable Waypoint server](https://youtu.be/YkOOCyK6Yak?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=3680)
- [01:14:20 - Using Waypoint Server and doing more deploys!](https://youtu.be/YkOOCyK6Yak?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=4460)
- [01:49:11 - Wrap-up and Goodbye!](https://youtu.be/YkOOCyK6Yak?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=6551)

## Week in Review

### Core K8s

- [K8s 1.20alpha3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#v1200-alpha3)
    - November 12 code freeze, keep an eye on [k8s.dev/release](https://k8s.dev/release) for the schedule
    - [Security Advisory] [Multiple secret leaks when verbose logging is enabled](https://discuss.kubernetes.io/t/security-advisory-multiple-secret-leaks-when-verbose-logging-is-enabled/13282)

### K8s and Cloud Native Ecosystem

- [My Kubernetes aliases and functions using fzf](https://bradbarrows.com/post/kubernetesaliasesfunctions)
- The learnk8s folks have put together a comprehensive [comparison of Kubernetes ingress controllers](https://docs.google.com/spreadsheets/d/191WWNpjJ2za6-nbG4ZoUMXMpUK8KlCIosvQB0f-oq3k/edit#gid=907731238) (Google Spreadsheet)
- Important Ops PSAs (help spread the word):
    - Important weekly reminder about the [Helm changes happening in November](https://helm.sh/blog/helm-v2-deprecation-timeline/)!
        - [CNCF Announcement](https://www.cncf.io/blog/2020/10/07/important-reminder-for-all-helm-users-stable-incubator-repos-are-deprecated-and-all-images-are-changing-location/)
    - Docker Hub [changes take effect](https://www.docker.com/blog/scaling-docker-to-serve-millions-more-developers-network-egress/) November 1st!
- [Kubeacademy AMA recording with Joe Beda](https://tanzu.vmware.com/content/webinars/oct-22-ask-me-anything-with-joe-beda-co-creator-of-kubernetes)
- [Hands-on CKA/CKAD with a CKS flavour](https://www.youtube.com/watch?v=jZOs8Oips7Q) next Friday 8:30 PM IST with friends of the show @SaiyamPathak and @walidshaari contributors to [Certified-Kubernetes-Security-Specialist](https://github.com/walidshaari/Certified-Kubernetes-Security-Specialist) GitHub repo.
- [kubectl-aliases - Programmatically generated handy kubectl aliases](https://github.com/ahmetb/kubectl-aliases)
- [learnk8s telegram group](https://t.me/learnk8s)


## Show Notes

- [HashiCorp Waypoint](https://learn.hashicorp.com/waypoint)
- [HashiCorp Waypoint @GitHub](https://github.com/hashicorp/waypoint)
- [Automating Execution](https://www.waypointproject.io/docs/automating-execution) -- see sub-bullets for CI integration
- [Uber Peloton talk from KubeCon 2018](https://www.youtube.com/watch?v=USgbj87Ztlk)
