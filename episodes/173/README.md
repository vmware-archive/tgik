# Episode 173 : Pulumi Kubernetes Operator

- Hosted by [\@jbeda](http://twitter.com/jbeda)
- Guests: [\@briggsl](http://twitter.com/briggsl), [\@viveklak](http://twitter.com/viveklak)
- Recording date: 2021-11-05

<a href="https://www.youtube.com/watch?v=zvJ8PsSlchc" target="_blank"><img src="https://i.ytimg.com/vi/zvJ8PsSlchc/maxresdefault.jpg" border="10" /></a>

## Table of Contents

- 00:00:33 - Welcome to TGIK and intros
- 00:06:49 - Week in Review
- 00:21:21 - Digging into the Pulumi Kubernetes Operator
- 00:22:31 - What is Pulumi? How does it fit with other projects
- 00:29:44 - Pulumi pre-req including state file and secret management
- 00:38:53 - Installing the operator along with core arch
- 00:52:43 - Pulumi imperative/declarative split
- 01:04:18 - Applying the first stack (and debugging and Q&A!)
- 01:23:07 - Sidebar on impressions and where operator fits in
- 01:30:32 - Success! Dev wordpress running. Now looking at "prod" config.
- 01:47:55 - Secret generation and storage in Pulumi
- 02:07:00 - Success! And quick conclusions.

## Week in Review
* [Knative hits 1.0!](https://knative.dev/blog/articles/knative-1.0/)
* [GKE image streaming](https://cloud.google.com/blog/products/containers-kubernetes/introducing-container-image-streaming-in-gke)
    * [Containerd IPFS support](https://twitter.com/TokunagaKohei/status/1456561135725940741)
    * [Nercdtl IPFS support soon](https://github.com/containerd/nerdctl/issues/465)
* [Nice blog post on k8s logging](https://codersociety.com/blog/articles/kubernetes-logging)
* [Introducing Azure Container Apps](https://techcommunity.microsoft.com/t5/apps-on-azure/introducing-azure-container-apps-a-serverless-container-service/ba-p/2867265?ocid=AID3042118)
    * Related, [Dapr (Distributed Application Runtime) joins CNCF Incubator](https://www.cncf.io/blog/2021/11/03/dapr-distributed-application-runtime-joins-cncf-incubator/)

## Show Notes
* [Pulumi Kubernetes Operator Tutorial](https://www.pulumi.com/docs/guides/continuous-delivery/pulumi-kubernetes-operator/)
* [TGIK operator demo](https://github.com/jaxxstorm/tgik-operator-demo/)
* [Pulumi Kubernetes Operator Hits 1.0!](https://www.pulumi.com/blog/pulumi-kubernetes-operator-1-0/)
* [Policy as Code](https://www.pulumi.com/docs/guides/crossguard/)
* [Pulumi Registry](https://www.pulumi.com/registry/)

Pure OSS Pulumi references:
* [Self managed backend](https://www.pulumi.com/docs/intro/concepts/state/#logging-into-a-self-managed-backend)
* [Secret encryption](https://www.pulumi.com/docs/intro/concepts/secrets/#configuring-secrets-encryption)
* [Automation API](https://www.pulumi.com/docs/guides/automation-api/)
    * [Automation API Kubernetes Example](https://github.com/jaxxstorm/pulumi-productionapp)
* [Aliases](https://www.pulumi.com/blog/cumundi-guest-post/)


### Managing AWS secrets
Create secret from AWS env variables:
```
kubectl create secret generic aws-creds-secret \
    --from-literal=AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
    --from-literal=AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
```

Init env from AWS CLI config: https://github.com/jbeda/dotfiles/blob/master/source/50_aws_env.sh

## Download operator:
https://github.com/pulumi/pulumi-kubernetes-operator/archive/refs/tags/v1.2.1.tar.gz

## Using an S3 backend:
`pulumi login s3://<bucket>`

## CR docs
https://github.com/pulumi/pulumi-kubernetes-operator/blob/master/docs/stacks.md

```yaml
apiVersion: pulumi.com/v1
kind: Stack
metadata:
  name: tgik-demo
spec:
  backend: "<backend>"
  secretsProvider: "awskms:///<kms-key>"
  envRefs:
    AWS_ACCESS_KEY_ID:
      type: Secret
      secret:
        name: aws-creds-secret
        key: AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY:
      type: Secret
      secret:
        name: aws-creds-secret
        key: AWS_SECRET_ACCESS_KEY
    AWS_DEFAULT_REGION:
      type: Literal
      literal:
        value: "us-east-1"
  stack: "s3backend.tgik.dev"
  projectRepo: https://github.com/jaxxstorm/tgik-operator-demo
  branch: refs/heads/main
  destroyOnFinalize: true
```