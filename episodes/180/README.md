# Episode 180 : Kubernetes 1.23 release review

- Hosted by @jbeda
- Recording date: 2022-01-07 

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=KlA8alUCz-8" target="_blank"><img src="https://i.ytimg.com/vi/KlA8alUCz-8/maxresdefault.jpg" border="10" /></a>

## Video Timeline

- 00:03:49 - Welcome to TGIK and welcome Rey!
- 00:07:11 - Week in Review
- 00:26:10 - What is a Kubernetes Release Lead?
- 00:33:36 - Dual stack IPv4/v6
- 00:37:37 - PodSecurity Admission to beta
- 00:42:06 - TTL Controller
- 00:44:43 - HPA v2 API to GA
- 00:52:17 - Ephemeral Containers and `kubectl debug`
- 00:59:16 - Auto Remove PVCs from StatefulSets to alpha
- 01:00:44 - SLSA Level 1 compliance
- 01:05:26 - 1.23 Lightning Round
- 01:06:52 - Thank you and goodbye

## Summary

For a special #TGIK8s this week we have a special guest. @reylejano was 1.23 release lead and will be going over how a release happens and what his favorite enhancements are for this release! 

## Week in review
- [DockerShim deprecation timeline](https://kubernetes.io/blog/2022/01/07/kubernetes-is-moving-on-from-dockershim/)
- [Minecraft as k8s admin tool](https://eric-jadi.medium.com/minecraft-as-a-k8s-admin-tool-cf16f890de42)
- [Great breakdown of Carvel and packages](https://beyondelastic.com/2022/01/04/tanzu-packages-explained/)
- [IPv6 support in EKS](https://twitter.com/_msw_/status/1479165244798685184)
- [IPv6 support in AKS (Preview)](https://twitter.com/LachlanEvenson/status/1479523186995712003)
- [Shout out to Liz Rice who is not standing for CNCF TOC membership again](https://twitter.com/lizrice/status/1479448125253242897)
- [Register for virtual KubeCon! $10 early bird deal ends today](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/register/)


## Show Notes

His favorite enhancements:

* Dual-stack IPv4/IPv6 is GA
* PodSecurity Admission goes to beta (PSP replacement)
* TTL Controller (Cleanup Jobs and Pods of Jobs) goes to GA
* HPA v2 API is GA
* Ephemeral Containers (supports kubectl debug) goes to beta
* Auto remove PVCs from StatefulSets is intro'd as alpha
* SLSA Level 1 compliance (provenance attestation files are generated that describe staging and release phases)
* Quick review:
    * Kubelet CRI went to beta (pre docker-shim)
    * Supply chain: taint based propagation analysis. How does data move within an application and make sure we don't log secrets.

v1.24 -- cycle starts on monday through April 24th.  Release cycle is 15-16 weeks.

Here's the [release blog post](https://kubernetes.io/blog/2021/12/07/kubernetes-1-23-release-announcement/).
About the [Release Team](https://github.com/kubernetes/sig-release/blob/master/release-team/README.md)
[Release Team Selection](https://github.com/kubernetes/sig-release/blob/master/release-team/release-team-selection.md)
[SIG Release mailing list](https://groups.google.com/g/kubernetes-sig-release)
[#sig-release Slack channel](https://kubernetes.slack.com/archives/C2C40FMNF)
More about [release roles](https://github.com/kubernetes/sig-release/tree/master/release-team/role-handbooks)
[Deprecated API Migration Guide](https://kubernetes.io/docs/reference/using-api/deprecation-guide/)
