# Episode 147: CoreDNS

- Hosted by Luke Short
    - Twitter: [@ekultails](https://twitter.com/ekultails)
    - LinkedIn: https://www.linkedin.com/in/luke-short-68872813a/
    - Blog: https://blog.ekultails.com/
    - GitHub: https://github.com/ekultails
    - Technical notes: 
        - Stable: https://ekultails.github.io/rootpages/
        - Latest: http://k8s.ekultails.com/rp/
- 03/05/2021

<a href="https://www.youtube.com/watch?v=44wv-d7H-HE
" target="_blank"><img src="http://img.youtube.com/vi/44wv-d7H-HE/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents (Timestamps)

- 00:00:40 - Welcome to TGIK!
- 00:06:03 - Week in Review
- 00:24:55 - Katacoda Kubernetes Playground
- 00:26:33 - CoreDNS Introduction
- 00:30:56 - CoreDNS Plugins
- 00:32:38 - CoreDNS Internal Plugin: File
- 00:34:28 - CoreDNS DNS Records
- 00:37:19 - Story time! Luke's first container was a DNS server
- 00:38:44 - Kubernetes at Scale: https://openai.com/blog/scaling-kubernetes-to-7500-nodes/
- 00:39:34 - Chat - What DNS servers do you use?
- 00:41:38 - CoreDNS Pods on Kubernetes
- 00:43:57 - Luke making spelling mistakes :-)
    - Luke thought he was being trolled by him resizing the window. Sometimes that will move the cursor around in the Kubernetes Playground in unpredictable ways.
- 00:45:10 - Moving to Luke's home Kubernetes cluster
    - Configuring CoreDNS in Kubernetes
- 00:49:28 - Load balancing with DNS
- 00:51:48 - View all CoreDNS plugins 
    - Preview the "records" and "unbound" plugins
    - Authoritative vs recursive
- 00:54:32 - CoreDNS external plugins require recompilation
- 00:55:55 coredns-unbound project
- 00:56:24 Public container registries
- 00:59:29 How to compile CoreDNS with plugins
- 01:01:22 CoreDNS demo of customizing Corefile on Kubernetes
- 01:21:18 CoreDNS demo of adding customized DNS records
- 01:40:09 Documentation for customizing CoreDNS in Kubernetes
- 01:43:29 Outro

## Week in Review

### Core K8s

- [Kubernetes 1.21 beta releasing this upcoming Tuesday March 9, 2021](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.21#timeline)
    - [Cluster API being updated to v1alpha4, preparing for release](https://cluster-api.sigs.k8s.io/roadmap.html#v04-v1alpha4--q4-2020)
    - [Service API LoadBalancer will no longer require NodePorts (can use VIPs only)](https://github.com/kubernetes/enhancements/issues/1864)
        - May be deffered to 1.22
    - [More than one Service API LoadBalancer is supported!](https://github.com/kubernetes/kubernetes/pull/98277)

### K8s and Cloud Native Ecosystem

- [Google Autopilot provides fully managed Kubernetes](https://thenewstack.io/googles-new-autopilot-for-kubernetes/)
- [Datastax Astra Offers a Serverless Option for Cassandra](https://thenewstack.io/datastax-astra-offers-a-serverless-option-for-cassandra/)
- [RediSearch 2.0 is twice as fast as 1.0](https://thenewstack.io/redis-redisearch-secondary-index-responds-faster-streamlines-indexing/)
- [CIVO Moving from OpenStack to Kubernetes](https://youtu.be/i9SrkkSzDnI?t=939)
- [CIVO #KUBE100 beta (free Kubernetes!) ending soon](https://www.civo.com/pricing)
- [Cluster API - KubeAcademy ](https://kube.academy/courses/cluster-api)

## Show Notes

- [Kubernetes Playground](https://www.katacoda.com/courses/kubernetes/playground)
- [CoreDNS Manual](https://coredns.io/manual/toc/)
- [CoreDNS Internal Plugins](https://coredns.io/plugins/)
- [CoreDNS External Plugins](https://coredns.io/explugins/)
- [List of Top-Level Domains (TLDs)](https://data.iana.org/TLD/tlds-alpha-by-domain.txt)
- [GitHub ekultails/container-coredns-unbound Project](https://github.com/ekultails/container-coredns-unbound)
- [Kubernetes Customizing DNS Service](https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/)

### Exposing DNS Port

DNS needs to listen on a port on a Node. Ingress cannot expose it because only handles HTTP connections.

Networking options:

- Service NodePort = Uses 3000+ port range which is not ideal. DNS needs to listen on the default port of 53 for compatibility with most programs.
- hostNetwork = Does not use CNI. A Pod can see other Pods and programs listening locally on 127.0.0.1.
- hostPort = Uses CNI for IP addressing but exposes specific ports on a Node.

tl;dr? Prefer hostPort for DNS today!

https://medium.com/@maniankara/kubernetes-tcp-load-balancer-service-on-premise-non-cloud-f85c9fd8f43c

Long-term, the Gateway API will provide an alternative to Ingress that will expose ports for any Pod (even those that aren't HTTP).

### Kubernetes Scale

Kubernetes has been tested to run at a scale of up to 7500 nodes!

https://openai.com/blog/scaling-kubernetes-to-7500-nodes/

### DNS Servers

DNS servers used by folks in the TGIK community (not necessarily for Kubernetes):

- bind
- coredns
- dnsmasq
- external-dns
- hosts files
- infoblox
- route53
- systemd-resolve

### Public Container Registries

Container regisitries used by the TGIK community:

- Amazon ECR
- Artifactory
- DTR
- Docker Hub
- Harbor
- Nexus
- Pulp
- Quay
