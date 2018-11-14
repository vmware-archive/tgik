# Episode 53 : Cluster API for AWS

- Hosted by @kris-nova
- Recording date: 20181012
- Live notes: https://hackmd.io/yrLz1UgVSF2xqvELw5eZaQ

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=e1XCsuTYUa4
" target="_blank"><img src="http://img.youtube.com/vi/e1XCsuTYUa4/hqdefault.jpg" width="480" height="360" border="10" /></a>

- 00:00:00 - Welcome to TGIK!
- 00:05:10 - Week in Review
  - 00:15:00 - Spark on Kubernetes
- 00:17:30 - Function Types in Go
- 00:23:30 - The Cluster API
- 00:34:30 - Cluster API Provider AWS

## Show Notes

 - [Cluster API](https://github.com/kubernetes-sigs/cluster-api)
 - [Cluster API API](https://github.com/kubernetes-sigs/cluster-api/tree/master/pkg/apis/cluster/v1alpha1)
 - [AWS Cluster API](https://github.com/kubernetes-sigs/cluster-api-provider-aws)
 - [Getting Started on AWS](https://github.com/kubernetes-sigs/cluster-api-provider-aws/blob/master/docs/getting-started.md)
 - [Kubectl Sudo](https://github.com/postfinance/kubectl-sudo)
 - [Baremetal Install of Kubernetes](https://www.reddit.com/r/kubernetes/comments/9nbehq/baremetal_install_of_kubernetes/)
 - [Function Types in Go](https://goplay.space/#qZ_xY-LRc_u)
 - [Deploying Spark on Kubernetes](https://testdriven.io/deploying-spark-on-kubernetes#.W732ZWgHcB4.reddit)

ClusterCTL command

```
clusterctl create cluster -v4 --provider aws  -m ./clusterctl/examples/aws/out/machines.yaml   -c ./clusterctl/examples/aws/out/cluster.yaml   -p ./clusterctl/examples/aws/out/provider-components.yaml --existing-bootstrap-cluster-kubeconfig ~/.bootstrap-kubeconfig --existing-bootstrap-cluster-kubeconfig=/Users/nova/.kube/config

```

## Reference Links

 - [Go 2 Draft Designs](https://go.googlesource.com/proposal/+/master/design/go2draft.md)
 - [Provisioning of a Minikube bootstrap cluster for Cluster API](https://github.com/kubernetes-sigs/cluster-api/pull/195)
 - [How to Build a Kubernetes Cluster with ARM Raspberry Pi then run .NET Core on OpenFaas - Scott Hanselman](https://www.hanselman.com/blog/HowToBuildAKubernetesClusterWithARMRaspberryPiThenRunNETCoreOnOpenFaas.aspx)
 - [AWS Vault](https://github.com/99designs/aws-vault)