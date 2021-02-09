# Episode 145 : Duck Typing in Kubernetes
- Hosted by @n3wscott
- 02/05/2021

<a href="https://www.youtube.com/watch?v=slX2nAFHeK0
" target="_blank"><img src="http://img.youtube.com/vi/slX2nAFHeK0/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review
- 00:11:20 - We are here to talk about ducks.
- 00:13:50 - Frame the prototype code for the Maths demo.
- 00:15:00 - Kind Cluster booting.
- 00:17:14 - [ko](https://github.com/google/ko) apply the Maths controller.
- 00:18:18 - Apply the starter demo resources.
- 00:24:15 - Take a tour of the Maths controller, api and duck type definition.
- 00:32:45 - Frame the "Square" CRD, start with controller runtime.
- 00:36:18 - ko-ify the controller runtimes output to let us use ko for deploying.
- 00:40:15 - Made a Square CRD but found it does not have a status. The drama begins.
- 00:41:05 - Update our Square CRD with the ResultsType duck type shape.
- 00:43:10 - oh noes, forgotten how to pull a resource out of the controller runtime client using the reconciler key, look it up.
- 00:45:00 - Saved by Ben Moss and friends with a copy/pasta controller runtime client get for the resource. On to implementing the simple controller.
- 00:50:21 - Try out our new Square resource controller with an object!
- 00:51:00 - Wait! remember to update the resource status, CR does not do this for you.
- 00:51:57 - Detour to talk about Knative Style reconcilers for a minute. 
- 00:54:41 - Applied our updated Square instance, but no status update... Let's dig in.
- 00:59:33 - Found a missing status annotation, put in the wrong place. 
- 01:20:30 - We now have the controller updating the status, time to tie it into the original Math controller.
- 01:21:00 - Add the aggregated RBAC for the results duck type. (note there is an issue in the Maths RBAC file that gets fixed soon)
- 01:15:15 - Attempt to show integration with the Maths controller.
- 01:26:36 - Status is gone again on Square! Panic mode!
- 01:29:34 - Correcting the RBAC, still not working, more panic, time to look at some controller logs. :D
- 01:35:30 - WORKING DUCK TYPE!
- 01:37:22 - Let's talk about duck typing.
- 01:38:06 - Per request, detour to poke the demo again. Find that there is still some status update issues from the controller runtime side.  The code still has issues. Fix them!
- 01:47:25 - Back to duck typing talk.
- 01:49:37 - Pitch the API Machinery Duck Type Prototype repo proposal.
- 01:53:00 - Thanks so much! See you on the internet.

## Week in Review

### Core K8s

- [Reminder: v1beta1 versions going away in 1.22: CRD and AdmissionWebhookConfiguration](https://discuss.kubernetes.io/t/reminder-v1beta1-versions-going-away-in-1-22-crd-and-admissionwebhookconfiguration/14684)
- [Kind v0.10 released](https://github.com/kubernetes-sigs/kind/releases/tag/v0.10.0). This is two weeks old now. Missed it last week!
- [Kubernetes 1.21 - Enhancements Freeze](https://kubernetes.slack.com/archives/C2C40FMNF/p1611600310047600): Feb. 9th.
    - Release target date for 1.21: April 8th.

### K8s and Cloud Native Ecosystem

- Occasional TGIK host Tiffany Jernigan wrote up an excellent getting started guide to Velero - https://tanzu.vmware.com/developer/guides/kubernetes/velero-gs/
- [Pinniped v0.5!](https://pinniped.dev/posts/multiple-pinnipeds/).  [We gotta do a TGIK on this sometime soon! -- Joe]
- [Contour v1.12 released](https://projectcontour.io/contour_v1120/).  This release brings local rate limiting, dymamic request headers and header has loadbalancing. Great to seee Contour continue to evolve.
- Controversial view of Helm: [Why Helm never felt like it belonged](https://littlechimera.com/posts/helm-design/) from LittleChimera.
- The core container image distribution system going to the CNCF. [Donating Docker Distribution to the CNCF
](https://www.docker.com/blog/donating-docker-distribution-to-the-cncf/)
- [Exploring the Kubernetes Operator Pattern](https://iximiuz.com/en/posts/kubernetes-operator-pattern/) nice long form blog post with diagrams.
- [Demystifying Google Container Tool Jib](https://ashishtechmill.com/demystifying-google-container-tool-jib-java-image-builder) Someone needs to take jib the rest of the way to a [ko](https://github.com/google/ko)-like experiance. :D
    - [jo](https://github.com/kameshsampath/jo) is a `ko` knock-off for Java.
    - In [TGIK 102](http://tgik.io/102) Joe B. goes over what Ko and Configula are
- Updating container images using just Go code in [this tutorial by ahmetb](https://ahmet.im/blog/building-container-images-in-go/)p
- [Knative in Action](https://www.manning.com/books/knative-in-action) has been released to printers, meaning physical copies will be shipping in a few weeks.
- [kpack 0.2.0](https://github.com/pivotal/kpack/releases) released today with Windows building support, Notary V1 image signing, and better logging.


## Show Notes

Shoutout to ramb.ly for hosting [TGIK8s](ramb.ly/youtube:slX2nAFHeK0) 

Follow-up: [Duck Types for Kubernetes](https://groups.google.com/g/kubernetes-sig-api-machinery/c/S6BTVjIRTlM) thread on API Machinery. Discuss at the next [API Machinery WG meeting](https://goo.gl/0lbiM9), Feb 10th, 11am PST.

Repos used in this talk:
- github.com/n3wscott/maths the starting repo with Add and Subtract resources.
- github.com/n3wscott/multiply the resulting demo repo adding the Square resource.

Kubecon talks on duck typing:
- [Extending Knative for Fun and Profit](https://www.youtube.com/watch?v=Mb8c5SP-Sw0)
- [Polymorphic Reconcilers in Kubernetes](https://www.youtube.com/watch?v=kldVg63Utuw)
