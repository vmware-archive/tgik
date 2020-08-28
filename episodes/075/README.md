# Episode 75 : Troubleshooting Container Networking

- Hosted by @mauilion
- Recording date: 2019-05-03

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://youtu.be/IhbJ3ll4usI" target="_blank"><img src="http://img.youtube.com/vi/IhbJ3ll4usI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

- [Effective Secrets with Vault and Kubernetes](https://itnext.io/effective-secrets-with-vault-and-kubernetes-9af5f5c04d06)
- [Grafana Dashboards for k8s administrators](https://povilasv.me/grafana-dashboards-for-kubernetes-administrators/)
- Release Stuff
    - k8s [1.15 alpha 2 is out](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.15.md#v1150-alpha2)
    - k8s [1.11.10 is out](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md)
    - [Deadlines are horrible](https://groups.google.com/forum/#!msg/kubernetes-sig-release/dGVBrlkOXQo/5m1zFTT7AwAJ)
- [Wireguard for Kubernetes](https://gravitational.com/blog/announcing_wormhole/)
- [Prior art for wiregaurd on kubernetes](https://github.com/squat/kilo)
- [Why script based deployments don't scale](https://blog.armory.io/why-script-based-deployments-to-kubernetes-dont-scale/)
- [How you can help localize k8s docs](https://kubernetes.io/blog/2019/04/26/how-you-can-help-localize-kubernetes-docs/)
- [Docker Hub Maintenance](https://success.docker.com/article/docker-hub-maintenance) - (DevOps tip 'o the day: don't depend on external services in production!)


## Show Notes

Kris has covered different CNIs in their own episodes, check 'em out:
- [Weave](https://github.com/heptio/tgik/blob/master/episodes/050/README.md)
- [Flannel](https://github.com/heptio/tgik/blob/master/episodes/049/README.md)
- [Cilium](https://github.com/heptio/tgik/blob/master/episodes/047/README.md)
- [Calico](https://github.com/heptio/tgik/blob/master/episodes/045/README.md)


## Reference Links
repo with sources: github.com/mauilion/kind-cni-playground
[calico-docs](https://docs.projectcalico.org/v3.7/introduction/)
[flannel docs](https://github.com/coreos/flannel)
[cilium docs](https://docs.cilium.io/en/v1.5/)
[canal docs](https://docs.projectcalico.org/v3.7/getting-started/kubernetes/installation/flannel#installing-with-the-kubernetes-api-datastore-recommended)
https://medium.com/@anilkreddyr/kubernetes-with-flannel-understanding-the-networking-part-2-78b53e5364c7
https://github.com/nicolaka/netshoot
https://github.com/cilium/cilium/pull/7839
