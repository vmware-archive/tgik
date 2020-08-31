# Episode 073 : Exploring Garden with Kubernetes

- Hosted by @krisnova
- Recording date: 2019-04-19

Episode issue: 
<!--- Thumbnailed embed of the video, NFnpUlt0IuM is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=Xfi9XqcZ76M
" target="_blank"><img src="http://img.youtube.com/vi/Xfi9XqcZ76M/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:04:00 - Week in Review


## Week in Review

- Lee Briggs compares [using Fargate](https://leebriggs.co.uk/blog/2019/04/13/the-fargate-illusion.html) from a K8s user standpoint
- Tinder's notes on [moving to Kubernetes](https://medium.com/@tinder.engineering/tinders-move-to-kubernetes-cda2a6372f44)
- k8s [1.14.1 is out](https://github.com/kubernetes/kubernetes/releases/tag/v1.14.1) 
    - 1.15 code freeze is next week. I am not kidding.
- How [godaddy manages external secrets](https://godaddy.github.io/2019/04/16/kubernetes-external-secrets/)
- Youichi Fujimoto has a fast guide to [using kind](https://itnext.io/starting-local-kubernetes-using-kind-and-docker-c6089acfc1c0)
    - Check out the [kind here](https://kind.sigs.k8s.io/)
- Comparing [local development](https://codefresh.io/kubernetes-tutorial/local-k8s-draft-skaffold-garden/) with Dreaf, Skaffold, and Garden 

## Show Notes

```bash
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole cluster-admin \
  --user $(gcloud config get-value account)
```

Garden example project, [`go3dprint`](https://github.com/garden-io/go3dprint/).


## Reference Links

 - [Garden GitHub](https://github.com/garden-io/garden)
 - [Quick Start with Garden](https://docs.garden.io/basics/quick-start)
 - [Blogs on Garden](https://medium.com/garden-io)
