# Episode 113 : Kubernetes Secrets Take 3

- Hosted by @joshrosso
- 4/10/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=an9D2FyFwR0
" target="_blank"><img src="http://img.youtube.com/vi/an9D2FyFwR0/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://youtu.be/an9D2FyFwR0)
- [00:04:41 - Week in Review](https://youtu.be/an9D2FyFwR0?t=273)
- [00:23:01 - Episode Overview](https://youtu.be/an9D2FyFwR0?t=1381)
- [00:23:33 - Default Secret Behavior](https://youtu.be/an9D2FyFwR0?t=1413)
- [00:41:06 - Encryption at Rest](https://youtu.be/an9D2FyFwR0?t=2466)
- [00:58:41 - External Provider Trade-Offs](https://youtu.be/an9D2FyFwR0?t=3521)
- [01:05:31 - Integration with Secrets: App Level vs Platform Level](https://youtu.be/an9D2FyFwR0?t=3931)
- [01:09:02 - Vault Secret Injection via Agent](https://youtu.be/an9D2FyFwR0?t=4142)
- [01:39:31 - Sealed Secrets](https://youtu.be/an9D2FyFwR0?t=5971)
- [01:46:09 - CSI Secret Driver](https://youtu.be/an9D2FyFwR0?t=6369)
- [01:51:46 - Wrap-Up](https://youtu.be/an9D2FyFwR0?t=6677)

## Week in Review

### Core k8s

- [k8s 1.18.1 is out](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#v1181)
  - Here are all the blog entries you should check out for 1.18 features:
    - [API Priority and Fairness](https://kubernetes.io/blog/2020/04/06/kubernetes-1-18-feature-api-priority-and-fairness-alpha/)
    - [Server Side Apply](https://kubernetes.io/blog/2020/04/01/kubernetes-1.18-feature-server-side-apply-beta-2/)
    - [Feature Topology Manager](https://kubernetes.io/blog/2020/04/01/kubernetes-1-18-feature-topoloy-manager-beta/)
    - [Ingress API improvements](https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/)
    - [Windows CSI Support](https://kubernetes.io/blog/2020/04/03/kubernetes-1-18-feature-windows-csi-support-alpha/)
- [WG Resource Management](https://groups.google.com/forum/#!topic/kubernetes-dev/3OyFrkQfRTA) has met it's goals and is closing down.
- It's not too late to be involved in 1.19! Check out the [shadow application form](https://forms.gle/ujChMWB1wZZAYY7r6) if you're interested, deadline is 4/14!
- [k8s Office Hours](https://git.k8s.io/community/events/office-hours.md) are April 15th, this next Wednesday:  [EU Session](https://youtu.be/6zk-vYxjQjA) and [West Coast US session](https://youtu.be/DLTU5mL-7aA)

### k8s Ecosystem

- [Migrating zookeeper into Kubernetes](https://product.hubspot.com/blog/zookeeper-to-kubernetes-migration) [HN comments](https://news.ycombinator.com/item?id=22814599)
- [Intro to Service Mesh Interface via the CNCF](https://www.cncf.io/webinars/introduction-to-service-mesh-interface/)
- The Docker-compose [specification](https://github.com/compose-spec/compose-spec/blob/master/spec.md) has been published
- [Kubevious](https://github.com/kubevious/kubevious) - another k8s visualization tool
- ArgoCD has been accepted to the CNCF, here's a blog on [setting up multicluster gitops](https://www.infracloud.io/multicluster-gitops-argocd/)
- [Sidekick](https://github.com/minio/sidekick) - "By attaching a tiny load balancer as a sidecar to each of the client application processes, you can eliminate the centralized loadbalancer bottleneck and DNS failover management."
- Only mildly k8s related but friend of the show and k8s maintainer Ben Elder has an article on making your own [open source virtual background](https://elder.dev/posts/open-source-virtual-background/)
- Article by Steve Wade on [how Mettle uses gitops!](https://itnext.io/how-we-do-gitops-mettle-4cc771a6c029)

## Show Notes

- Old episodes:
    - tgik.io/065
    - tgik.io/066

- [x] Secret Review
    - [x] Default
    - [x] Encryption at Rest
    - [x] KMS (Envelope Encryption)
    - [x] External Provider

- [x] Today's Options
    - [x] Vault Injection
    - [x] Sealed Secrets (kind of different concern)
    - [x] [csi-secret-driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver)


## Reference Links

* CSI Secret Driver: https://github.com/kubernetes-sigs/secrets-store-csi-driver
* Go-Daddy External Secrets: [Kubernetes External Secrets](https://github.com/godaddy/kubernetes-external-secrets)
* Kubevault (by appscode): https://kubevault.com
* Vault injector how-to: https://learn.hashicorp.com/vault/getting-started-k8s/sidecar
* [Hashicorp announcing the vault helm chart](https://www.hashicorp.com/blog/announcing-the-vault-helm-chart/)
* [Hashicorp secrets store csi driver provider vault](https://github.com/hashicorp/secrets-store-csi-driver-provider-vault)
