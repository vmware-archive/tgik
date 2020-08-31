# Episode 061 : Falco with Kubernetes Dynamic Auditing

- Hosted by @krisnova
- Recording date: 2018-01-11

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=fjoxmTeh3Fk
" target="_blank"><img src="http://img.youtube.com/vi/fjoxmTeh3Fk/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review


## News from Around the K8s Community

- Two big CVEs:
    - [API Server](https://discuss.kubernetes.io/t/security-impact-of-kubernetes-api-server-external-ip-address-proxying/4072)
    - [Dashboard 1.10.1](https://discuss.kubernetes.io/t/security-release-of-dashboard-v1-10-1-cve-2018-18264/4069) 
        - [Detecting this CVE with Falco](https://sysdig.com/blog/privilege-escalation-kubernetes-dashboard/)
- [k8s 1.13.2 is out](https://github.com/kubernetes/kubernetes/releases/tag/v1.13.2)
    - [k8s 1.14 draft schedule](https://github.com/kubernetes/sig-release/pull/431) - not final, your mileage may vary, thar be dragons, etc. [README.md](https://github.com/kubernetes/sig-release/blob/254fcf6ade5d97035facaa1163f8ce8577b8c596/releases/release-1.14/README.md)
- [List of videos and slides from Kubecon](https://github.com/cloudyuga/kubecon18-NA) - Thanks Cloudyuga!
- [Another list of videos and slides from Kubecon](https://github.com/warmchang/KubeCon-North-America-2018) - Thanks William Zhang!
- Other software releases: [Weave Flux 1.9.0 is out](https://github.com/weaveworks/flux/releases/tag/1.9.0), [Heptio Ark .10.1](https://discuss.kubernetes.io/t/heptio-ark-v0-10-1-released/4166)
- [Crossplane](https://github.com/crossplaneio/crossplane)
- SIG apps wants to remind people about some [API deprecations](https://github.com/kubernetes/kubernetes/pull/70370) 
- [`systemd` CVE](https://www.openwall.com/lists/oss-security/2019/01/09/3)
    - Not all distros affected, check with your distro/vendor.
- Kubecon CFP closes in a week! Check out the info [here](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2019/cfp/) - Hope to see a bunch of you in Barcelona!
- Idea(s) for future TGIK sessions
    - Spinnaker
        - New [LA Meetup for Spinnaker](https://www.meetup.com/spinnaker-la/)

### Blog posts from around the web

- Anyone got anything interesting to add here? -jorge


## Show Notes

[Kate Kuchin on dynamic auditing](https://www.youtube.com/watch?v=7dJiagat4FI)


## Reference Links
 - [Falco Rules File Docs](https://github.com/falcosecurity/falco/wiki/Falco-Rules)
 - [Sysdig Filtering/Output Fields](https://github.com/draios/sysdig/wiki/Sysdig-User-Guide#all-supported-filters)
 - [Audit Example](https://github.com/falcosecurity/falco/tree/dev/examples/k8s_audit_config)
 - [Extra Rules](https://github.com/falcosecurity/falco-extras)
 - [Kernel Module](https://github.com/draios/sysdig/tree/dev/driver)
 - [DKMS](https://en.wikipedia.org/wiki/Dynamic_Kernel_Module_Support)
 - [Mark's Kubecon 2018 Seattle Talk on Falco K8s Audit](https://kccna18.sched.com/event/I1p9/deep-dive-falco-mark-stemm-sysdig)

## Docs

Thanks to friends at SysDig: Mark Stemm, Dan Papandrea (aka POP), Michael Ducy

```bash
# Create a 1.13 cluster with Dynamic Auditing with Kubicorn
kubicorn create tgik-falco-cluster -S kubicorn -p aws -M serverPool.bootstrapScripts[0]=kubicorn/amazon_k8s_ubuntu_16.04_master.sh -N serverPool.bootstrapScripts[0]=kubicorn/amazon_k8s_ubuntu_16.04_node.sh
kubicorn apply tgik-falco-cluster -S kubicorn

# Note: To later delete the cluster:
kubicorn delete tgik-falco-cluster -S kubicorn
```

Install falco

```bash
kubectl create -f k8s-with-rbac/falco-account.yaml
kubectl create configmap falco-config --from-file=k8s-with-rbac/falco-config
kubectl create -f k8s-with-rbac/falco-daemonset-configmap.yaml 

# Install falco service
kubectl apply -f falco-service.yaml
```

Create audit sink

```bash
kubectl apply -f audit-sink.yaml
```


Create role bindings with cluster-admin

```bash

# This is bad, don't do this
kubectl create clusterrolebinding cluster-admin --clusterrole=cluster-admin \
--user=user1 --user=user2 --group=group1
```
