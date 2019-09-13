# Episode 090 : Grokking Kubernetes : kube-proxy

- Hosted by @mauilion
- 09/13/2019

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=wEyLsEaomfA
" target="_blank"><img src="http://img.youtube.com/vi/wEyLsEaomfA/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

- Friday the 13th. :pray: demo gods

- palindrome week!
    - 9/10/19 P
    - 9/11/19 A
    - 9/12/19 L
    - 9/13/19 I
    - 9/14/19 N
    - 9/15/19 D
    - 9/16/19 R
    - 9/17/19 O
    - 9/18/19 M
    - 9/19/19 E

- Kubernetes 1.16 has been [bumped to Sept 18th](https://groups.google.com/forum/#!topic/kubernetes-dev/06N2zsOT_YY/discussion)
    - [Patch release schedule](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)!

- [Mistake that cost thousands](https://medium.com/@gajus/mistake-that-cost-thousands-kubernetes-gke-2212ea663e1f)

- Google's Kubernetes Podcast is checking out [Windows Server Containers](https://kubernetespodcast.com/episode/070-windows-server-containers/) with Patrick Lang

- CNCF's [cloud native survey](https://www.surveymonkey.com/r/cloudnativesurvey2019) is live!

- [Isopod](https://medium.com/cruise/isopod-5ad7c565d350), a DSL for k8s config

- If you [help users out on discuss k8s you can win a tshirt](https://discuss.kubernetes.io/t/help-someone-out-here-win-some-prizes/7877/)

- We're holding a [webinar on Octant](https://vmware.zoom.us/webinar/register/WN_PCRyynG_RRGE__ljwS7U8Q) next week if you're interested in checking it out.
    - Speaking of webinars did you know CNCF runs [cloud native webinars](https://www.cncf.io/webinars/) all the time?


## Show Notes
- Show Outline:
    - [x] The parts and the [config reference](https://kubernetes.io/docs/reference/#config-reference)
        
        - [x] kubelet
            - [x] client/server auth
            - [x] cri and all the other c\*i integrations
            - [x] kubelet api https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/server/server.go
            - [x] configz
            - [x] metrics
            - [x] theory of operation
                - [x] For exec: https://erkanerol.github.io/post/how-kubectl-exec-works/
            - [x] static pods
            - [x] PLEG
              
        - [x] kube-proxy
        - https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/
            - [x] auth to apiserver
            - [x] iptables
            - [ ] ipvs
            - [x] services
            - [x] theory of operation
            - https://github.com/kubernetes/enhancements/issues/536
            - https://kubernetes.io/docs/tutorials/services/source-ip/
            - [x] configz
            - [x] metrics
        
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