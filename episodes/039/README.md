# TGIK Episode #39

- Hosted by @jbeda
- Recording date: 2018-06-15

Note: Timestamps are [hh:mm] into the stream according to the local time, this might not match the final YouTube video timestamp exactly, but should get you close to the subject

## News and Announcements

- From now on we'll start the stream 5 minutes early to give people time to settle in.
- People ask where we get the cool Heptio swag, we'll work on how best to get this to you.
- The artwork on Joe's tshirt is courtesy of Ashley McNamara:
    - https://github.com/ashleymcnamara/gophers
    - https://twitter.com/ashleymcnamara


#### Kubeadm 1.11 notes (Thanks tstclair!)

[00:08] 1.11 is getting closer.  Daily burn down next week.

* kubeadm v1alpha2 Config
    * Massive refactor to align with plans to move to GA and enable mo-better clusterAPI integration
* Kubelet Config
    * Use Global Component Config w/versioning (no more unit file mucking)
* CoreDNS is now the default
* Better CRI integration and support for air-gapped environments
* Improved auto-detection for other CRIs.
* New CLI commands
    * `kubeadm config print-default`
    * `kubeadm config migrate`
    * `kubeadm config images list`
    * `kubeadm config images pull`
    * `kubeadm upgrade node config`
* Security improvements around kubelet
    * removing cadvisor port
    * Node taints on join, kubelets permissions are locked out after init/join
* Several bug fixes, most notably etcd upgrades

If you're interested in following along the kubeadm development, check out [SIG Cluster Lifecycle](https://github.com/kubernetes/community/tree/master/sig-cluster-lifecycle).

## News and Links from the Kubernetes Community

[00:11]

* https://github.com/kubernetes/community/blob/master/events/office-hours.md
* https://blog.heptio.com/ark-v0-9-alpha-now-with-restic-14ad6b402ab3
* https://docs.portworx.com/scheduler/kubernetes/ark.html
* [00:14] https://aws.amazon.com/eks/
* https://github.com/weaveworks/eksctl

## Related Links for today's session

* https://blog.heptio.com/congratulations-to-our-friends-at-microsoft-532cb41254f9
* [00:21] https://github.com/negz/kuberos
* [00:21] https://medium.com/@mrbobbytables/kubernetes-day-2-operations-authn-authz-with-oidc-and-a-little-help-from-keycloak-de4ea1bdbbe (Bob Killen, the author of this article will be participating in Kubernetes office hours regularly, so if you have auth questions bring them to office hours!)
* [00:22] https://github.com/appscode/guard

# Authenticating the your Cluster

[00:18] Joe will be covering adding github authentication to a kubernetes cluster in this episode using openid with dex.

* https://github.com/coreos/dex
* https://github.com/heptiolabs/gangway

Cloud providers don't provide this, the way to do this is to change flags in the API server. Their value add is the integration with their cloud, so if you're using AKS, you're using Active Directory out of the box, for example.

We'll cover using your own id provider 

## Contour and cert manager and DNS

[00:24] Let's start configuring and theory-crafting 
[00:26] Install contour:
```
kubectl apply -f yaml/01-contour
kubectl apply -f yaml/02-kuard
```

Get ELB address and point `*.tgik.io` and `tgik.io` to that in Route 53
```
kubectl get -n heptio-contour service contour -o wide
```

Now install cert manager and httpbin to confirm that TLS is all working.

```
kubectl apply -f yaml/03-cert-manager
kubectl apply -f yaml/04-httpbin
```

Now we should be able to go to https://httpbin.tgik.io.

These instructions/yaml are based on [Dave's article](https://blog.heptio.com/how-to-deploy-web-applications-on-kubernetes-with-heptio-contour-and-lets-encrypt-d58efbad9f56) and instructions in the [Contour repo](https://github.com/heptio/contour)

[00:29] How does OAuth work? (Notepad time)

- https://github.com/kubernetes-incubator/external-dns (User shout out to Zalando!)

## Dex and sample application

Now install dex.  Check out the dex config file at yaml/dex/03-dex-configmap.yaml
```
kubectl apply -f 05-dex
```

[00:52] The instructions for installing dex on Kubernetes aren't great. I had to do some experimentation to get to this point.

Navigate to https://dex.tgik.io/.well-known/openid-configuration

We should be able to point the example app at dex with:
```
bin/example-app --issuer https://dex.tgik.io/
```

Set up API server to talk to our instance of dex.  SSH in and modify api server manifest.
```
    - --oidc-issuer-url=https://dex.tgik.io/
    - --oidc-client-id=example-app
    - --oidc-username-claim=email
    - --oidc-groups-claim=groups
```

Configure kubectl directly here...
```
kubectl config set-credentials jbeda  \
    --auth-provider=oidc \
    --auth-provider-arg=idp-issuer-url=https://dex.tgik.io/ \
    --auth-provider-arg=client-id=example-app \
    --auth-provider-arg=client-secret=ZXhhbXBsZS1hcHAtc2VjcmV0 \
    --auth-provider-arg=refresh-token=Chljdmo0M21mYzZ5YnBzanJqYmRxc3g1dWYzEhl4NXJtNWFma2hvMjNzbHV6aWE1aHVpa3Zz \
    --auth-provider-arg=extra-scopes=groups

kubectl config set-context jbeda --user=jbeda --namespace=default --cluster=kubernetes
kubectl config use-context jbeda
kubectl get pods
kubectl config use-context admin@kubernetes
kubectl create rolebinding jbeda-admin --clusterrole=admin --user=joe.github@bedafamily.com
kubectl config use-context jbeda
```

## Gangway

Enable a new static client in dex.

Change the `client-id` on the API server...
```
    - --oidc-issuer-url=https://dex.tgik.io/
    - --oidc-client-id=gangway
    - --oidc-username-claim=email
    - --oidc-groups-claim=groups
```

Install gangway

```
kubectl apply -f 06-gangway
```

- [01:04] Checking out the redirects 
- [01:24] The server= line is empty
- [01:26] Logged in, and Joe can't `get pods` because .... RBAC! (This is a good thing)
- [01:28] Looks like Jose can log in
    - Simon gets: Error from server (Forbidden): pods is forbidden: User "simon@gmail.com" cannot list pods in the namespace "default"
    - RBAC is doing what it's supposed to. :+1: 
- [01:31] Let's try to give Simon access to Joe's cluster
- [01:33] Simon crushes Joe's wallet and launches a 100 instances. 
 
## Questions from the audience

- [01:00] Jim Angel asks "Do you think Dex will die out as RH merges Tectonic, Dex's main driving force, with OpenShift (since OpenShift already has auth using OCP cli + poor Dex documentation)?"
    - Bismark Paliz adds that Keycloak is also a Red Hat product
- [01:08] Any idea if we can use https://sovrin.org/? 
