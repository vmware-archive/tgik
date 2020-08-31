# Episode 099 : Grokking Kubernetes : kube-apiserver

- Hosted by @mauilion
- Recording date: 2019-12-13

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=582HunYxrgY
" target="_blank"><img src="https://i.imgur.com/huW7Okj.jpg" width="480" height="360" border="10" /></a>
## Table of Contents

- 00:00:00 - Welcome to TGIK!![]()

- 00:00:00 - Week in Review

## Week in Review

- [1.17 is out!](https://kubernetes.io/blog/2019/12/09/kubernetes-1-17-release-announcement/)
    - Check out [the release retrospective](https://www.youtube.com/watch?v=XQaC5ke9SHc) if you're interested in seeing how the sausage is made
- [There is a performance regression](https://github.com/kubernetes/kubernetes/pull/85810)
- [to try 1.17 in kind!](https://twitter.com/BenTheElder/status/1204184462088519686?s=20)
- [kubernetes podcast covers the release with Guin!](https://kubernetespodcast.com/episode/083-kubernetes-1.17/)
- [Azure get's an ALB](https://azure.microsoft.com/en-au/blog/application-gateway-ingress-controller-for-azure-kubernetes-service/)
- [Great doc by Daniele Polencic on troubleshooting deployments](https://learnk8s.io/troubleshooting-deployments)
- [kube-state metrics new release!](https://twitter.com/tariq1890/status/1204121224428711936?s=20) [release notes](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.9.0-rc.0)
- [Stackrox dives into new features in 1.17](https://www.stackrox.com/post/2019/12/whats-new-in-kubernetes-1.17-a-deeper-look-at-new-features/)
- [My presentation at blackhat with Ian Coldwater is finally up!](https://twitter.com/IanColdwater/status/1205350978360414209?s=20)
- [The Podlets!](https://thepodlets.io/)
- [Sidecar containers](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md)

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
            - [ ] Admission Control [docs](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)
                - [ ] builtins
                - [ ] PSP
            - [ ] Exploring the API.
                - [ ] kubectl explain
                - [ ] openapi stuffs
                - [ ] deprecation [docs](https://kubernetes.io/docs/reference/using-api/deprecation-policy/)
                - [ ] [api reference docs](https://kubernetes.io/docs/reference/using-api/api-concepts)
        
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
