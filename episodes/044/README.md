# Episode #44 : Knative

- Hosted by @jbeda 
- Recording Date: 20180727
- [Live Notes](https://hackmd.io/bzjUvaJRTjGUnLsB0oCo5w#) - Final version of notes will be pushed to this repo after the brodcast.

<iframe width="560" height="315" src="https://www.youtube.com/embed/n_zqedVM0oM?rel=0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

## Table of Contents

Note: Timestamps are [hh:mm] into the stream according to the local time, this might not match the final YouTube video timestamp exactly, but should get you close to the subject.

- 00:00:00 - Welcome to TGIK!
- 00:04:22 - Week in Review
- 00:14:00 - Introduction into Knative 
- 00:19:11 - Getting involved in the Knative Community
- 00:23:00 - Getting invoved in Knative Development 
- 00:30:30 - Installation Preparation and YAML Inspection
- 00:46:30 - Installing Knative
- 01:16:30 - Installing Knative (for real this time!)
- 00:00:00 - Success!
- 00:00:00 - Concluding Thoughts

## News and Announcements

- k8s wins [most impact award at OSCON](https://kubernetes.io/blog/2018/07/19/kubernetes-wins-2018-oscon-most-impact-award/)
    - What happens to the trophy? Passed around like the Stanley Cup? Paris thinks CNCF has it? 
- Kubernetes is now in the stable channel in [Docker Desktop](https://blog.docker.com/2018/07/kubernetes-is-now-available-in-docker-desktop-stable-channel/)
- k8s-related Announcements from Google Next
    - [GKE On-Prem](https://cloud.google.com/gke-on-prem/)
        - In early access
        - Based on VMware for provisioning
    - Knative (today's discussion!)
- [Kubecon CFP closes in ~2 weeks (August 12)](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2018/program/call-for-proposals-cfp/)
    - You don't need to be an expert, share your experience wherever you are in your Kubernetes journey. 
    - Kate Kuchin [gave a talk](https://www.youtube.com/watch?v=X2icFkQMg60) on her experience coming from front-end development that a lot of people appreciated.

## Knative
(pronounced kay-native)
- [Homepage](https://cloud.google.com/knative/)
- [Github](https://github.com/knative/)

### Community

- [Slack](https://knative.slack.com/)
    - Eventually join at [slack.knative.dev](https://slack.knative.dev)
    - For now, use [this link](https://knative.slack.com/join/shared_invite/enQtNDA1ODA1MTc2NjE1LWYxYjhhNDVjOWY1ZTMzMzMwZjE4YTA5YWUxODBmYWQzZTNlMjc0YzlhNmZhYmNhMTk4MzExOTU5MjU5MDcxNzA)
- [knative-users google group](https://groups.google.com/forum/#!forum/knative-users). There is nothing here but this is important so you get access to shared documents.
    - There is also a [knative-dev](https://groups.google.com/forum/#!forum/knative-dev) group, you can apply to join and to get invites to their community meetings. 

### Basic Information

- Is based on Istio as the ingress.
    - You can't really swap that for nginx or kong at the moment.
- Where are we at really vs. the marketing
    - Built on Kubernetes and Istio, part of Google Cloud.
        - Because it's built on k8s it can deploy in lots of different places
        - Over time you'll get the value of the network effect of being able to run in many places
        - A quote from one of the GCP Next sessions on Knative: "Extensible at the top and pluggable at the bottom"
    - What does serverless mean in the knative context?
- Operator vs. developer friendliness
    - Advertised as operator friendly
- [Knative Logo](https://avatars2.githubusercontent.com/u/35583233?s=200&v=4)
    - 7 sided logo with a K in it, with a 9 sided logo with an n in it.
    - "7 of 9" inside joke. 

### Installation

- Used to be per-provider, PR submitted to make them generic Kubernetes instructions
- Unfortunately EKS doesn't have the mutating admission controller flag enabled to have this work out of the box as KNative relies on [Istio automatic sidecar injection](https://istio.io/docs/setup/kubernetes/sidecar-injection/), work in progress, beta, etc.
- Pro-tip: Pasting files into `kubectl apply` is like curling to bash, so you should check out the yaml files before blindly pasting them into a cluster. 
- Can this be installed on on-prem install? You should be able to.
- Pablo recommended using helm to install istio
    - Joe doesn't normally install Tiller as it's difficult in a secure way.

#### Components

- Lots of large YAML files, but most of them for Knative itself is for monitoring. Look at the `-no-mon.yaml` files for smaller installs
    - Matthew Moore: We're splitting it. "Build is required for
      serving today, we have plans for looser coupling post 0.1"
    -  "part of splitting off monitoring is putting the default flow on a diet"
- We really need something like MPM to manage this
- This is a complex beast to install and deploy. 
  The user experience can be great, but we need to pay attention
  to making this stuff work and maintaining it over time and
  giving administrators insight into what is going on.

##### Istio
- Contains:
    - Galley: Replaces Istio-auth
    - statsd-prom-bridge: Envoy only pushes statsd metrics, so
      need something to for Prometheus to scrape.
- Manifest doesn't include RAM requests, but is it in its own 
  namespace so will not have quotas set.
- Require mutating admission controller enabled on the API server
    - `MutatingAdmissionWebhook` added to `--enable-admission-plugins` (for 1.10+)
- After installation, we run `kubectl label namespace default istio-injection=enabled`
    - This will have Istio modify all pods deployed to the
      default namespace to inject the Istio sidecars.

##### Knative
- Installs fluentd and node-exporter as a daemonset on all nodes.
    - We need to figure out how to integrate with people's 
      existing installations.
- [Service (services.serving.knative.dev) is very confusing](https://github.com/knative/serving/issues/1397)

### Concerns for multi-tenancy
Want to lock down a domain names to namespaces, so that
different tenants cannot impersonate each other. Istio
doesn't currently allow you to lock it down.

It would be great to swap Istio for other systems like
Contour that give controls for multi-tenancy.

## Concluding Thoughts

Istio brings value, but it's a lot of complexity which make it
a lot harder from an operational pov. It's great that knative
is made up of building blocks that come together.

It'd be great to separate some of this more, and allow, 
for example, scale-to-zero to be handled by something other
than Istio?

This will need more loose coupling, but Knative makes a great
start.

## Reference Links 

- [Distributing Software in a Massively Parallel Environment](https://www.youtube.com/watch?v=hJAIL-Syzxw) - Dinah McNutt, Google, Inc. Presented at LISA14
- [Maintaining Consistency in a Massively Parallel Environment](https://www.youtube.com/watch?v=_uJlTllziPI) - Dinah McNutt, Google, Inc. Presented at the 2013 USENIX Configuration Management Summit (UCMS '13)

