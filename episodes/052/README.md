# Episode 52 : Instrumenting with Prometheus

- Hosted by @kris-nova
- Recording date: 20181005

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=zY2vpAf1SXA
" target="_blank"><img src="http://img.youtube.com/vi/zY2vpAf1SXA/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

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


## Reference Links

- [Instrumenting Go with Prometheus](https://prometheus.io/docs/guides/go-application)
-
