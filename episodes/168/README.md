# TGIK Episode 168: Tasting the rainbow of GKE, Minishift, Tanzu, win-dev-tools, local-up-cluster.sh, and Minishift

https://unsplash.com/photos/QuoNoM9nyz0

## News

- Whats this all about https://github.com/kubernetes/kubernetes/pull/93898 
- Why you cant  ping a service (not new, but news to some folks, worth reading) https://nigelpoulton.com/why-you-cant-ping-a-kubernetes-service/ 
- https://github.com/kubernetes/website/pull/29783/files  how to build your own KPNG backend by lars
- https://stackoverflow.com/questions/69318744/does-go-mod-resolve-transitive-dependencies-for-an-dep-by-default-and-if-so-wh 
- 

## TASTE THE RAINBOW !

- Azure
  - whats a service principal?
  - whats the RBAC option to?
  - What calico version to i get and hows it configured?
      - 10->12 core quota
  - why do i need an azure file share to make a azure shell ?
  - oh, you can kubectl debug a node 
      - ??? still didnt see any calico infrastructure ???
- GKE
    - konnectivity server: Interesting way for K8s controlplane to talk to the nodes?
    - CNI: Cillium or GKE native
- VMWare Tanzu
  - Antrea
  - Calico
  - ClusterAPI
  	- MachineDeployments
	- VSphereClusters
	- Clusters
	- KubeadmConfig
		- cloud init userdata, base 64 encoded
  - Also supports:
	  - EC2
	  - Azure
  - For UI, try Tanzu Mission Control or TGKS
- Openshift
    - To try OpenShift faster way use the [OpenShift Playground](https://learn.openshift.com/playgrounds/openshift47) 
    - Local OpenShift is [CodeReadyContainers (CRC)](https://access.redhat.com/documentation/en-us/red_hat_codeready_containers/1.32/html/getting_started_guide/installation_gsg#minimum-system-requirements_gsg) You need free RedHat account to download [crc](https://console.redhat.com/openshift/create/local)
    - Without Pull Secrets [OKD CRC](https://www.okd.io/crc.html) - To try OpenShift faster way use the [OpenShift Playground](https://learn.openshift.com/playgrounds/openshift47) 
- EKS Anywhere 
    -  [Local Cluster](https://anywhere.eks.amazonaws.com/docs/getting-started/local-environment/)
- k3s... fun to use on a rasberry

## Video

- 3:00:00 News of the week
- 5:15:00 How does the kubelet figure out its IP when it changes?
- 7:00:00 lets look at a one line PR with 50 comments
- 15:00:00 reboot tests
- 18:00:00 kpng blog post by lars
- 20:00:00 backends in kpng
- 21:00:00 go mod - how can you import only one package?
- 24:00:00 AKS clusters and the Azure terminal
- 28:00:00 Kubenet, azure cni + add on network policy
- 32:00:00 Why no cillium on Azure
- 33:00:00 looking at whats on an azure cluster ootb
- 37:00:00 kubectl debug thanks moz!
- 39:00:00 moving on to tanzu
- 44:00:00 management and workload clusters on Tanzu 1.4 + Vsphere 
- 45:00:00 ytt to describe tanzu clusters
- 50:00:00 how kubeadm and user_data cloud init works
- 59:00:00 minishift vs CodeReadyContainers (okd)
- 1:10:00 service Accounts, GCE, and making openshift in the cloud
- 1:18:00 hack/local-up-cluster.sh the most powerful dev tool ever
- 1:24:00 breaking the kubelet and watching it fail
- 1:41:00 GKE whats the konnectivity stuff about ? looking at the KEP
- 1:44:00 checking out sig-windows-dev-tools to make a cluster from source w windows nodes
- 1:48:00 carlos impromptu shows us some ibm/openshift stuff
- 1:51:00 the windows dev tools come up...
- 1:55:00 ricardos on prem openshift environment comes up

