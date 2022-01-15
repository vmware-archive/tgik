# Episode 181: KPNG Part Dos !

## News of the week

- dockershim is really OUT now : https://kubernetes.io/blog/2022/01/07/kubernetes-is-moving-on-from-dockershim/  
- sig-Windows-dev-tools demo: https://www.youtube.com/watch?v=3Z0NOrETjxY
- CSI proxy on windows: https://www.phillipsj.net/posts/how-to-use-the-windows-csi-proxy-and-csi-smb-driver-for-kubernetes/
- StatefulSets and DNS: https://jayunit100.blogspot.com/2021/12/looking-at-headless-services-coredns.html

## Show notes

- KPNG = [Kubernetes Proxy Next Gen](https://github.com/kubernetes-sigs/kpng)
- KPNG overall project update
    - code split up
    - ci jobs
    - bugs we found in our initial implementation 
- KPNG local-up-cluster, does it still work
- how to run `Kind` on IPv6 ?
- https://jayunit100.blogspot.com/search?q=triage
- https://miro.com/app/board/o9J_lxxSiFw=/
- https://github.com/kubernetes/kubernetes/pull/97081
- https://docs.google.com/presentation/d/1glOnUoiflU41xg34fVz9Mwl_4MyWl_5ubqhLXrVME9w/edit#slide=id.g10150b2b564_0_9
- 

```
1) docker build -t kpng:test -f Dockerfile .
2) hack/test_e2e.sh -b iptables -d -i ipv4
```

Come to antrea.io/live this wednseday to see Zac's and Amims Service-LB kube-proxy conformance suite !!!!!!
