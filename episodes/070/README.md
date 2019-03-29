# Episode 70 : Assuming AWS roles with kube2iam/kiam

## - Hosted by @jbeda
- Recording date:  20190329

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=vgs3Af_ew3c
" target="_blank"><img src="http://img.youtube.com/vi/vgs3Af_ew3c/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:20 - Week in Review
- 00:12:00 - Getting started with kube2iam/kiam

## Week in Review

- [Kubernetes 1.14 has been released!](https://kubernetes.io/docs/setup/release/notes/)
- [Kubernetes release announcement](https://kubernetes.io/blog/2019/03/25/kubernetes-1-14-release-announcement/)
- Check out our [Kubernetes Release Party](https://twitter.com/cloudnativeapps/status/1111343973996609536) right after this
    - With Nicholas Lane, Duffie Cooley, and a handful of other engineers showcasing and kicking the tyres on the new release.
    - We'd love to do these for every release, thoughts, ideas or feedback? Let us know!
- [Kubernetes blog post: A guide to Kubernetes admission controllers](https://kubernetes.io/blog/2019/03/21/a-guide-to-kubernetes-admission-controllers/)
- [New – Advanced Request Routing for AWS Application Load Balancers](https://amzn.to/2TGTcaB)
- [Ubuntu 19 beta is available](https://itsfoss.com/ubuntu-19-04-release-features/)
    - > Beta release date: March 28, 2019
    - > 3. Linux Kernel 5.0
- A couple of weeks old! - [Ensuring Kubernetes Cost Efficiency across (many) Clusters - DevOps Gathering 2019](https://www.slideshare.net/try_except_/ensuring-kubernetes-cost-efficiency-across-many-clusters-devops-gathering-2019)

## Show Notes

- secure your aws credentials in regular use: [link](https://github.com/99designs/aws-vault)


## Reference Links

- [Episode Issue](https://github.com/heptio/tgik/issues/175)
- Analysis:
    - [kube2iam vs kiam](https://www.bluematador.com/blog/iam-access-in-kubernetes-kube2iam-vs-kiam)
    - [Comparison of options](https://docs.google.com/document/d/1rn-v2TNH9k4Oz-VuaueP77ANE5p-5Ua89obK2JaArfg/edit)
- Projects in this class:
    - [kube2iam](https://github.com/jtblin/kube2iam)
    - [kiam](https://github.com/uswitch/kiam)
    - [iam4kube](https://github.com/ash2k/iam4kube)
    - [kube-aws-iam-controller](https://github.com/mikkeloscar/kube-aws-iam-controller)
    - [AWS Future Plan](https://github.com/aws/containers-roadmap/issues/23)
- Similar for GCP: [k8s-gke-service-account-assigner](https://github.com/imduffy15/k8s-gke-service-account-assigner)
- Simple CLI to generate SSL certificates on any platform: [onessl](https://github.com/kubepack/onessl)
-

