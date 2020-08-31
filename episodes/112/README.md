# Episode 112 : Deep dive into admission controllers

- Hosted by @jbeda
- Recording date: 2020-04-03

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=fEvOzL_eosg
" target="_blank"><img src="http://img.youtube.com/vi/fEvOzL_eosg/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

YouTube crapped out a couple of times here. There are some gaps/jumps.

- 00:00:00 - Welcome to TGIK!
- 00:19:07 - Start digging into Admission Controllers
- 00:28:47 - Start looking at code
- 00:32:53 - Start trying to use ko for deploy
- 00:47:04 - Give up on ko and use a plain old Dockerfile
- 00:50:33 - Get deployment with TLS
- 01:10:46 - Configure webhook admission controller
- 01:19:35 - Start testing admission controller
- 01:21:51 - why isn't stuff logging?
- 01:29:12 - Working! Playing with admission controller

## Week in Review

### Core k8s

- [A KEP about KEPs!](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/617-improve-kep-implementation)
    - Updates to the KEPs and repo subdirectory layouts
- [k8s.gcr.io cutover](https://github.com/kubernetes/release/issues/270) to Kubernetes Community owned infrastructure in progress
    - Originally opened by Joe back in 2017!
- Discussing revamping the [issue triage workflow in `kubernetes/kubernetes`](https://github.com/kubernetes/enhancements/pull/1554)

### Cloud Native Ecosystem

- [videos.cncf.io](https://videos.cncf.io) it's an ai powered searchable youtube search tool [read more](https://www.cncf.io/blog/2020/04/03/introducing-a-new-tool-to-make-finding-your-favorite-cncf-videos-easier/)
- [Kubie](https://github.com/sbstp/kubie) - kubie is an alternative to kubectx, kubens and the k on prompt modification script. Check out [the blog post with why they built this tool](https://blog.sbstp.ca/introducing-kubie/)
- Using [linkerd with knative](https://linkerd.io/2020/03/23/serverless-service-mesh-with-knative-and-linkerd/) without istio.
- [MKIT](https://github.com/darkbitio/mkit) - query and validate several common security-related configuration settings of managed Kubernetes cluster objects and the workloads/resources running inside the cluster.
- Alex Ellis reflects on [5 years of Raspberry Pi Clusters](https://medium.com/@alexellisuk/five-years-of-raspberry-pi-clusters-77e56e547875)
- Google [has announced kpt](https://opensource.googleblog.com/2020/03/kpt-packaging-up-your-kubernetes.html), pronounced "kept". [Documentation](https://googlecontainertools.github.io/kpt/)
- [Cortex 1.0](https://twitter.com/bboreham/status/1245687848248004609?s=20)
- [ArgoCD 1.5.0 released](https://github.com/argoproj/argo-cd/releases/tag/v1.5.0)
- [from the Velero team] As mentioned in this week’s community meeting, we would like to hear from Velero community members on how we can help you use Velero, and build out a catalog of common usage patterns for backing up stateful applications. We would love to hear from the community at large here, and see how we can help protect their most valuable persistent apps.
    - Please take our 2 minute survey: https://forms.gle/4FEn97DfbX1PUrYF6
- CNCF SIG Contributor Strategy is live! - https://github.com/cncf/sig-contributor-strategy
    > Responsible for contributor experience, sustainability, governance, and openness guidance to help CNCF community groups and projects with their own contributor strategies for a healthy project.



## Show Notes

Code from the episode is up at https://github.com/jbeda/test-admission.

## Reference Links
* [Blog post](https://kubernetes.io/blog/2019/03/21/a-guide-to-kubernetes-admission-controllers/) on Admission Controllers
* [Docs on Dynamic Admission Controllers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/)
* [e2e test code](https://github.com/kubernetes/kubernetes/tree/master/test/images/agnhost/webhook)
* Brendan Burns [Customizing and Extending the Kubernetes API with Admission Controllers](https://youtu.be/P7QAfjdbogY)
*  Haowei Cai, Google [Admission Webhooks: Configuration and Debugging Best Practices](https://youtu.be/r_v07P8Go6w)
