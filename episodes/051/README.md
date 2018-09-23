# Episode 51 : Linkerd 2.0

- Hosted by @jbeda
- Recording date: 20180920

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=1dXJ34-KLe8
" target="_blank"><img src="http://img.youtube.com/vi/1dXJ34-KLe8/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Technical difficulties. Mic broken. Chat not scrolling. Sorry!
- 00:04:17 - Welcome to TGIK and hello!
- 00:08:01 - Week in review: Kubernetes Steering Committee Elections
- 00:10:06 - Week in review: Chucks cool YAML tool - kubeyaml.com
- 00:11:40 - Week in review: Kubernetes 1.12 release Real Soon Now
- 00:12:21 - Week in review: Pulumi k8s support
- 00:15:30 - Week in review: Heptio Contour 0.6 and Heptio Ark 0.9.6
- 00:17:47 - Starting to look at Linkerd 2.0
- 00:21:13 - Installing the Linkerd CLI
- 00:25:28 - Cluster preflight validation
- 00:26:34 - linkerd install - yaml spelunking
- 00:41:39 - Discussion on running/aggregating prometheus
- 00:44:08 - Actually install linkerd
- 00:46:15 - linkerd dashboard
- 00:46:46 - Digression on "fabulous mode" using lolgopher
- 00:51:32 - linkerd top
- 00:52:24 - Installing demo app
- 00:56:38 - Digression around cardinality when monitoring
- 01:00:43 - Pointing tgik.io to emojivoto
- 01:02:11 - Injecting sidecar to emojivoto
- 01:03:55 - Watch it run with dashboard and CLI tools
- 01:07:30 - Digression around source IP vs. true source
- 01:15:11 - Starting to wrap up. Some comparison to Istio
- 01:17:10 - Trying to remove linkerd from example app
- 01:20:02 - Scaling out the demo app
- 01:22:21 - "The cats and the dogs and the rainbows are all fighting"
- 01:23:44 - Stream goes down because Kris overloads the Heptio network

## Show Notes

[LinkerD Getting Started](https://linkerd.io/2/getting-started/index.html)
[CLI Installation](https://linkerd.io/2/getting-started/#step-1-install-the-cli) - as usual we recommend inspecting any .sh file that you curl to install something.
Things learned from Linkerd devs in chat:
 - CLI client and control plane is golang
 - Data plane is in rust (this is new)
`linkerd install` doesn't do anything, it outputs a yaml file for you. So no dep on helm or other tools, just the raw yaml. You then pipe this to a file and/or kubectl. install takes some flags for versions, etc.

Checking the yaml file (~854 lines.)
- Controllers are read-only, no reading of secrets. (We like to see things like this)
- Not running as root (user 2102) in SecurityContext
    - upstream k8s will start to enforce things like this soon, openshift already does this.

Exploring the dashboard - https://linkerd.io/2/getting-started/index.html#step-4-explore-linkerd

## Reference Links

* Heptio Academy with Craig McLuckie: [San Francisco](https://www.eventbrite.com/e/heptio-academy-hands-on-kubernetes-training-in-san-francisco-tickets-49439435683), [New York].(https://www.eventbrite.com/e/heptio-academy-hands-on-kubernetes-training-in-new-york-tickets-49439430668)
* [Election](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-dev/0gEdp_xdzEI/214wEaVbCgAJ)
* [Chuck's YAML Tool](https://kubeyaml.com/)
* [1.12 is here](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.12/README.md)
* [Pulumi k8s stuff](https://info.pulumi.com/press-release/pulumi-cloud-native-sdk-delivers-cloud-native-infrastructure-as-code-for-kubernetes)
* [Service Mesh with William Morgan](https://softwareengineeringdaily.com/2017/06/26/service-mesh-with-william-morgan/)
* [LinkerD](https://kubernetes.io/blog/2018/09/18/hands-on-with-linkerd-2.0/)
  * [LinkerD getting started](https://linkerd.io/2/getting-started/)
* New Heptio releases - [Ark 0.9.6](https://github.com/heptio/ark/releases) and [Contour 0.6](https://github.com/heptio/contour/releases)
