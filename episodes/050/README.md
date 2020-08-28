# Episode 050 : Weave (CNI)

- Hosted by @krisnova
- Recording date: 2018-09-07
- HackMD: https://hackmd.io/Vd_cF06PRvSovEoMvn2Riw?both

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url
Note the 0.jpg for the thumbnail --->

<a href="http://www.youtube.com/watch?feature=player_embedded&v=2YoK4bBy3CM
" target="_blank"><img src="http://img.youtube.com/vi/2YoK4bBy3CM/0.jpg"
alt="IMAGE ALT TEXT HERE" width="640" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Show Notes

Deploy a cluster without Weave using Kubciron

```bash

# From this directory
kubicorn create tgik-weave-cluster -S kubicorn -p aws -M serverPool.bootstrapScripts[0]=kubicorn/amazon_k8s_ubuntu_16.04_master.sh
kubicorn apply tgik-weave-cluster -S kubicorn

```

## Reference Links

 - https://medium.com/google-cloud/managing-helm-releases-the-gitops-way-207a6ac6ff0e
