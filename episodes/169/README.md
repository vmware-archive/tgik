# Episode 169: Idemopotency, Leaky resources, and CSI

## News of the week

- Linkerd just released version 2.11, now support authorization polocy https://linkerd.io/2021/09/30/announcing-linkerd-2.11/
- KubeCon NA October 11-15 2021 https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/
- Production Identity Day https://events.linuxfoundation.org/production-identity-day-spiffe-spire-north-america/
- And many more Con!😧 https://events.linuxfoundation.org/
- Velero just released v1.7.0 https://twitter.com/projectvelero/status/1443941852265205765
- Falco 0.30.0 https://falco.org/blog/falco-0-30-0/
- https://github.com/sarun87/k8snetlook 
## Show notes

### CSI Idempotency

#### Idempotency requirements in CSI Spec
https://github.com/container-storage-interface/spec/blob/master/spec.md#createvolume

#### How CSI Drivers handle this
- GCEPD CSI Driver
- AWS EBS CSI Driver
- vSphere CSI Driver

### Avoid Leaking Resources
#### Volumes
- Fixed potential leak of volumes after CSI driver timeouts
  https://github.com/kubernetes-csi/external-provisioner/pull/312
- Always Honor PV Reclaim Policy
  https://github.com/kubernetes-csi/external-provisioner/issues/546
  https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2644-honor-pv-reclaim-policy
- Volume leak after CreateVolume started and before PV creation
  https://github.com/kubernetes-csi/external-provisioner/issues/486

#### Snapshots
- VolumeSnapshot design to avoid resource leaking
  https://github.com/kubernetes-csi/external-snapshotter/blob/v4.2.1/pkg/common-controller/snapshot_controller.go#L506
- Changes to always retry when create snapshot fails
  https://github.com/kubernetes-csi/external-snapshotter/blob/v4.2.1/pkg/sidecar-controller/snapshot_controller.go#L304

#### Interesting Links
- [Klusterd fix/break k8s clusters](https://www.youtube.com/playlist?list=PLz0t90fOInA5IyhoT96WhycPV8Km-WICj)
Next show  Nov 18, 2021, will be between AWS EKS team and Equinix folks https://www.youtube.com/watch?v=kjZ33qIzK6s
- [SDC2021](https://storagedeveloper.org/events/sdc-2021/schedule) was earlier this week, Sept 28-29
One CSI driver presented is [beegfs, included in the k8s csi driver list ](https://kubernetes-csi.github.io/docs/drivers.html) ([github](https://github.com/NetApp/beegfs-csi-driver))


#### Show contents

- 00:00:01 welcome to TGIK with xing ! episode 169
- 00:01:00 News of the week
- 00:10:00 CSI controller service RPC 
- 00:15:00 CreateVolume should be idempotent, so should most other things
- 00:17:00 Create volume names returns an ID
- 00:24:00 Volume UIDs and handles, AWS, Vsphere
- 00:27:00 DeleteVolume succeeds, even if a volume isnt there
- 00:36:00 AWS EBS csi driver idempotency support and parameter mismatches
- 00:51:00 Demo of deleting PVs, deletion timestamps, vSan leaked storage example

