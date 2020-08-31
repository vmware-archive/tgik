# Episode 52 : Instrumenting with Prometheus

- Hosted by @krisnova
- Recording date: 2018-10-05

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=zY2vpAf1SXA
" target="_blank"><img src="http://img.youtube.com/vi/zY2vpAf1SXA/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## K8s news

- [Kubernetes 1.12 is out](https://kubernetes.io/blog/2018/09/27/kubernetes-1.12-kubelet-tls-bootstrap-and-azure-virtual-machine-scale-sets-vmss-move-to-general-availability/). 1.12.1 came out just now!
- [Steering Committee Election Results](https://groups.google.com/forum/#!topic/kubernetes-dev/EL4W0gQGxMA) - congrats to Aaron Crickenberger, Tim St. Clair, and Davanum Srinivas 
- The folks from automaticserver have created a [CRI shim for LXC/LXD](https://discuss.kubernetes.io/t/lxe-released-a-kubernetes-integration-of-lxc-lxd/3022/1)
- [Gardener .12 released](https://discuss.kubernetes.io/t/gardener-0-12-0-released/2935)
- [Schedule for Kubecon is now available](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2018/schedule/)
- [Teleport 3.0](https://gravitational.com/blog/teleport-release-3/) is out
- [Digital Ocean's k8s now in limited availability](https://blog.digitalocean.com/announcing-limited-availability-of-digitalocean-kubernetes/)

## Show Notes

## Set up cluster with Prometheus

Create a cluster in AWS with Kubicorn

```bash
kubicorn create prom-cluster -paws
kubicorn apply prom-cluster
```

Install Prometheus and other apps

```bash
kubectl apply -f bundle.yaml
kubectl apply -f prometheus.yaml
```

Poke a hole in the firewall

 - Open the Node security group
 - Expose the TCP Ports to each LoadBalancer created
    - 9093 -> alertmanager
    - 3000 -> grafana
    - 9090 -> prometheus

Change the Makefile to fit your docker registry if needed, and build the example app

```bash
make container push
```

Deploy the example app

```bash
kubectl run tgik-app -n monitoring --image krisnova/tgik-instrumentation-app
kubectl expose -n monitoring tgik-app --port 9093 --target-port 1313
```


To reload the application

```bash
kubectl scale tgik-app -n monitoring --replicas=0
kubectl scale tgik-app -n monitoring --replicas=1
```

To reload Prometheus without taking it down:

```bash
curl -i -X POST localhost:9090/-/reload
```

**Do not expose these lifecycle endpoints to the public**

## Reference Links

- [Instrumenting Go with Prometheus](https://prometheus.io/docs/guides/go-application)
- [Joe's Prometheus v1](https://www.youtube.com/watch?v=pDb2psNcvKU)
- [All the Autoscaling](https://caylent.com/kubernetes-autoscaling/)
- [Step by Step CRDs](https://medium.com/@dimitris.kapanidis/tailor-kubernetes-to-your-needs-with-crds-7a1e322116a1)
- [Cloud Native London Videos](https://skillsmatter.com/conferences/10160-cloudnative-london-2018#program)
- [Prometheus Query Language](https://prometheus.io/docs/prometheus/latest/querying/basics/)
- [Original Source of Prometheus](https://github.com/giantswarm/kubernetes-prometheus)



https://github.com/prometheus/client_golang/blob/master/prometheus/go_collector.go#L272
