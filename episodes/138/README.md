# Episode 138 : Harbor
- Hosted by @tiffanyfayj and @pczarkowski
- 10/30/2020

<!--- Thumbnailed embed of the video, gNMZOlEKJnI is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=gNMZOlEKJnI
" target="_blank"><img src="http://img.youtube.com/vi/gNMZOlEKJnI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:03:40 - Week in Review
- 00:21:00 - Installing Harbor
- 00:38:51 - Harbor Achitecture
- 00:44:29 - Vulnerability Scanning
- 00:48:30 - Creating a new Project / Repository
- 00:50:23 - Create robot account & kubernetes pull secret
- 01:07:20 - Uploading and Sharing Helm Charts
- 01:23:43 - Proxying through to dockerhub
- 01:35:55 - Replicating from another Harbor
- 01:50:00 - Wrapup

## Week in Review

### Core K8s

- [Reminder: Kubernetes is considering moving to 3 releases per year! Add your input.](https://github.com/kubernetes/sig-release/issues/1290)
- [SCTP is GA (feature gate now on)](https://github.com/kubernetes/kubernetes/pull/95566)
    - [Stream Control Transmission Protocol](https://en.wikipedia.org/wiki/Stream_Control_Transmission_Protocol)

### K8s and Cloud Native Ecosystem

- [PVC AutoResizer: Expanding PVCs automatically](https://blog.kintone.io/entry/pvc-autoresizer)
- [Exploring Waypoint with Saiyam](https://www.civo.com/learn/waypoint-solving-the-build-deploy-and-release-problem)
- [Update on Helm Chart Repo Deprecation](https://helm.sh/blog/charts-repo-deprecation/)
    - Nov 6 - stable / incubator removed from artifacthub
    - Nov 13 - stable / incubator become read only
    - _after_ Nov 13 - Downloads of charts will be redirected to a readonly archive on GH pages
        - See [Switching to use read only archive](https://helm.sh/docs/faq/#i-am-getting-a-warning-about-unable-to-get-an-update-from-the-stable-chart-repository)
- [Open Telemetry Release Cadence announced (path to GA!)](https://medium.com/opentelemetry/tracing-specification-release-candidate-ga-p-eec434d220f2)
- [Cluster API EKS Support!](https://www.weave.works/blog/introducing-eks-support-in-cluster-api)

## Show Notes

[install scripts](./install-scripts)
[Harbor Website](https://goharbor.io/)
[Harbor GitHub](https://github.com/goharbor/harbor)

- https://github.com/goharbor/harbor/wiki/Architecture-Overview-of-Harbor


- Pushing an image to to Harbor and running it on Kubernetes
    ```
    docker tag <image-name> <harbor-fqdn>/<project-name>/<image-name>
    docker push <harbor-fqdn>/<project-name>/<image-name>

    kubectl run <image-name> --image=<harbor-fqdn>/<project-name>/<image-name>
    ```

- Pulling down a helm chart and pushing it to Harbor
    ```
    helm fetch bitnami/nginx

    # Go to Harbor -> Projects -> <project-name> -> Upload tar file

    helm repo add <repo-name> https://<harbor-fqdn>/chartrepo/<project-name>
    helm repo update
    helm install nginx <repo-name>/nginx
    ```

- Create a robot user and set it up as a kubernetes secret for a pod
    - [here](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#create-a-secret-by-providing-credentials-on-the-command-line)

- Pushing to Harbor with Helm
    Install plugin: https://github.com/chartmuseum/helm-push

- eBPF Summit YouTube links
  * [Day 1](https://www.youtube.com/watch?v=1GSgyvn4N7E&amp;feature=youtu.be)
  * [Day 2](https://www.youtube.com/watch?v=jw8tEPP6jwQ&amp;feature=youtu.be)

- Software Circus Fest YouTube Streams
  * [Schedule](https://www.softwarecircus.io/software-circus-nightmares-on-cloud-street-schedule)
  * YouTube Streams are no longer available, watch https://www.softwarecircus.io/ (presumeably they'll post recordings there)

- DockerHub changes this weekend:
  * [Pull rate-limiting starts Nov 2](https://www.docker.com/blog/what-you-need-to-know-about-upcoming-docker-hub-rate-limiting/)
  * [Checking Your Current Docker Pull Rate Limits and Status](https://www.docker.com/blog/checking-your-current-docker-pull-rate-limits-and-status/)
  * [Image retention changes delayed to mid-2021](https://www.docker.com/blog/docker-hub-image-retention-policy-delayed-and-subscription-updates/)
  * [Registry Creds Operator - one solution for Docker Hub auth](https://inlets.dev/blog/2020/10/29/preparing-docker-hub-rate-limits.html)

- Mentioned in chat:
  * Josh Rosso mentioned http://lwkd.info/ for k8s core development news.