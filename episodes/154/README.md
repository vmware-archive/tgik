# Episode 154 : Pixie
- Hosted by @e_k_anderson
- 05/14/2021

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=CYQsLfVcgYI
" target="_blank"><img src="http://img.youtube.com/vi/CYQsLfVcgYI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:02:40 - Week in Review
- 00:14:10 - Getting started with Pixie
- 00:16:00 - Architecture overview
- 00:18:00 - No instrumentation vs instrumentation
- 00:26:30 - Looking at the Pixie install
- 00:34:00 - Starting `px-sock-shop` demo
- 00:38:50 - Log in to Pixie Cloud
- 00:45:00 - OOM the 8GB Minikube
- 00:52:00 - Get a bigger Minikube
- 00:57:20 - Installing Pixie and demo again
- 01:01:00 - Up and running with `px-online-boutique` demo
- 01:04:20 - Switch to "data isolation mode" and explore passthrough vs direct
- 01:11:00 - Exploring the Pixie dashboards
- 01:13:50 - Exploring the "scratch pad"
- 01:17:00 - 3 minute audio glitch... oops!
- 01:23:20 - "PxL" language, a DSL like Pandas
- 01:40:00 - Evan gives up on learning PxL on the fly
- 01:40:00 - Discussing dynamic tracing/logging with eBPF
- 01:45:00 - Comparision of Prometheus & Pixie models and usage (speculation)
- 01:47:00 - Tracing TLS methods
- 01:52:30 - Looking for apiserver traces
- 01:57:20 - Closeout

## Week in Review

### Core K8s

* [Release 1.21 webinar](https://community.cncf.io/events/details/cncf-cncf-online-programs-presents-cncf-live-webinar-kubernetes-121-release/). Hear about the new features in 1.21 from the Release and Enhancement leads who shepherded the process!
* [Swap KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2400-node-swap) -- making swap not destroy your cluster, but enable it for certain workloads

### K8s and Cloud Native Ecosystem

* [Kubecon](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) last week! What were your favorite talks?
   * There were 17 day-zero and day-minus-one events!
   * [The Long, Winding and Bumpy Road to CronJob’s GA - Maciej Szulik, Red Hat & Alay Patel, Red Hat](https://youtu.be/o5h6s3A9bXY)
   * [Hacking into Kubernetes Security for Beginners - Ellen Körbes, Tilt & Tabitha Sable, Datadog](https://youtu.be/mLsCm9GVIQg)
   * [Kubecon video playlist](https://www.youtube.com/playlist?list=PLj6h78yzYM2MqBm19mRz9SYLsw4kfQBrC)
       * [Automating Your Home with K3s and Home Assistant - Eddie Zaneski & Jeff Billimek](https://www.youtube.com/watch?v=icyTnoonRqI&list=PLj6h78yzYM2MqBm19mRz9SYLsw4kfQBrC&index=97)
* New project: [kcp: minimal Kubernetes API server](https://github.com/kcp-dev/kcp)
* [Antrea CNI](https://github.com/antrea-io/antrea) accepted as a CNCF sandbox project
* [ChaosBlade chaos engineering toolkit](https://github.com/chaosblade-io/chaosblade) accepted as a CNCF sandbox project
* [Fluid dataset orchestrator](https://github.com/fluid-cloudnative/fluid) accepted as a CNCF sandbox project
* [Submariner cross-cluster connector](https://github.com/submariner-io/submariner/) accepted as a CNCF sandbox project
* [Vineyard zero-copy data sharing](https://github.com/alibaba/v6d/) accepted as a CNCF sandbox project
* [WasmEdge WASM virtual machine](https://github.com/WasmEdge/WasmEdge) accepted as a CNCF sandbox project
* Blog Post: [LyftLearn: ML Model Training Infrastructure built on Kubernetes](https://eng.lyft.com/lyftlearn-ml-model-training-infrastructure-built-on-kubernetes-aef8218842bb)
* OSS from Google: [Modernizing Oracle operations with Kubernetes and El Carro](https://opensource.googleblog.com/2021/05/modernizing-oracle-operations-with-kubernetes-el-carro.html). Opeerator for Oracle.
* Falco response engine, four approaches: [Kubeless](https://falco.org/blog/falcosidekick-reponse-engine-part-1-kubeless/), [OpenFaaS](https://falco.org/blog/falcosidekick-reponse-engine-part-2-openfaas/), [Knative](https://falco.org/blog/falcosidekick-reponse-engine-part-3-knative/), [Tekton](https://falco.org/blog/falcosidekick-response-engine-part-4-tekton/)

## Show Notes

* Looking at https://pixielabs.ai and https://px.dev
* Two versions:
   * New Relic hosted: https://pixielabs.ai
   * Non-vendor / OSS: https://px.dev
* [Pixie on Kubernetes Podcast ](https://kubernetespodcast.com/episode/150-pixie/)