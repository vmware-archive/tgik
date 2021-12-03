# Episode 174 : Verifying Signed Images with Connaisseur

- Hosted by Pushkar Joglekar [\@PuDiJoglekar](http://twitter.com/PuDiJoglekar)
- Guest: Christoph Hamsen [\@xophham](https://twitter.com/xophham)
- Recording date: 2021-11-19 5 PM GMT / 9 AM PST

<a href="https://youtu.be/LFAmi39CBb4" target="_blank"><img src="https://i.ytimg.com/vi/LFAmi39CBb4/maxresdefault.jpg" border="10" /></a>

Join Pushkar with Christoph from SSE - Secure Systems Engineering GmbH, as we
explore, why having signed images only is not enough. But having a way to verify
them is equally important. In this episode, we will explore connaisseur project
that acts as an admission controller to verify signed images at deployment time.

## Table of Contents
- 00:00:00 - Welcome to TGIK + Guest Intro!
- 00:04:20 - Week in Review
- 00:09:12 - Origin Story of Connaisseur
- 00:22:10 - Cloning Connaisseur Github repo
- 00:25:00 - Install Connaisseur with Helm
- 00:33:00 - Create kind cluster
- 00:38:00 - Deploy Hello World image pod
- 00:45:00 - Deploy unsigned image
- 00:50:00 - Deploy signed image
- 01:05:00 - Deep dive into Connaisseur Configuration File
- 01:32:00 - Discussing Threats and potential Open source contributions
- 01:38:00 - Wrapping up

## Week in Review

* *Kubernetes v1.23* is almost there - Code freeze is now ongoing, so only
  patches and PRs that blocks releases will be accepted!
* November Kubernetes patch releases are out!
  - https://twitter.com/puerco/status/1461176447742226440?s=20
* Go v1.18 is almost there, and this
  cool [Twitter thread](https://twitter.com/mvdan_/status/1456947756925399040?s=12)
  covers A LOT of the new stuff: fuzzing, generics, and even a new net/netip
  package and a new strings/bytes.Cut()
* Gatekeeper 3.7.0
  was [released](https://github.com/open-policy-agent/gatekeeper/releases/tag/v3.7.0)
  with a bunch of new features: Mutation moved to Beta (yeey), and a new CLI (
  alpha) so you can test constraints and constrainttemplates without Kubernetes
* Dockershim is deprecated and marked for removal in v1.24 and SIG-Node is
  collecting some feedback
  - https://kubernetes.io/blog/2021/11/12/are-you-ready-for-dockershim-removal/
* Pod Security Admission moving to Beta and will be enabled by default in v1.23.
  This tweet has some useful
  links: https://twitter.com/tallclair/status/1460386502555230216
*

## Show Notes

* Introductions
* TGIK schedule: https://github.com/vmware-tanzu/tgik
* More APAC / EMEA timezone friendly TGIKs are here!
* What are signatures and how do you verify them?
* Why sign images?
* Why verify images?
* [Connaisseur Repository](https://github.com/sse-secure-systems/connaisseur)
* [Connaisseur Documentation](https://sse-secure-systems.github.io/connaisseur/v2.2.0/)
  * [Quick Start](https://sse-secure-systems.github.io/connaisseur/v2.2.0/#quick-start)
    provides a simple demo to get started with minimal configuration
  * [Getting Started](https://sse-secure-systems.github.io/connaisseur/v2.2.0/getting_started/)
    with signing and verifying your first own images
* [Recent Blog Post](https://medium.com/sse-blog/verify-container-image-signatures-in-kubernetes-using-notary-or-cosign-or-both-c25d9e79ec45)
  on Signature Verification for both, notary and cosign

## Wait there's more!!

- TGIK will have its first ever double header. Next session is happening today,
  on regularly scheduled time 1pm PST: tgik.io/175

