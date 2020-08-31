# Episode 054 : Building containers and images for Kubernetes with Buildah

- Hosted by @krisnova
- Recording date: 2018-10-19

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=Z9t1DvNcseA
" target="_blank"><img src="http://img.youtube.com/vi/Z9t1DvNcseA/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:02:00 - Hackmd and tgik repo description
- 00:04:20 - Build without a docker daemon
- 00:06:00 - multicluster-controller operator pattern for many clusters neat!
- 00:08:22 - Buildah repo
- 00:09:00 - kubebox
- 00:11:30 - dig in to build tags in go
- 00:18:00 - A bit about docker in FreeBSD that Kris worked on.
- 00:22:20 - Buildah Intro
- 00:32:50 - install time and examples
- 00:35:00 - episode code located [here](https://github.com/heptio/tgik/tree/master/episodes/054)
- 00:37:00 - What does buildah do.
- 00:41:00 - What is an [oci spec](https://github.com/opencontainers/image-spec)
- 00:45:00 - Let's build an image in Buildah.
- 01:06:00 - push to gcr.
- 01:10:00 - run this image in Docker on the laptop.
- 01:12:00 - let's run this container in k8s.
- 01:13:00 - What else can we do (explore tags and shas and k8s runs)
- 01:19:00 - what even is an unshare :smile: 
- 01:23:00 - Wrap up :smile: 



## Show Notes

Question about comparing Build with [Jesses' genuinetools/img tool](https://github.com/genuinetools/img)
Maybe a series on tooling like this.



## Reference Links

 - [Build without Docker daemon](https://thenewstack.io/red-hat-buildah-provides-a-way-to-build-containers-without-the-docker-daemon/)
 - [Multicluster Controller/Operator Library](https://admiralty.io/blog/introducing-multicluster-controller/)
 - [Buildah Repository](https://github.com/containers/buildah)
 - [Kubebox](https://github.com/astefanutti/kubebox)
 - [Unshare(2) Man Page](http://man7.org/linux/man-pages/man2/unshare.2.html)
 - [great write up on how `top & free` work in containers](https://t.co/ZGpnKq2ddu)
 - [Docker on FreeBSD](https://wiki.freebsd.org/Docker)
 - [Docker Hack Day 2017](https://wiki.freebsd.org/DockerHackDay2017)
 - [Dave Cheney's Blog on Conditional Build Tags](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
 - [img build tool](https://github.com/genuinetools/img)
 - [kaniko - Build Images In Kubernetes](https://github.com/GoogleContainerTools/kaniko)
 - [Serverless Jenkins with Jenkins X](https://medium.com/@jdrawlings/serverless-jenkins-with-jenkins-x-9134cbfe6870)
 - [Building Docker Images Without Docker](https://www.youtube.com/watch?v=qhykcC94ukg)
 - [OpenShift Commons Briefing #102: Introducing Buildah with Nalin Dahyabhai & Dan Walsh](https://youtu.be/bOzJ9RJ4elU)
