# Episode 055 : Creating container images with Kaniko

- Hosted by @kris-nova
- Recording date

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=rt0fqOekgjI
" target="_blank"><img src="http://img.youtube.com/vi/rt0fqOekgjI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Show Notes

#### Creating a build context

```bash
tar -C ./context -zcvf context.tar.gz .
gsutil cp context.tar.gz gs://krisnova-tgik
```

## Reference Links

 -
 - [Pod Lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)
 - [Link to running Kaniko in Kubernetes](https://github.com/GoogleContainerTools/kaniko#running-kaniko-in-a-kubernetes-cluster)
 - [Kaniko vs Others](https://github.com/GoogleContainerTools/kaniko/blob/master/README.md#comparison-with-other-tools)

