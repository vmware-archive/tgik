# Episode 156 : "TGI Kubernetes 156: Exploring upstream k8s builds w/ multi-arch"

- Hosted by @jayunit100
- Recording date: 2021-05-28

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url 

<a href=https://www.youtube.com/watch?v=l3TWbrWkVzY"></aref>
--> 

This week we're going to dig into image builder, imgpkg, agnhost, and other image'y stuff in the K8s universe.

Show notes: https://hackmd.io/oC8hn8dJTLGFiU5Jz51b3w

## Table of Contents

00:03:00 - audio check

00:05:00 - Week in review

00:10:00 - Intro to multi OS containers, pause image, agnhost on kubernetes

00:21:00 - Dockerfile for windows

00:30:00 - peri's blog post on buildx

00:38:00 - speaking of manifest

00:35:00 - deep dive into OCI and umoci

00:45:00 - intro to OVAs on windows

01:00:00 - nssm, systemctl, windows, linux


## Week in Review

- https://stackoverflow.blog/2021/05/19/rethinking-system-architecture-can-kubernetes-help-to-solve-rewrite-anxiety/
- https://github.com/k3s-io/kine 
- Episode 154
    - Episode 154, we found the bug ! it was nftables
    - https://github.com/kubernetes-sigs/kpng/pull/29/files
- Azure AKS added Containerd support for Windows
    - https://azure.microsoft.com/en-us/updates/public-preview-aks-support-for-containerd-for-windows-server-containers/
- https://increment.com/containers/


## Show Notes

### K/K 
- Playin w AGNHost images
    - Why cant you run scratch image on windows ????
    - https://perithompson.netlify.app/blog/creating-multiarch-containers/
        https://github.com/kubernetes/kubernetes/blob/master/build/pause/Makefile#L131 
    - docker run -t -i k8s.gcr.io/e2e-test-images/agnhost:2.21
    - buildx allows you to ADD -> simple commands into a  
        - https://github.com/kubernetes/kubernetes/tree/master/build/pause
        - https://github.com/kubernetes/kubernetes/blob/master/build/pause/Makefile

### Kubelets, Hosts, images

- Image-Builder (OVA)
    - https://jayunit100.blogspot.com/2021/04/building-windows-ovas-on-os-x.html?q=image-builder 
    - `make build-node-ova-local-windows-2019`
    - https://spin.atomicobject.com/2013/06/03/ovf-virtual-machine/
    - "your always calling things burritos"
- Image-Builder (VHD)
    - customizing
    - exploring the VM
    - goss  
        - https://github.com/aelsabbahy/goss
### OCI
- https://github.com/opencontainers/image-spec/blob/master/spec.md
- Skopeo : build + fetch
    - ubuntu, build in docker, and apt install `libgpgme11-dev`
    - whats libdevmapper ? 
    - ... apt installed it
- Using skopeo and umoci 
    ```
        skopeo copy docker://opensuse/amd64:42.2 oci:opensuse:latest
        umoci ls --layout opensuse
        umoci ls --layout opensuse
    ```

### IMAGE STUFF

### Demo
