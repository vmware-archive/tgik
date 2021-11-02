# Episode 155 : Backstage
- Hosted by @jbeda
- 05/21/2021

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=KeQnIagZYA4
" target="_blank"><img src="http://img.youtube.com/vi/KeQnIagZYA4/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:03:20 - Week in Review
- 00:20:17 - Starting with backstage
- 00:30:41 - Getting node env set up
- 00:32:20 - Creating the backstage app
- 00:35:54 - Exploring app structure
- 00:38:40 - Launching the app
- 00:46:38 - Configuring Auth
- 00:59:16 - Thinking about confused deputies
- 01:04:34 - Starting to explore catalog w/ Kubernetes
- 01:41:33 - Create new component from template
- 01:53:38 - Summing up

## Week in Review

### Core K8s
* [KEP-2558: Publish versioning information in OpenAPI](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2558-publish-version-openapi)
* [KEP-2590: Kubectl Subresource Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2590-kubectl-subresource)
* [KEP-2133: Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers)
* [KEP-2040: Kubelet CRI support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2040-kubelet-cri)
* [KEP-2413: Enable seccomp by default](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2413-seccomp-by-default)
* [KEP-2579: Pod Security Admission Control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2579-psp-replacement) (PSP Replacement)
* [1.22 release info](https://www.kubernetes.dev/resources/release/)

### K8s and Cloud Native Ecosystem
* Look at all the sandbox projects! https://www.cncf.io/sandbox-projects/
* [vcluster](https://www.vcluster.com/)
* Blog Post: [The Easiest Way to Debug Kubernetes Workloads](https://martinheinz.dev/blog/49)
* [Kind v0.11.0 Release](https://github.com/kubernetes-sigs/kind/releases/tag/v0.11.0)
* [WTFisSRE Conference](https://www.cloud-native-sre.wtf) May 20th, talks recorded and still available if you register and login.
  * [Full day stream of Track 1](https://youtu.be/aQ8ivGNDOyQ)
  * Excellent talk by sig-security chair Tabitha Sable & Ellen Körbes: [Hacking into Kubernetes Security for Beginners](https://youtu.be/aQ8ivGNDOyQ?t=31013)

## Show Notes

* Main page for Backstage: backstage.io
* Must use Node 14. Tricky to do with homebrew. Ended up installing nvm via [this guide](https://yoember.com/nodejs/the-best-way-to-install-node-js-with-yarn/).  I'm a total noob here so don't listen to me if you know better.
