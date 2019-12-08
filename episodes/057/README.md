# Episode 057 : Vault and Kubernetes

- Hosted by @kris-nova
- Recording date: 20181117

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=7aR6k9xBN94
" target="_blank"><img src="http://img.youtube.com/vi/7aR6k9xBN94/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!


## Show Notes

Auth done using the Kubernetes Auth backend:
https://www.vaultproject.io/docs/auth/kubernetes.html

`jwt` is something of an under-description. In granting the service account token to Vault, you are authing using a shared-secret credential that let's one use the serviceaccount generally; so be aware (but if you can't trust Vault or the wire to it ... you're in trouble)

When configuring `kubernetes_ca_cert=@ca.crt` it is worth keeping in mind this part of the [`write`](https://www.vaultproject.io/docs/commands/write.html) documentation

> Data is specified as "key=value" pairs. If the value begins with an "@", then it is loaded from a file. If the value is "-", Vault will read the value from stdin.
> 
As for the value you could extract from your kubeconfig as e.g.

```bash
kubectl config view --raw -o json | \
jq -r '. as $raw | .clusters[] | 
select(
  .name == (
    $raw.contexts[] | select( .name == ($raw."current-context")
  ) | .context.cluster)
) | .cluster | ."certificate-authority-data"' | \
base64 --decode > ca.crt
# or there may be a key "certificate-authority" which
# will contain a file name you can just use with @
```

## Reference Links

 ```bash
 sed -i s/Heptio/Heptio + VMware/g
 ```
 Thanks to Justin from Lithuania for our keychains. TODO: we should check in a photo. 
 
 - [1.13alpha3 is out](https://github.com/kubernetes/kubernetes/releases/tag/v1.13.0-alpha.3) - 1.13 is in code freeze, last release of the year!
 - [VMware update from Craig](https://blog.heptio.com/heptio-will-be-joining-forces-with-vmware-on-a-shared-cloud-native-mission-b01225b1bc9e)
 - [Nova at Heptio](https://twitter.com/krisnova/status/932685715934187520?s=21)
 - [Dockerbox (Seth Pollack)](https://github.com/sethpollack/dockerbox) - a busybox for Docker
 - [Heptio Intro to Containers and Kubernetes training](https://www.eventbrite.com/e/intro-to-containers-and-kubernetes-seattle-tickets-50999092659?utm_medium=youtube&utm_source=tgik&utm_campaign=tgik_heptio_training)
 - [Kelsey Hightower's Vault on GKE](https://github.com/kelseyhightower/vault-on-google-kubernetes-engine)
  - [Vault Operator](https://github.com/coreos/vault-operator)
 - [Dynamic Secrets](https://www.hashicorp.com/blog/why-we-need-dynamic-secrets)
 - [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
 - [Service Account Token](https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration#authenticating-inside-the-cluster)
 - [Shamir's Secret (Seal/Unseal)](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing)
 - [Dynamic Secrets](https://medium.com/@gmaliar/dynamic-secrets-on-kubernetes-pods-using-vault-35d9094d169)
 - [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
 - [Kubicorn](https://github.com/kubicorn/kubicorn)
 - [Klone](https://github.com/kris-nova/klone)
