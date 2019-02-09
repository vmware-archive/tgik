# Episode 64 : Kubernetes config with Pulumi

- Hosted by @jbeda
- Recording date: 2019-02-08

https://www.youtube.com/watch?v=ILMK65YVSKw
<a href="https://www.youtube.com/watch?v=ILMK65YVSKw
" target="_blank"><img src="http://img.youtube.com/vi/ILMK65YVSKw/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

(timestamp is time after the hour, not show time)

- 00:00:00 - Welcome to TGIK!
- 00:04:00 - Week in Review
- 00:14:00 - Pulumi!
- 00:24:00 - Creating our new cluster for Pulumi
-
-

## Week in Review

- All quiet in upstream k8s land, no releases or milestones this week
- [Kubernetes @ CERN](https://speakerdeck.com/rochaporto/kubernetes-at-cern-use-cases-integration-and-challenges)- Presentation
    - [Video](https://www.youtube.com/watch?v=2PRGUOxL36M) of Kubecon Presentation
- [Troubleshooting & Debugging Microservices in Kubernetes](https://www.youtube.com/watch?v=2hxTTyc6IH8) - GOTO 18 talk • Ray Tsang & Robert Kubis
- Argo blog: [The State of Kubernetes Configuration Management: An Unsolved Problem](https://blog.argoproj.io/the-state-of-kubernetes-configuration-management-d8b06c1205)
- Last weekend was [FOSDEM](https://fosdem.org/2019/)
    - Lots of great content slides and videos of sessions hosted at fosdem are all available via the site.
    - from [falco team](https://fosdem.org/2019/schedule/event/falco_container_monitoring/) - we did an episode on Falco, check out [Episode 61](https://www.youtube.com/watch?v=fjoxmTeh3Fk)
    - Seth Vargo on [kubernetes secrets](https://fosdem.org/2019/schedule/event/base64_not_encryption/)
    - [Cilium update](https://fosdem.org/2019/schedule/event/cilium_overview_and_updates/) - We looked at Cilium in [Episode 47](https://www.youtube.com/watch?v=I8Tp7jU2oJk)
    - Our own Kris Nova talks about the wonderful [Kubernetes Code base](https://fosdem.org/2019/schedule/event/kubernetesclusterfuck/)
- Some [Heptio projects have been renamed](https://blogs.vmware.com/cloudnative/2019/02/05/welcoming-heptio-open-source-projects-to-vmware/)
- [k9ss.io](https://k9ss.io/) - ncurses terminal UI for k8s clusters


## Show Notes

List of languages supported:
- Typescript, JavaScript, Python, Go
- C# (community supported) (Joe Duffy: "Early .NET is looking nice")

Joe Duffy: the @ means it's namespaced in an org. We tend to put all Pulumi packages, like @pulumi/aws, @pulumi/azure, @pulumi/kubernetes in our org, so it's easy to know what's Pulumi specific.


## Reference Links

* [Pulumi on k8s overview](https://www.pulumi.com/kubernetes/)
* [Getting Started](https://pulumi.io/quickstart/)
* [blog on managing Kubernetes with pulumi](https://blog.pulumi.com/program-kubernetes-with-11-cloud-native-pulumi-pearls)
* [GKE tutorial](https://pulumi.io/quickstart/gcp/tutorial-gke.html)
* [Velero on Pulumi](https://github.com/joeduffy/velero/blob/joeduffy/pulumify/conf/README.md)
* [ConfigMap rollout](https://pulumi.io/quickstart/kubernetes/tutorial-configmap-rollout.html)
* Pulumi Social
    * [Twitter](https://twitter.com/pulumicorp)
    * [Community Slack](https://slack.pulumi.io/)
    * [API & Docs](https://pulumi.io)