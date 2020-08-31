# Episode 101 : Grokking Kubernetes : kube-apiserver

- Hosted by @mauilion
- Recording date: 2020-01-17

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=ZB4LhAeiTCE
" target="_blank"><img src="http://img.youtube.com/vi/ZB4LhAeiTCE/hqdefault.jpg"/></a>
## Table of Contents

- 00:00:00 - Welcome to TGIK!![]()

- 00:00:00 - Week in Review

## Week in Review

- Kubernetes now has a [Bug Bounty Program](https://kubernetes.io/blog/2020/01/14/kubernetes-bug-bounty-announcement/)!
- Kubernetes holds [lots of meetings](https://twitter.com/castrojo/status/1217885251755814912)
   - Stuff from the [now monthly Kubernetes Community Meeting](https://discuss.kubernetes.io/t/kubernetes-community-meeting-notes/35/73)
       - Thursday, March 05: Week 9 - Code Freeze
       - Monday, March 16: Week 11 - Docs must be completed and reviewed
       - Tuesday, March 24: Week 12 - Kubernetes v1.18.0 released
       - Patch Release Updates https://git.k8s.io/sig-release/releases/patch-releases.md
         - 1.17.1 released Jan. 14
         - 1.16.5 coming today Jan. 16
         - 1.15.8 coming today Jan. 16
         - 1.14.11 coming today Jan.16
       - Check out Josh Berkus and Noah Kantrowitz's [This Week in Kubernetes Development](http://lwkd.info/2020/20200113) for more summaries of what's going on. (Subscribe to this, you won't regret it. --jorge)
- [Crossplane](https://github.com/crossplaneio/crossplane) has turned [0.6](https://blog.crossplane.io/crossplane-v0-6-enabling-application-delivery-platforms-on-the-road-towards-production-ready/)
- [Header and Host Rewrite with Contour 1.1](https://projectcontour.io/announcing-contour-1.1/)
- Michael Hausenblas brings you more [Kubernetes on Rpi's](https://mhausenblas.info/kube-rpi/)
- Duffie loves to see people trying bare metal, here's how [thebsdbox does it](https://thebsdbox.co.uk/2020/01/02/Designing-Building-HA-bare-metal-Kubernetes-cluster/).
- [KubeVault](https://kubevault.com/) is now 0.3. Overview on [ByteBuilder](https://blog.byte.builders/post/kubevault-v0.3.0/).
- [Werf goes 1.0](https://medium.com/flant-com/announcing-werf-1-0-stable-813b664a06ae) - A gitops CLI tool from Flant.
- A collection of documentation, how-tos, tools and other information on debugging and identifying Kubernetes/container workload failures, performance and reliability considerations, and other [kubernaughties](https://github.com/jnoller/kubernaughty/blob/master/docs/part1-introduction-and-problem-description.md). 
    

## Show Notes

- Outline:
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
            - [x] metrics
        
        - [x] kube-scheduler
            - [x] https://kubernetes.io/docs/concepts/scheduling/kube-scheduler/
            - [x] auth to apiserver
            - [x] theory of operation
            - [x] https://kubernetes.io/docs/concepts/extend-kubernetes/poseidon-firmament-alternate-scheduler/
            - [x] :warning:can be skipped!:warning: (So not a security thing!)
            - [x] leader election!
            - [x] direct scheduling
            - [x] multiple schedulers!
            - [x] configurable!
            - [x] metrics
    
        - [x] kube-apiserver
            - [x] auth to etcd
                - [x] Other neat etcd tricks!
            - [x] auth to kubelet
            - [x] How the apiserver handles authentication [docs](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)
                - [x] certs, jwts, tokens
                - [x] Kubernetes components
                - [x] People and bots
            - [x] How the apiserver handles authorization [docs](https://kubernetes.io/docs/reference/access-authn-authz/authorization/)
                - [x] 
            - [x] Admission Control [docs](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)
                - [x] builtins
                - [x] PSP
            - [x] Exploring the API.
                - [x] kubectl explain
                - [x] openapi stuffs
                - [x] deprecation [docs](https://kubernetes.io/docs/reference/using-api/deprecation-policy/)
                - [x] [api reference docs](https://kubernetes.io/docs/reference/using-api/api-concepts)
        
                - [ ] validating and mutating webhooks [docs](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/)
                - [ ] static pods
            - [x] theory of operation
                - [ ] access patterns.![https://i.imgur.com/mxwlWPe.png](https://i.imgur.com/mxwlWPe.png) 
            - [ ] ha [docs](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/high-availability/)
                - [ ] internal apiserver lb
                - [ ] external apiserver lb
            - [ ] kubectl get --raw /metrics
    - [ ] General systems stuff.
        - [ ] show direct scheduling!
        - [ ] The ways that each are configured/configurable.
        - [ ] The access patterns for each.
        - [ ] The authentication mechanisms for each.
        - [ ] What even is edge vs level triggered?
        - [ ] What is a watch?
        - [ ] Why is all this so darn stable?

## Reference Links
- [Alibaba Cloud with 10,000 Nodes k8s cluster](https://medium.com/@Alibaba_Cloud/how-does-alibaba-ensure-the-performance-of-system-components-in-a-10-000-node-kubernetes-cluster-ff0786cade32)
- [how kubectl exec works](https://erkanerol.github.io/post/how-kubectl-exec-works/)
- [Awesome micro-framework to get started quickly with admission control](https://github.com/elithrar/admission-control) 
- [Postman Desktop Client](https://www.getpostman.com/downloads/)
