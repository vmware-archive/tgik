# Episode 131 : No mo' yolo with Validating Admission Controllers
- Hosted by @ciberkleid and @pczarkowski
- 09/11/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=RVDK0m2XQeg
" target="_blank"><img src="http://img.youtube.com/vi/RVDK0m2XQeg/hqdefault.jpg" width="480" height="360" border="10" /></a>



## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### Core K8s

- 1.19.1 is out, here's the [changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md/#changelog-since-v1190)
- Point releases [v1.19.2/1.18.9/v1.17.12 are scheduled for September 16](https://groups.google.com/g/kubernetes-dev/c/btWkJGkLkhA)
- Proposal for [WG Reliability](https://groups.google.com/g/kubernetes-dev/c/wx2qjnPEzc4)
- Check out the blog post on [warnings](https://kubernetes.io/blog/2020/09/03/warnings/)
- Steering Committee [Election Update](https://groups.google.com/g/kubernetes-dev/c/sskISyF_5xk)
    - Nomination/exception processes are now closed. 
- KubeCon videos are now [all on YouTube](https://www.youtube.com/playlist?list=PLj6h78yzYM2O1wlsM-Ma-RYhfT5LKq0XC) if you wanna catch up. 


### K8s and Cloud Native Ecosystem

- The Jetstack folks outline [performing a live CNI migration](https://blog.jetstack.io/blog/cni-migration/)
- Jakub Pavlik looks at [Kubernetes Storage performance comparison](https://medium.com/volterra-io/kubernetes-storage-performance-comparison-v2-2020-updated-1c0b69f0dcf4)
- CrunchyData has released a [postgres operator](https://github.com/CrunchyData/postgres-operator)
- Bookmark this - [https://k8syaml.com/](https://k8syaml.com/)


## Show Notes

Reference docs:

- https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers

Examples:

1. Validating Admissions Controller from Eric Shanks' ***The IT Hollow*** blog
    - [Write-up](https://theithollow.com/2020/05/26/kubernetes-validating-admission-controllers) and [code repo](https://github.com/theITHollow/warden)
    - This controller is written in Python and provides a clear and simple example of the minimum controller & webhook requirements
    - Slightly modified [fork](https://github.com/paulczar/ithollow-warden/tree/tgik)  (uses a cert manager & includes a [kind](https://kind.sigs.k8s.io/) cluster config file)

2. Mutating Admissions Controller  by Paul
    - [Code repo](https://github.com/paulczar/m13k)
    - This controller is written in Golang and uses [ytt](https://github.com/k14s/ytt) to manipulate the incoming request; the controller is generic and reads the specific ytt manipulations from a ConfigMap
    - Includes steps to run locally, deploy with `kubectl`, or deploy using `helm`



