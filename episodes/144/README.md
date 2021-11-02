# Episode 144 : Exploring The State of K8s on Windows

- Hosted by @jayunit100
- Recording date: 2021-01-28 

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=WnXnO97tsNo" target="_blank"><img src="https://i.ytimg.com/vi/WnXnO97tsNo/maxresdefault.jpg" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK and introductions/announcements
- 00:15:18 - Looking at how NetworkPolicy Truth tables work on unsupported clusters
- 00:32:29 - Viewing working NetworkPolicies in Windows with Calico 3.16
- 00:45:01 - Lookint at CAPI on Windows
- 00:55:40 - How CNIs are installed with post-kubeadm/preBootstrap commands on Windows
- 01:02:15 - How runtimeClasses, taints, nodeSelectors work together to schedule Windows pods
- 01:13:53 - How Cluster API works on NSX for windows with VSphere and NSX as the network plane
- 01:15:34 - How to use hub.docker.com to lookup windows images matching your OS
- 01:19:07 - Looking at CSI Proxy, briefly
- 01:26:45 - The most interesting problem in containerd and windows networking : The CNI ADD Codepath!

## News of the Week

* [CKA/CKAD Exam tips](https://itnext.io/kubernetes-journey-cka-ckad-exam-tips-ff73e4672833) from Brad McCoy
* From David McKay and Matt Turner: [Fun live coding of a KubeCtl plugin](https://www.youtube.com/watch?v=9FGA5vDeaOk) to add noun-verb ordering. (Joe agrees that noun verb is better!)
* OpenAI talks about [scaling k8s to 7500 nodes](https://openai.com/blog/scaling-kubernetes-to-7500-nodes/).
* David Giffin shares a [post on how to debug CrashLoopBackOff](https://releaseapp.io/blog/kubernetes-how-to-debug-crashloopbackoff-in-a-container).
* https://github.com/cockroachlabs/crl-scheduler
* https://learnk8s.io/troubleshooting-deployments 
  - images and wrong architectures
  - windows: "issue w/ kubelet" -> issue w/ HNS, or CNI provider
  - kube-proxy may not be syncing all the hns/iptables/ipvs ruls
  - firewall rules or network policies
* https://github.com/kubernetes/sig-release/tree/master/releases/release-1.21
* https://twitter.com/Aspenwilder/status/1354842524683100205?s=20

### VMware Tanzu related stuff
* Joe did an AMA for [kube.academy](https://kube.academy/) yesterday. [Video is now up](https://tanzu.vmware.com/content/webinars/jan-28-ask-me-anything-with-joe-beda-co-creator-of-kubernetes) (Warning: webinar registration link)
* Brian McClain wrote up [an article about "The Hate for YAML"](https://tanzu.vmware.com/developer/blog/the-hate-for-yaml-the-hammer-or-the-nail/) based on the "I'm sorry about the YAML" talk Joe gave at Software Circus
* Get an [early ebook copy of "Production Kubernetes"](https://tanzu.vmware.com/content/ebooks/production-kubernetes) from Josh Rosso, Rich Lander, Alex Brand and John Harris.  All great folks with a lot to share. (Warning: contact capture to download.)
* Tanzu Tuesdays! Tiffany and Ollie demonstrate the latest on [running Spring Boot apps on k8s](https://tanzu.vmware.com/developer/tv/tanzu-tuesdays/0039/).
* Paul presents a guide to [using Harbor to deal with Docker Hub rate limits](https://tanzu.vmware.com/developer/guides/kubernetes/harbor-as-docker-proxy/).

## Show Notes

## Windows K8s current news

- The "real" K8s Contributor experience on Windows
    - WSL2 : build, test, hack, build exes, copy them to c:/
    - `hnsdiag` and `hcsdiag` my 2 best friends
    - e2e tests, agnhost, ...
    - `e2e.exe` ... yes it works !
    - RDP!
        - getting felix / node logs
        - eventviewer and kube proxy / HNS logs
    - testing network policies

- Priveleged containers in windows https://github.com/kubernetes/enhancements/pull/2288#issuecomment-767300232 , for internal details, read about https://docs.microsoft.com/en-us/windows/win32/procthread/job-objects 

- Containerd on windows 
  - BLEEDING EDGE OMG
    - https://github.com/containerd/containerd/issues/4851
    - https://github.com/vmware-tanzu/antrea/issues/1581
    - https://github.com/projectcalico/calico/issues/4334
    - antrea : [same ipam, calls to HNSEndpoint](https://github.com/containernetworking/plugins/blob/master/plugins/ipam/host-local/main.go)
    - Kubelet -> ContainerD <-> CNI (race) ContainerD->HCSShim VNIC vs. docker bridge, see [docker design doc antrea](https://docs.google.com/document/d/1lSis0XnKz8UcJSkxTgRtDhP2DAwQtcZDjT6ZRySUb48/edit) original antrea design doc
- windows networkpolicies
    - Help us test them ! https://github.com/kubernetes/kubernetes/pull/98077 
    - Run the hacked e2e binary `wget https://storage.googleapis.com/jayunit100/content/e2e.test.win.12` followed by `./e2e.test.win.12 --provider=local --kubeconfig=/home/kubo/.kube/config --ginkgo.focus="Netpol.*" --ginkgo.skip="udp|Slow|SCTP" --node-os-distro="windows" --dump-logs-on-failure=false`
    - taints? you might want to add the `--non-blocking-taints="os,node-role.kubernetes.io/master,node.kubernetes.io/not-ready"`

# The ecosystem... 

- https://drive.google.com/file/d/1ZS4YGgPGXT5P-fpL2j0Ehdk4QhzVkZCE/view?usp=sharing - Kubelet, containerd, hcmshim, CNI interplay by Ravi.G

## CNI providers on Windows

- look at container networks `hnsdiag.exe list all`

- Calico, docker, containerd
    
    - calico: (quick demo)
        - https://aws.amazon.com/blogs/containers/open-source-calico-for-windows-containers-on-amazon-eks/
        - https://github.com/projectcalico/calico/blob/master/scripts/install-calico-windows.ps1
    - antrea:  (quick demo)
        - installer https://github.com/jayunit100/k8sprototypes/blob/master/windows/peri-min.yaml.sh#L290

- https://github.com/vmware-tanzu/antrea/blob/main/pkg/agent/util/net_windows.go#L304 

## Images

- powershell: `[System.Environment]::OSVersion.Version`
- images need to match your OS, look it up in this table! https://hub.docker.com/_/microsoft-windows


## Storage

- csi proxy: https://github.com/kubernetes-csi/csi-proxy/pull/106
- Interested in how to build a csi-proxy compatible storage solution? https://github.com/gab-satchi/vsphere-csi-driver#windows-prototype 

## CAPI and windows

- scaffolding for Windows MachineDeployment
https://github.com/jayunit100/k8sprototypes/blob/master/windows/peri-min.yaml.sh#L222
- Now: post kubeadm + image builder
  - Cloud init ignition
  - cloud-init
  - windows sysprep
- Node agent:
    
    - kubeadm post actions
      - 1) encapsulate common use cases into node agent
          - http proxies/certs arent easy to embed
      - 2) get rid of user_data size limits
    - reach out to CAPI / naadir / etc for future of images on CAPI


## CSI on Windows 
- https://github.com/gab-satchi/vsphere-csi-driver#windows-prototype
- 

## Add Windows Monitor to Prom Operator

- To create the prom daemonset https://raw.githubusercontent.com/jayunit100/k8sprototypes/master/windows/tgik-prometheus.yaml

- kubectl exec privileged powershell.exe get-hnsnetwork 

## CSI proxy
  - 

## image builder

## Containerd
- containerd / crio / ... -> lxc / runc / hcsshim  
- CRI : image + runtime 
- --> Calls CNI to setup networking
- runtime -> shim (`runc`, `runhcs`,`kata`, `firecracker`,`gVisor`,`plugin` )
  - shim binary : substitute w/ custom runhcs.exe
- Example:
    - run a windows iis pod  https://github.com/containerd/containerd/blob/master/pkg/cri/config/config_windows.go 

## Priveliged containers

kubectl exec privileged powershell.exe get-hnsnetwork will be working

# CNI

## Calico

## Antrea

### Create Folders
``` powershell
New-Item -ItemType Directory -Force -Path C:\k\antrea
New-Item -ItemType Directory -Force -Path C:\k\antrea\logs
New-Item -ItemType Directory -Force -Path C:\k\antrea\bin
New-Item -ItemType Directory -Force -Path C:\var\log\kube-proxy
```

### Set IP in kubeadm-flags.envs and set Node_name environment variable
``` powershell
$env:HostIP = (
  Get-NetIPConfiguration | Where-Object {
     $_.IPv4DefaultGateway -ne $null -and $_.NetAdapter.Status -ne "Disconnected"
  }
).IPv4Address.IPAddress
$file = 'C:\var\lib\kubelet\kubeadm-flags.env'
$newstr="--node-ip=" + $env:HostIP
$raw = Get-Content -Path $file -TotalCount 1
$raw = $raw -replace ".$"
$new = "$($raw) $($newstr)`""
Set-Content $file $new


[Environment]::SetEnvironmentVariable("NODE_NAME", (hostname).ToLower())
```

### Download Antrea Components
curl.exe -LO "https://raw.githubusercontent.com/vmware-tanzu/antrea/master/hack/windows/Install-OVS.ps1"
### Install Open Vswitch

We need to enable test signing if this is the first time...
``` powershell
invoke-expression "bcdedit /set TESTSIGNING ON" 
Restart-computer

./Install-OVS.ps1 
```

Important that installing on the host requires a kubeconfig to get the secrets and nodes  for the demo we have given the rights to system:nodes C:/etc/kubernetes/kubelet.conf << this needs to be changed in future!

### Setup services 

*This cannot be done over ssh (RDP or terminal in)*
    ``` powershell
    cd c:\k\antrea

    curl.exe -LO https://raw.githubusercontent.com/vmware-tanzu/antrea/master/hack/windows/Helper.psm1
    curl.exe -LO http://w3-dbc302.eng.vmware.com/rcao/image/containerd/antrea-agent.exe

    mv antrea-agent.exe c:\k\antrea\bin
    Import-Module ./helper.psm1

    Install-AntreaAgent -KubernetesVersion "v1.19.1" -KubernetesHome "c:/k" -KubeConfig "C:/etc/kubernetes/kubelet.conf" -AntreaVersion "v0.12.0" -AntreaHome "c:/k/antrea"
    New-KubeProxyServiceInterface

    nssm install kube-proxy "c:/k/kube-proxy.exe" "--proxy-mode=userspace --kubeconfig=C:/etc/kubernetes/kubelet.conf --log-dir=c:/var/log/kube-proxy --logtostderr=false --alsologtostderr"
    nssm install antrea-agent "c:/k/antrea/bin/antrea-agent.exe" "--config=c:/k/antrea/etc/antrea-agent.conf --logtostderr=false --log_dir=c:/k/antrea/logs --alsologtostderr --log_file_max_size=100 --log_file_max_num=4"

    nssm set antrea-agent DependOnService kube-proxy ovs-vswitchd
    nssm set antrea-agent Start SERVICE_DELAYED_START

    start-service kube-proxy
    start-service antrea-agent            
    ```

# Windows commands !

## Do it all in powershell
- Tail a file: `Get-Content myfile.txt -Wait`
- Find a file: `Get-Childitem –Path C:\ -Recurse -Name *ctr*`
- List processes:`Get-Process *ovs* | Format-Table -Property Name`
- Chocolatey: `Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))`, then run `refreshenv` in the shell 
- use Vim: `choco install vim -y`
- You can run `ctr --namespace=k8s.io containers list`
- Look at VNIC Events `Get-WinEvent Microsoft-Windows-Hyper-V-VmSwitch-Operational` and `Microsoft-Windows-Hyper-V-Compute-Operational`
- sort cmd line output:`| Sort-Object`
- look at the containers `hcsdiag.exe list all`
- look at container networks `hnsdiag.exe list all`
- look at all containers `ctr.exe --namespace=k8s.io c ls`
- version `[System.Environment]::OSVersion.Version`
- grep ~ `Select-String`, i.e. `Get-WinEvent -ListLog * | Out-String -Stream | Select-String SSH`

# Log File locations

Kube-proxy: c:/var/log/kube-proxy
Antrea-Agent: c:/k/antrea/logs
Kubelet: c:/var/log/kubelet
Containers: c:/var/log/containers
OVS: c:/openvswitch/var/log/openvswitch

.

# REFERENCES!

https://github.com/kubernetes/kubernetes/issues/98102
https://github.com/vmware-tanzu/antrea/issues/1581
https://github.com/containerd/containerd/issues/4851
https://www.youtube.com/watch?v=FKoVztEQHss

# OTHER STUFF
- cluster api internal stuff
    - 
- What is Windows sig currently focused on?
    - Having a proper end-to-end post merge release blocking test suite using containerd. Ravi.G, James et al. are working on it
        - https://github.com/kubernetes/kubernetes/issues/93276
    - 1.21 KEP
        - KEP 2258: Use kubectl to view system service logs #2271
        - KEP 1981: Windows privileged container KEP updates for alpha #2288

## Reference Links
- https://drive.google.com/file/d/1ZS4YGgPGXT5P-fpL2j0Ehdk4QhzVkZCE/view?usp=sharing - Kubelet, containerd, hcmshim, CNI interplay by Ravi.G
