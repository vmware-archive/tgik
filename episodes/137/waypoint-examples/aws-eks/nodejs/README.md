# Waypoint AWS-EKS Example

|Title|Description|
|---|---|
|Pack|Cloud Native Buildpack|
|Cloud|AWS|
|Language|NodeJS|
|Docs|[Kubernetes](https://www.waypointproject.io/plugins/kubernetes)|
|Tutorial|[HashiCorp Learn](https://learn.hashicorp.com/tutorials/waypoint/get-started-kubernetes)|

Requirements:

- You must create an ECR registry named `waypoint-example` (or choose your own name and edit `waypoint.hcl`).
- Create an EKS cluster as shown at [HashiCorp Learn](https://learn.hashicorp.com/tutorials/terraform/eks?in=terraform/kubernetes).
- Configure `kubectl` to communicate to the Kubernetes cluster (also covered in the tutorial above).
