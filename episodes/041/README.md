# Episode 41 : Traefik

- Hosted by @kris-nova
- Recording date

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url
Note the 0.jpg for the thumbnail --->

<a href="http://www.youtube.com/watch?feature=player_embedded&v=NObVcDG3ADM
" target="_blank"><img src="http://img.youtube.com/vi/NObVcDG3ADM/0.jpg"
alt="IMAGE ALT TEXT HERE" width="640" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Show Notes


## Reference Links

 - https://docs.traefik.io/user-guide/kubernetes/
 - https://github.com/kris-nova/krex
 - https://www.lfasiallc.com/events/kubecon-cloudnativecon-china-2018/program/call-for-proposals-cfp/
 - https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2018/
 - https://github.com/kubernetes/website/blob/release-1.11/content/en/blog/_posts/2018-06-26-kubernetes-1.11-release-announcement.md

## Stateful Application

This is the demo application Kris will use ot test ingress with Traefik.

Create DB

```
kubectl exec -it $(kubectl get po | grep postgres | cut -d " " -f 1) -- bash -c "psql -c 'CREATE DATABASE stateful_app_development;' -U postgres"
```

Scale deployment
```
kubectl scale deploy/statefulapp --replicas 1
```
