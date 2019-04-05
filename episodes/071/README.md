# Episode 71 : Exploring the Open Policy Agent (OPA)

- Hosted by @jbeda
- 04/05/2019

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=QU9BGPf0hBw
" target="_blank"><img src="http://img.youtube.com/vi/QU9BGPf0hBw/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:01:00 - Week in Review
- 00:11:00 - OPA, starting with how we should pronounce it
-

## Week in Review
- [Great article highlighting the work migrating kubernetes to go modules](https://medium.com/programming-kubernetes/using-go-modules-with-kubernetes-api-and-client-go-projects-2f3fdd5589a)
- No major Kubernetes release things to report, the cycle starts in earnest on Monday
    - Please check [this kubectl CVE](https://discuss.kubernetes.io/t/announce-security-release-of-kubernetes-kubectl-potential-directory-traversal-releases-1-11-9-1-12-7-1-13-5-and-1-14-0-cve-2019-1002101/5712)!
        - See also [a write up from twistlock on this CVE](https://www.twistlock.com/labs-blog/disclosing-directory-traversal-vulnerability-kubernetes-copy-cve-2019-1002101/).
    - Please check this [CNI CVE](https://discuss.kubernetes.io/t/announce-security-release-of-kubernetes-affecting-certain-network-configurations-with-cni-releases-1-11-9-1-12-7-1-13-5-and-1-14-0-cve-2019-9946/5713) too!
- [Anatomy of CVE-2019-5736: A runc container escape!](https://aws.amazon.com/blogs/compute/anatomy-of-cve-2019-5736-a-runc-container-escape/)
- [How I passed the CKA](https://medium.com/@krystiannowaczyk/how-i-passed-the-cka-certified-kubernetes-administrator-exam-f94b11566528)
- [Grafana Logging using Loki](https://blog.giantswarm.io/grafana-logging-using-loki/)
- [OpenTracing + Open Census](https://medium.com/opentracing/merging-opentracing-and-opencensus-f0fe9c7ca6f0)
    - Also check out [zipkin](https://zipkin.io/)
- [K8s Clusters? Oh Biff ‘em Popeye!](https://medium.com/@fernand.galiana/k8s-clusters-oh-biff-em-popeye-637e9312963)

## Show Notes

TIL: It's "Opa" as in "Opa!" when you go to a Greek restaurant, not "Oh P Aye"

Notes from Torin:
> I recommend getting started with the Kubernetes admission control tutorial on openpolicyagent.org: https://www.openpolicyagent.org/docs/kubernetes-admission-control.html. The tutorial takes you through deploying OPA from scratch, enabling the webhook, deploying a policy, exercising the policy, updating the policy, etc. This is how most people try out OPA for admission control for the first time.
>
> In addition to the tutorial, we recently put together a primer/guide on writing admission control policies: https://github.com/open-policy-agent/opa/blob/master/docs/content/docs/guides-kubernetes-admission-control.md. The guide goes into detail on how admission control policies are structured and evaluated.
>
> A couple other pointers that might be helpful:
>
> * [Language Reference](https://www.openpolicyagent.org/docs/language-reference.html). This page lists all the various built-in functions for things like string manipulation, regexp, etc.
>
> * [Rego Cheatsheet](https://github.com/open-policy-agent/opa/blob/master/docs/content/docs/rego-cheatsheet.md). Compact reference doc for common expressions in OPA (e.g., looking up values, iterating over arrays, etc.)
>
> Lastly, when it comes to writing Rego, you could check out the [VS Code extension](https://github.com/tsandall/vscode-opa) and [play.openpolicyagent.org](https://play.openpolicyagent.org/). Both provide ability to evaluate the policies interactively.
>
> As far as Gatekeeper goes, we're in the middle of restructuring the controller quite a bit and we don't have the same kind of docs yet. The policy authoring experience is more or less the same in Gatekeeper (i.e., users write `deny` rules that enforce constraints on `input` which gets bound to the AdmissionReview object.) Once Gatekeeper is ready, there will be additional features like a constraint library and audit capability that might be good fodder for a future episode :D

## Reference Links
- [tgik issue: 135](https://github.com/heptio/tgik/issues/135)
- Admission Controllers:
    - https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#admission-webhooks
    - https://banzaicloud.com/blog/k8s-admission-webhooks/
- https://www.cncf.io/blog/2019/04/02/toc-votes-to-move-opa-into-cncf-incubator/
- https://www.openpolicyagent.org/
- https://www.openpolicyagent.org/docs/
- https://github.com/open-policy-agent/gatekeeper
- https://github.com/open-policy-agent/opa
- In depth perf article: https://itnext.io/optimizing-open-policy-agent-based-kubernetes-authorization-via-go-execution-tracer-7b439bb5dc5b?gi=414d224ebdc7
- https://db-blog.web.cern.ch/blog/lukas-gedvilas/2018-02-creating-tls-certificates-using-kubernetes-api