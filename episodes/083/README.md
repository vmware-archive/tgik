# Episode 084 : Minecraft Controller with Kubebuilder

- Hosted by @jbeda
- 2019-07-19

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=tv-HJuwC4yI
" target="_blank"><img src="http://img.youtube.com/vi/tv-HJuwC4yI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:29 - Week in Review
- 00:17:42 - Putting Kubebuilder in context
- 00:22:56 - Minecraft image we'll be using
- 00:26:07 - Running Minecraft on k8s manually
- 00:34:10 - Ingress for Minecraft?
- 00:39:17 - Starting with kubebuilder
- 00:47:46 - Adding our first API type
- 00:53:43 - Aside: Operators vs. Controllers
- 01:02:18 - First run of controller!
- 01:04:11 - Leaders election and HA in k8s controllers
- 01:14:32 - Starting our reconciler
- 01:16:00 - RBAC requirements via annotations
- 01:29:47 - Aside: `generatedName` and its history
- 01:52:33 - Bug: reversed objects in `SetControllerReference` (James Munnelly caught it but I missed it in the chat!)
- 01:58:00 - First run with reconciler!
- 01:59:44 - Debugging first errors
- 02:08:20 - Almost there! CrashLoopBackoff
- 02:13:48 - Success!
- 02:14:00 - Wrapping up


## Week in Review
- In 1.16 some of the [deprecated apis will be removed!](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/)
- Kubecon SD CFP is closed! Good luck to those who have submitted!
- [Sysdig gives a rundown](https://sysdig.com/blog/33-kubernetes-security-tools/) of 33 Kubernetes specific security tools.
- Did you know that all the [videos from the 2019 Contributor Summit](https://www.youtube.com/playlist?list=PL69nYSiGNLP2WTJ6P8sQenhf0RY-JqF5L) in Barcelona are online?
- A new [special interest group](https://groups.google.com/forum/#!topic/kubernetes-sig-usability/2reiWMRYapU) emerges! Jump in!
- Release stuff!
    - Check out the new release notes website! [relnotes.k8s.io](https://relnotes.k8s.io/) by Jeff Sica
    - Stephen Augustus recorded the 1.16.0-alpha.1 release live. Check it out [part 1](https://youtu.be/VYhBqBoeAVY) [part 2](https://youtu.be/aOyZGleHGf0)
    - SIG-Release [update](https://docs.google.com/presentation/d/1t-bOgt6IfHW-TrdMfE3oopleSEqRgvECybZTq7GqVHU/edit#slide=id.g401c104a3c_0_0) and [video](https://youtu.be/RWbNg4Wjwpg?t=1337) from community meeting.
- IBM releases [3 new projects](https://developer.ibm.com/blogs/cloud-native-apps-kubernetes-kabanero/).


## Show Notes
- [crd structural schema changes](https://kubernetes.io/blog/2019/06/20/crd-structural-schema/)
- [Programming Kubernetes](https://learning.oreilly.com/library/view/programming-kubernetes/9781492047094/)
- [Kubebuilder Docs](https://book.kubebuilder.io/quick-start.html)
- [Minecraft Image](https://github.com/itzg/dockerfiles/tree/master/minecraft-server)
- [rcon-cli](https://github.com/itzg/rcon-cli) uses [rcon library](https://github.com/james4k/rcon)
- [Minecraft protocol](https://wiki.vg/Protocol#Handshake)
- [Basic Minecraft Operator using Operator-SDK](https://github.com/stgarf/minecraft-operator-go)
