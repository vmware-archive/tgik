# Episode 153 : "KubeProxy and KPNG"
- Hosted by @jayunit100, @rikatz, @mclaseau
- Recording date: 2021-04-30

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=UF1SglWIYz4" target="_blank"><img src="https://i.ytimg.com/vi/UF1SglWIYz4/maxresdefault.jpg" border="10" /></a>

This week we are going to hop around and look at whats goin on w/ Kube proxy:
- the backlog: a good way to get a feel for some of the KP pain points
- the KPNG project: a decoupled, easy to extend KP for the future

## Table of Contents

## Week in Review

VLADs new test framework 

https://github.com/vladimirvivien/e2e-framework/blob/initial-poc/examples/suites/hello2_test.go

- https://github.com/cncf/toc/issues/650
- https://github.com/submariner-io/submariner

INGRESS vs GATEWAY

https://kubernetes.io/docs/concepts/services-networking/ingress/ vs https://kubernetes.io/blog/2021/04/22/evolving-kubernetes-networking-with-the-gateway-api/ 

NETPOL Subproject Status

- https://kubernetes.io/blog/2021/04/20/defining-networkpolicy-conformance-cni-providers/ 
 
- Anything else worth discussing? 

## Show Notes

### Why were going to talk about KPNG


KPNG KEP !

https://github.com/kubernetes/enhancements/pull/2094/files


KUBE PROXY AS A BOTTLENECK

- https://openai.com/blog/scaling-kubernetes-to-7500-nodes/  scaling
- https://github.com/kubernetes/kubernetes/issues/96935 WINDOWS priority
- https://github.com/kubernetes/kubernetes/issues/23864 load balancing
- https://github.com/kubernetes/kubernetes/issues/97052 Technical complexity
- https://github.com/kubernetes/kubernetes/issues/98302 Command config is nested, provider specific, messy
- /SOURCE/go/src/k8s.io/kubernetes/pkg/proxy/ipvs generic logic out of providers
- Lately alot of CNIs are off making their own proxies!

LOOK AT THIS COPY PASTE ! 

- https://github.com/vmware-tanzu/antrea/pull/772/files#diff-f035aa6b882a07149c42554fa347c5983a99d5678d192256b84494e15140d82c
- https://github.com/kubernetes/enhancements/pull/2635

KPNG SOLVES
- DaemonSet problem
- Decoupling/TechDebt problem

### Understanding kube-proxy rules

- https://github.com/jayunit100/k8sprototypes/blob/master/kind/kind-local-up.sh generic calico / antrea cluster...
- annotated iptables trace from calico
  - Jumps start after the point where they left off
  - Ordering is important
  - The Sync function is huge :
- https://github.com/jayunit100/k8sprototypes/blob/master/2020kubecon/iptables-save-calico.txt
- https://www.rkatz.xyz/post/2020-09-13-the-mix-of-packet-filter-linux-part1/

### KPNG Demo

KPNG ARCHITECTURE 

- https://github.com/kubernetes-sigs/kpng/blob/master/doc/arch.svg 

BUILD KIND FROM SOURCE

- https://github.com/kubernetes-sigs/kind/ so you can use `noProxy`
- https://github.com/kubernetes-sigs/kpng/tree/master/hack 
 
GENERIC KPNG CLUSTER SETUP

https://github.com/kubernetes-sigs/kpng/tree/master/hack

## KPNG arch Notes:
- Jobs: all processes are `Jobs`, each Job type is a different struct.   A job is essentially a way
of representing a controller.
- WatchState: WatchState connects the `Sink` with the store2localdiff functionality.
    - Initially empty.
    - Triggered by store2localdiff (store2local) *or* store2globaldiff (store2api or store2file)
    - When triggered...
        - revieves a localnet.v1Service (foo)
        - calculate the diff between empty vs localnet.v1Service 'foo'
        - calls the `SendUpdate`  to the `Sink`
- What is a Sink?
    - GenericSink: Generic impl of the Diff2EventLoop (a utility used by LocalSink and Global Sinc)
    - Sink interfaces which compose w/ the generic sinc:
        - LocalSink: used by store2local
        - GlobalSink: store2globaldiff used by the store2api
- There are 2 ways to run a backend:
    - by specifying it on startup: `kpng kube to-local to-nft` runs everything in local memory.
    - by separating server + client: `kpng kube to-api ` + `kube-proxy-nft` in a separate process which comms over GRPC
- KPNG Startup
    - Watchers (i.e. kube)
        - Start watchers: PICK ONE!!!
            - `kpng file` watches statefile ~ file2store
                - to-api, to-file, to-local ~ pick store consumer(s)
            - `kpng kube` watches apiserver ~ kube2store
                - to-api, to-file, to-local ~ pick store consumer(s)
             - api2store NOT THERE YET !!!
    - Consumers (i.e. to-file or to-local or... ) + (to-nft, to-ipvs, ...)
        - store2file.View(...) service - blocks
            - view(...)
                - every time something changes
                    - for node, add node to local file
                    - for service:
                    -   
                        - lookup all endpoints in proxystore
                            - add those endpoints
        - store2Diff.View(...)
            - view(...)
                - send Diff to the Since (STILL SAME MEMORY SPACE)
                - 
    - store2diff.View(...) services - blocks
        - wait(...)
        - view(...)
            - every time something changes
- User creates a service 'foo' (no endpoints yet)
    - kube2store.controller
        - Watch()<-- sees `v1.Service`
            - service-event-handler.go triggered
                - `func (h *serviceEventHandler) OnAdd(obj interface{}) {`
                - KPNG optimizes the v1.Service data
                    - so that it doesnt trigger unc events
                    - so that the storage size is min
                    - proxystore.go (Update)
                        - revision incremented (so clients can know if theyre uptodate)
                        - store2File() triggers the view(...) funciton above
                        - store2Local (relise on store2Diff) + store2LocalDiff() triggers the view(...) function
- User creates a pod

