# Episode069 : Exploring Gloo on Kubernetes

- Hosted by @jbeda
- Hosted by @krisnova
- Recording date: 2019-03-22

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=

" target="_blank"><img src="http://img.youtube.com/vi/q-L2woZMBHY/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Last week in Kubernetes Development (lwkd.info)
- [New scope doc in github.com/kubernetes/kubernetes](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/scope.md)
- [The documentation for HTTP readiness and liveness probes has always said that any HTTP response with a code between 200 and 400 would be considered a success, but until now 301 and 302 redirects have been expanded by the kubelet and queried again. As a compromise between matching the spec and not overly breaking backwards compatibility, redirects to the same hostname as used originally will still be followed, but a redirect to a new hostname will be considered a success and not followed. This may result in a loss of accuracy in some very rare cases such as clustered applications redirecting to each other. But on the plus side it improves the security of Kubernetes and closes a potential traffic amplification vector.](https://github.com/kubernetes/kubernetes/pull/75416)

## Kubeweekly https://t.co/JtzIHddg3D

- [Shout out to Tim St. ClairThe what and why of cluster-api](https://blogs.vmware.com/cloudnative/2019/03/14/what-and-why-of-cluster-api/) 
- [Shout out to Marko Mudrinic for "running kubernetes in ci (with kind!)"](https://www.loodse.com/blog/2019-03-12-running-kubernetes-in-the-ci-pipeline-/)
- Shout out to Paris Pittman for a [retro and look forward on contributor summit](https://kubernetes.io/blog/2019/03/20/a-look-back-and-whats-in-store-for-kubernetes-contributor-summits/)

## Show Notes

### Glö

"Overjoyed" in Icelandic

### About Gloo

Built on envoy. Great quote below.



###### Patrick Barker [10:51 AM]
> so the tldr on gloo: you know how istio, ambassador, contour, etc are all just ways of configuring envoy on kubernetes? Gloo decomposed that process down into more fundamental types https://gloo.solo.io/dev/ making it really a way to build envoy configuration services.


## Reference Links



 - [New Kubernetes Dashboard](https://hub.docker.com/r/herbrandson/k8dash)
 - [Gloo Userguides](https://gloo.solo.io/user_guides/)
 - [Envoy Universal data plane API](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a)
 - [Annotation Documentation](https://gloo.solo.io/user_guides/basic_ingress/)
