## for this episode i used kubeadm dind to host harbor.

git clone git@github.com:kubernetes-sigs/kubeadm-dind-cluster.git

cd kubeadm-dind-cluster

I am also using direnv to manage my environment variables. direnv.net

you can just 

```
export DNS_SERVICE="coredns"
export CNI_PLUGIN="calico-kdd"
export SKIP_SNAPSHOT="y"
```
and then to start your cluster you can do 
`./fixed/dind-v1.11.sh up`

Then you should see the following after the up completes:

```
 $ kubectl get nodes -o wide
NAME          STATUS    ROLES     AGE       VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE                       KERNEL-VERSION      CONTAINER-RUNTIME
kube-master   Ready     master    3m        v1.11.0   172.18.0.2    <none>        Debian GNU/Linux 9 (stretch)   4.15.0-32-generic   docker://17.3.2
kube-node-1   Ready     <none>    2m        v1.11.0   172.18.0.3    <none>        Debian GNU/Linux 9 (stretch)   4.15.0-32-generic   docker://17.3.2
kube-node-2   Ready     <none>    2m        v1.11.0   172.18.0.4    <none>        Debian GNU/Linux 9 (stretch)   4.15.0-32-generic   docker://17.3.2
```

Then I made some certs using certbot and manual verification mode since I own k8s.work I make a subdomain c1.k8s.work and put up the let encrypt secret in a TXT record:
```
*.c1.k8s.work.               A     172.18.0.3,172.18.0.4
_acme-challenge.c1.k8s.work. TXT   "P5QEuyPKdRM41uD61IWQnWtRJUKlxqBfk1wbz3lfzGc"
```

The * record will attract everything to my 2 worker node ip addresses. the acme-challenge is there so that i can issues certs against that subdomain using certbot from Lets Encrypt.

So I generate a * record and then I push the downloaded fullchain.crt and privkey to a secret in the namespace ingress with:
```
kubectl create ns ingress
kubectl create secret tls -n ingress defaultssl --cert=certs/c1.k8s.work/fullchain.pem --key=certs/c1.k8s.work/privkey.pem
```

Now down to it.

In this dir under harbor/ingress/ is the expanded nginx-ingress stable chart.
You can use helm to template some manifests here.

I use this command in the show:
```
 helm template --name std-ingress --namespace ingress --set controller.kind=DaemonSet,controller.hostNetwork=true,controller.extraArgs.default-ssl-certificate=ingress/defaultssl,controller.service.type=None .
```

This will template the manifests.

You can install them with:
```
helm template --name std-ingress --namespace ingress --set controller.kind=DaemonSet,controller.hostNetwork=true,controller.extraArgs.default-ssl-certificate=ingress/defaultssl,controller.service.type=None . | kubectl apply -f - --namespace=ingress
```

Now you should have an ingress controller running on each worker node serving with Hostnet: true means that each of the worker nodes now have an ingress controller  process using node port 80 and 443.

Next up testing:

The command I showed on this was a curl command:

curl --resolve blah.c1.k8s.work:443:172.18.0.3 https://blah.c1.k8s.work


Once all that worked I moved on to bringing up the harbor chart. For this I used the chart that is in this dir under harbor.

This chart has since been removed from the harbor repo. I don't yet know why?

The command to deploy to cluster 1 for harbor is:

```
helm init
helm install --name harbor --namespace registry  --set externalPort='',persistence.enabled=false,externalDomain=harbor.c1.k8s.work,email.host=172.18.0.1,email.from="c1 <admin-c1@harbor.com>" .
```

This will stand up harbor in your cluster without persistence.


After that we played with Harbor a bit and then did some docker things.

To grab the ca key here was the command:

```
kubectl get cm -n registry harbor-harbor-notary -o jsonpath="{.data.notary-signer-ca\\.crt}" > ~/.docker/tls/notary-harbor.c1.k8s.work/ca.crt

Then the env var to make docker point to our notary:

```
export DOCKER_CONTENT_TRUST=1
export DOCKER_CONTENT_TRUST_SERVER=https://notary-harbor.c1.k8s.work
```
and we are off to the races:

```
docker trust sign harbor.c1.k8s.work/library/kubeadm:v1.11.1
docker push !$
```


