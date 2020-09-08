# Episode 128 : RBAC tooling!

- Hosted by @mauilion
- Recording 08/14/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=ODTETX48yIU
" target="_blank"><img src="http://img.youtube.com/vi/ODTETX48yIU/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core K8s


- [Introducing Hierarchical Namespaces](https://kubernetes.io/blog/2020/08/14/introducing-hierarchical-namespaces/)
- 1.19 Release team is on break to prep for KubeCon, will resume on 24 August, 1.19 is close!
    - [Jeremy Rickard](https://twitter.com/jrrickard) will be your 1.20 release manager
- Dot Releases 
    - [1.18.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.8)
        - [Changes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#v1188)
    - [1.17.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.11)
        - [Changes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#v11711)
    - [1.16.14](https://github.com/kubernetes/kubernetes/releases/tag/v1.16.14)
        - [Changes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#v11614)

### Cloud Native Ecosystem

- Next week is [KubeCon + CloudNativeCon Europe Virtual](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/?gclid=Cj0KCQjw7Nj5BRCZARIsABwxDKKlZHz6lkgbTigmXCWfQLv2Q9X7HUJOnLkezvKn503XLhghNADt0WYaArO8EALw_wcB)
    - [Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/schedule/)
    - Check out the [mentoring sessions](https://groups.google.com/forum/#!topic/kubernetes-dev/FDwGepXJK5s)
- Docker has an updated [image retention policty](https://www.docker.com/pricing/retentionfaq)
- Helm Releases
    - [3.3.0](https://github.com/helm/helm/releases/tag/v3.3.0) focuses on `helm lint` and general stability fixes
    - [2.16.10](https://github.com/helm/helm/releases/tag/v2.16.10) final bugfix release for helm v2
- Mike Metral goes over the [Pulumi Kubernetes Operator](https://www.pulumi.com/blog/pulumi-kubernetes-operator/)
- Daniele Polencic goes over [Graceful Shutdown and Zero Downtime deployments](https://learnk8s.io/graceful-shutdown) - The [PDF diagram](https://learnk8s.io/a/graceful-shutdown.pdf) they use here is very useful! 
- First release of the [seccomp operator](https://github.com/kubernetes-sigs/seccomp-operator)
- Our break up with Weave Net and [how we did it with no downtime](https://blog.quentin-machu.fr/2020/08/07/our-breakup-with-weave-net/)


## Show Notes

## Reference Links

### at large!
- [rbac.dev](https://rbac.dev/)
- [consistent oidc on eks](https://aws.amazon.com/blogs/opensource/consistent-oidc-authentication-across-multiple-eks-clusters-using-kube-oidc-proxy/)

### client side!
- [rbac-view](https://github.com/jasonrichardsmith/rbac-view)
- [rbac-lookup](https://github.com/FairwindsOps/rbac-lookup)
- [access-matrix](https://github.com/corneliusweig/rakkess)
- [who-can](https://github.com/aquasecurity/kubectl-who-can)
- [kubectl-sudo](https://github.com/postfinance/kubectl-sudo)
- [audit2rbac](https://github.com/liggitt/audit2rbac)
### managers!
- [rbac-manager](https://github.com/FairwindsOps/rbac-manager)
- [permission-manager](https://github.com/sighupio/permission-manager)
