## Kubeadm 1.11 notes (Thanks tstclair!)

1.11 is getting closer.  Daily burn down next week.

* kubeadm v1alpha2 Config
    * Massive refactor to align with plans to move to GA and enable mo-better clusterAPI integration
* Kubelet Config
    * Use Global Component Config w/versioning (no more unit file mucking)
* CoreDNS is now the default
* Better CRI integration and support for air-gapped environments
* Improved auto-detection for other CRIs.
* New CLI commands
    * kubeadm config print-default
    * kubeadm config migrate
    * kubeadm config images list
    * kubeadm config images pull
    * kubeadm upgrade node config
* Security improvements around kubelet
    * removing cadvisor port
    * Node taints on join, kubelets permissions are locked out after init/join
* Several bug fixes, most notably etcd upgrades

## Links
* https://github.com/kubernetes/community/blob/master/events/office-hours.md
* https://blog.heptio.com/ark-v0-9-alpha-now-with-restic-14ad6b402ab3
* https://docs.portworx.com/scheduler/kubernetes/ark.html
* https://aws.amazon.com/eks/
* https://github.com/weaveworks/eksctl
* https://blog.heptio.com/congratulations-to-our-friends-at-microsoft-532cb41254f9
* https://github.com/negz/kuberos
https://medium.com/@mrbobbytables/kubernetes-day-2-operations-authn-authz-with-oidc-and-a-little-help-from-keycloak-de4ea1bdbbe
* https://github.com/appscode/guard
* https://github.com/coreos/dex
* https://github.com/heptiolabs/gangway

## Contour and cert manager and DNS

Install contour:
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

## Dex and sample application

Now install dex.  Check out the dex config file at yaml/dex/03-dex-configmap.yaml
```
kubectl apply -f 05-dex
```

The instructions for installing dex on Kubernetes aren't great. I had to do some experimentation to get to this point.

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