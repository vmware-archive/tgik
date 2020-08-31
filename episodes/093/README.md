# Episode 093 : Grokking Kubernetes : kube-controller-manager

- Hosted by @mauilion
- Recording date: 2019-10-11

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=78XQe6FdHac
" target="_blank"><img src="http://img.youtube.com/vi/78XQe6FdHac/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

- TGIK Git repo has [moved!](https://github.com/vmware-tanzu/tgik)

- [Announcing Kubernetes Community Days](https://www.cncf.io/blog/2019/10/03/announcing-kubernetes-community-days/)

- Schedule for all events in San Diego are up! Register for those sessions you really want to attend!
    - [contributor summit](https://events19.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/program/schedule/)
    - [kubecon 2019](https://events19.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2019/schedule/)
    - [Cloud Native Rejekts](https://cfp.cloud-native.rejekts.io/cloud-native-rejekts-na-2019/schedule/)
    - [Cloud Native Security Day](https://events19.linuxfoundation.org/events/cloud-native-security-day-2019/program/schedule/)

- CFP for kubecon EU 2020 is [OPEN](https://linuxfoundation.smapply.io/prog/kccnceu2020/)
    - CFP Close: 11:59PM PDT, Wednesday, December 4, 2019

- Jorge was interviewed on the [Kubernetes Podcast!](https://kubernetespodcast.com/episode/074-community-and-contributor-experience/)

- A note from Joe on [Open Source in VMware Tanzu!](https://blogs.vmware.com/cloudnative/2019/10/01/open-source-in-vmware-tanzu/)

- [Sloop](https://github.com/salesforce/sloop)

- a great writeup on [Clusters at scale with eBay](https://tech.ebayinc.com/engineering/scalability-tuning-on-tess-io-cluster/)

- [A PR to migrate leader election to the lease api](https://github.com/kubernetes/kubernetes/pull/81030)



## Show Notes
- Show Outline:
    - [x] The parts and the [config reference](https://kubernetes.io/docs/reference/#config-reference)
        
        - [x] kubelet tgik.io/086
            - [x] client/server auth
            - [x] cri and all the other c\*i integrations
            - [x] kubelet api https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/server/server.go
            - [x] configz
            - [x] metrics
            - [x] theory of operation
                - [x] For exec: https://erkanerol.github.io/post/how-kubectl-exec-works/
            - [x] static pods
            - [x] PLEG
              
        - [x] kube-proxy tgik.io/090
        - https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/
            - [x] auth to apiserver
            - [x] iptables
            - [x] ipvs
            - [x] services
            - [x] theory of operation
                - https://github.com/kubernetes/enhancements/issues/536
                - https://kubernetes.io/docs/tutorials/services/source-ip/
            - [x] configz
            - [x] metrics
        
        - [x] kube-controller-manager tgik.io/093
            - [x] [init code](https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-controller-manager/app/controllermanager.go#L373) 
            - [x] [controller code](https://github.com/kubernetes/kubernetes/tree/master/pkg/controller)
            - [x] control loops!
                - attachdetach
                - **bootstrapsigner**
                - cloud-node-lifecycle
                - clusterrole-aggregation
                - cronjob
                - csrapproving
                - csrcleaner
                - csrsigning
                - daemonset
                - deployment
                - disruption
                - endpoint
                - **endpointslice**
                - garbagecollector
                - horizontalpodautoscaling
                - job
                - namespace
                - nodeipam
                - nodelifecycle
                - persistentvolume-binder
                - persistentvolume-expander
                - podgc
                - pv-protection
                - pvc-protection
                - replicaset
                - replicationcontroller
                - resourcequota
                - root-ca-cert-publisher
                - route
                - service
                - serviceaccount
                - serviceaccount-token
                - statefulset
                - **tokencleaner**
                - ttl
                - ttl-after-finished
            - [x] disabled by default:
                - bootstrapsigner
                - endpointslice
                - tokencleaner
            - [x] :warning:can be skipped!:warning:
            - [x] theory of operation
            - [x] leader election!
            - [ ] metrics
        
        - [ ] kube-scheduler
            - [ ] auth to apiserver
            - [ ] theory of operation
            - [ ] direct scheduling
            - [ ] not to much magic here.
            - [ ] extensible tho
            - [ ] leader election!
            - [ ] metrics
    
        - [ ] kube-apiserver
            - [ ] auth to etcd
            - [ ] auth to kubelet
            - [ ] access patterns.
            - [ ] theory of operation
            - [ ] ha
            - [ ] kubectl get --raw /metrics
            - [ ] Exploring the API.
                - [ ] kubectl explain
                - [ ] openapi stuffs
                - [ ] [api reference docs](https://kubernetes.io/docs/reference/using-api/api-concepts)
        
    - [ ] General systems stuff.
        - [ ] show direct scheduling!
        - [ ] The ways that each are configured/configurable.
        - [ ] The access patterns for each.
        - [ ] The authentication mechanisms for each.
        - [ ] What even is edge vs level triggered?
        - [ ] What is a watch?
        - [ ] Why is all this so darn stable?

## Reference Links
- [keps](https://github.com/kubernetes/enhancements)
- [Tim Hockin reconciliation presentation](https://speakerdeck.com/thockin/kubernetes-what-is-reconciliation)
- [kube-csr](https://github.com/JulienBalestra/kube-csr)
