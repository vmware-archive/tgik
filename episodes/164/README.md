# Episode 164: Ingress dive in and some L7 stuff

- Hosted by @rikatz @jayunit100 @stevesloka
- Recording date: 2021-08-24

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url 

<a href=https://www.youtube.com/watch?v=l3TWbrWkVzY"></aref>
--> 

Show notes: https://hackmd.io/ECpQYlYCQX6h5w9yB02RHw
Scrib: 


## Week in Review

- sig-windows-dev-tools now has working CNI providers! https://github.com/kubernetes-sigs/sig-windows-dev-tools
- Ingress NGINX reaches v1.0.0 GA and drops support for Ingress API v1beta1
- https://github.com/kubernetes-sigs/kpng now has a dual stack implementations for iptables
- https://github.com/kubernetes/website/pull/29504 - improvements in Kubernetes Docs/sites generation time of more than 80%
- https://medium.com/@LachlanEvenson/mutating-kubernetes-resources-with-gatekeeper-3e5585d49ead - Gatekeeper now supports Mutating requests (instead of just validating) and Lachie wrote about this

## Table of Contents
0:00 - Introduction
3:00 - Week in Review
8:25 - SIG Docs is looking for help!
11:10 - The high level architecture of Ingress
15:40 - Getting Started with Contour
24:40 - Shoutout to Curiefense WAF
28:10 - Coming back to Contour and understanding its architecture
34:00 - Going deeper with complex Ingress Objects
38:30 - Watching what happens in Envoy
57:00 - Trying Gateway API
1:19:30 - About SCTP (going back to L4)
 
## Show Notes
* https://gateway-api.sigs.k8s.io/
* https://github.com/kubernetes-sigs/gateway-api/releases/tag/v0.4.0-rc1
* https://www.curiefense.io/
* https://gateway-api.sigs.k8s.io/v1alpha1/guides/getting-started/#installing-gateway-api-crds-manually
