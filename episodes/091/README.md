# Episode 091 : Kpack

- Hosted by @jbeda
- 09/20/19

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url 4zkRX9PSJ5k--->

<a href="https://www.youtube.com/watch?v=4zkRX9PSJ5k" target="_blank"><img src="https://i.ytimg.com/vi/4zkRX9PSJ5k/maxresdefault.jpg" border="50" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:19:43 - Starting in on kpack

## Week in Review

- Tidbits from the Kubernetes community:
  - Kubernetes 1.16 is has [gone GA](https://kubernetes.io/blog/2019/09/18/kubernetes-1-16-release-announcement/)
      - Reminder that [etcd 3.4](https://kubernetes.io/blog/2019/08/30/announcing-etcd-3-4/) is also here
      - If all your Helm charts broke then you missed all the API warnings on this show the past three months. :D Check [here](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/) for more information on the removal of deprecated APIs in 1.16.
      - TGIK Episode 84 on [removed APIs](https://www.youtube.com/watch?v=-U79ZLO_37E)
      - [Endpoint slices by Nigel Poulton](https://youtu.be/f3xusisgp74)
  - There's a medium severity [kubectl CVE](https://discuss.kubernetes.io/t/announce-security-release-of-kubectl-versions-v1-16-0-1-15-4-1-14-7-and-1-13-11-cve-2019-11251/7993) for v1.13.10, v1.14.6, and v1.15.3.
  - KubeCon + CloudNativeCon [schedule is up](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2019/schedule/).
      - Don't forget to check out the [Rejekt's Conference](https://cloud-native.rejekts.io/) prior to KubeCon
  - Registration for the Kubernetes Contributor Summit for San Diego [is open](https://kubernetes.io/blog/2019/09/24/san-diego-contributor-summit/)
  - k8s 1.17 [planning has begun](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.17)
      - k8s 1.16.1 is slated for October 2.
  - We're looking for experts to help our incoming new users on [discuss.k8s.io](https://discuss.kubernetes.io/t/help-someone-out-here-win-some-prizes/7877) - help new users out for a chance to win a k8s shirt!
- CNAB [has gone 1.0](https://github.com/deislabs/cnab-spec/blob/master/100-CNAB.md)
    - [Microsoft's Blog](https://cloudblogs.microsoft.com/opensource/2019/09/10/cloud-native-application-bundle-cnab-1-0-updates/) on the release


- From around the web:
    - Kubecost has a [A Practical Guide to Setting Kubernetes Requests and Limits ](http://blog.kubecost.com/blog/requests-and-limits/)
    - Gergely Brautigam has written [Using a Kubernetes based Cluster for Various Services with auto HTTPS](https://skarlso.github.io/2019/09/21/kubernetes-cluster/)
    - Karthik Sirasanagandla has published his [preparation checklist](https://blog.codonomics.com/2019/09/ckad-exam-preparation-checklist.html) on taking the CKAD.
    - @feloy has been working on putting the [man pages for Kubernetes online](https://kubernetes-manpages.feloy.dev/cgi-bin/man/man2html)



## Reference Links
* [kpack blog post](https://content.pivotal.io/blog/introducing-kpack-a-kubernetes-native-container-build-service)
* [kpack github](https://github.com/pivotal/kpack)
* [Cloud Native Buildpacks](https://buildpacks.io)
* [kpack tutorial](https://github.com/pivotal/kpack/blob/master/docs/tutorial.md)