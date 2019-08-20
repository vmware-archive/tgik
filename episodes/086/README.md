# Episode 086 : Grokking Kubernetes : kubelet

- Hosted by @mauilion
- 08/09/2019

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=CKpSyl5vgK8
" target="_blank"><img src="http://img.youtube.com/vi/CKpSyl5vgK8/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

- [John Harris on sudo like access with kubectl](https://johnharris.io/2019/08/least-privilege-in-kubernetes-using-impersonation/)
- Make sure you check out [CVE-2019-11247](https://www.stackrox.com/post/2019/08/how-to-remediate-kubernetes-security-vulnerability-cve-2019-11247/)
- https://github.com/trailofbits/audit-kubernetes
- The Kubernetes 3rd Party Security Assessment is [now available](https://www.cncf.io/blog/2019/08/06/open-sourcing-the-kubernetes-security-audit/)
- Ian Coldwater is on the [latest episode](https://kubernetespodcast.com/episode/065-attacking-and-defending-kubernetes/) of the k8s podcast from Google talking about attacking and defending Kubernetes.
- Vito Botta is doing a comparison of a [bunch of k8s storage solutions](http://vitobotta.com/2019/08/06/kubernetes-storage-openebs-rook-longhorn-storageos-robin-portworx/)
- [Tools and methods for auditing RBAC policies](https://www.nccgroup.trust/us/about-us/newsroom-and-events/blog/2019/august/tools-and-methods-for-auditing-kubernetes-rbac-policies/)
- [Building a Kubernetes platform at Pinterest](https://medium.com/pinterest-engineering/building-a-kubernetes-platform-at-pinterest-fb3d9571c948)
- Tim Hockin did some slides on [reconciliation](https://speakerdeck.com/thockin/kubernetes-what-is-reconciliation). 

## Show Notes
- Show Outline:
    - [x] The parts and the [config reference](https://kubernetes.io/docs/reference/#config-reference)
        
        - [x] kubelet
            - [x] client/server auth
            - [x] cri and all the other c\*i integrations
            - [x] kubelet api
            - [x] configz
            - [x] metrics
            - [x] theory of operation
            - [x] static pods
            - [x] PLEG
              
        - [ ] kube-proxy
            - [ ] auth to apiserver
            - [ ] iptables
            - [ ] ipvs
            - [ ] services
            - [ ] theory of operation
            - [ ] configz
            - [ ] metrics
        
        - [ ] kube-controller-manager
            - [ ] auth to apiserver
                - [ ] per controller!
            - [ ] control loops!
            - [ ] can be skipped!
            - [ ] theory of operation
            - [ ] leader election!
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