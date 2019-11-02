# Episode 096 : Grokking Kubernetes : kube-scheduler

- Hosted by @mauilion
- 11/1/2019

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=XxVHNWoZO_c
" target="_blank"><img src="http://img.youtube.com/vi/XxVHNWoZO_c/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

- [Contour 1.0 release!](https://projectcontour.io/announcing-contour-1.0/)
- [VMWorld EU](https://www.vmworld.com/en/europe/index.html) 
- [New Octant.dev website](https://octant.dev/)
- [Pod labels via status](https://github.com/kubernetes/kubernetes/pull/84260)
- [Kube Docs survey](https://kubernetes.io/blog/2019/10/29/kubernetes-documentation-end-user-survey/)
- [A Kubernetes crime story](https://engineering.prezi.com/https-engineering-prezi-com-a-kubernetes-crime-story-2e8d75a77630)
- [How to write a kubectl plugin](https://elsesiy.com/blog/how-to-kubectl-plugin)
- [kubernetespodcast.com](https://kubernetespodcast.com/)
- [Nividia Operator](https://devblogs.nvidia.com/nvidia-gpu-operator-simplifying-gpu-management-in-kubernetes/)



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
- [kubectl sudo like auth](https://johnharris.io/2019/08/least-privilege-in-kubernetes-using-impersonation/)
- [keps](https://github.com/kubernetes/enhancements)
- [Tim Hockin reconciliation presentation](https://speakerdeck.com/thockin/kubernetes-what-is-reconciliation)
- [kube-csr](https://github.com/JulienBalestra/kube-csr)
- [Bryan's Webinar on Octant](https://www.youtube.com/watch?v=D5PLsXP9aPc)
