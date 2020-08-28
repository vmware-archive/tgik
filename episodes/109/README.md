# Episode 109 : Kubernetes Dev Setup

- Hosted by @mauilion
- Recording date: 2020-03-13

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=Rce70OkfRhI
" target="_blank"><img src="http://img.youtube.com/vi/Rce70OkfRhI/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:00:00 - Week in Review

## Week in Review

### K8s Core

- k8s 1.18 is now [beta2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md/#changelog-since-v1180-beta1)
- Debug containers are here; try [kubectl alpha debug](https://github.com/kubernetes/enhancements/blob/c283365bfdc2abd035b0dc0480e9e096e4d4b522/keps/sig-cli/20190805-kubectl-debug.md)

### K8s Ecosystem

- Rich Stokes [has some goodies](https://github.com/richstokes/k8s-fah) on deploying Folding@Home on Kubernetes to help solve for COVID-19
- Play a supervillain with [kubethanos](https://github.com/berkay-dincer/kubethanos)



## Outline:

- [x] What is a contribution? [community](https://kubernetes.io/community)
    - [x] discuss.k8s.io
    - [x] slack.k8s.io
    - [x] open bugs!
    - [x] test things!
    - [x] Share your knowledge!
- [x] Contribute to docs! [how?](https://kubernetes.io/docs/contribute/start/#the-basics-about-our-docs)
- [x] Help with doc [translations!](https://kubernetes.io/docs/contribute/localization/)
- [x] How to contribute to the [codebase](https://github.com/kubernetes/community/tree/master/contributors/guide/contributor-cheatsheet#getting-started)
- [x] Check out the old doc.
- [x] Frame this with what it takes today!
- [x] Let's get to it!
    - [x] prereqs:
        - [x] git
        - [x] docker
        - [x] Caveats for Docker!
    - [x] direnv.net
    - [x] [gimme](https://github.com/travis-ci/gimme)
    - [x] go get!
        - Kubernetes DNS configs: https://git.k8s.io/k8s.io/dns
        - Vanity redirects (https://git.k8s.io/k8s.io/k8s.io)
            - go imports (`k8s.io/foo` --> `github.com/kubernetes/foo` and `sigs.k8s.io/bar` --> `github.com/kubernetes-sigs/bar`)
            - GitHub website redirects (`git.k8s.io` --> `github.com/kubernetes`)
    - [x] let's talk about building stuff!
    - [x] let's play!
- [x] tests! [Kind e2e](https://github.com/kubernetes-sigs/kind/blob/master/hack/ci/e2e-k8s.sh)
    - [x] [testing docs](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/testing.md)
    - [x] power tool! 
      ```
      build/run.sh make test
      ```
- [ ] commit messages! [write up](https://chris.beams.io/posts/git-commit/)


## Show Notes

General idea:

The web is littered with "setting up your development laptop for cloud native". 
Lots of them get out of date quickly and need love, like the page below which we've copied from the k8s developer guide.
How do the experts do it? What are their tricks, why is it when they do a talk they have all these awesome shortcuts, useful prompts, namespace switching thingers, and all these goodies? 

Duffie is going to set up his laptop the awesome way, and then as he's doing it we're going to rewrite this kubernetes contributor guide page. We're going to add all of your goodies and tips too, then when we're done, we're going to commit it back upstream -- so that the community defaults to the awesome setup. 


- Bonus: If you've never contributed to k8s when we're done one of you can submit this page as your very first pull request and @castrojo will review and help you get it up there. 

## Reference Links

Duffie will set up his laptop

Instructions from: https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md


Getting started locally
-----------------------


### Requirements

git
golang: https://github.com/travis-ci/gimme


#### Linux

`sudo apt install git docker.io`


#### MacOS

`xcode-select --install`

`brew install git`

`brew install docker` or [Docker Desktop for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac)

`brew install gimme`
bash if you're in Catalina

#### Windows





#### Docker

At least [Docker](https://docs.docker.com/installation/#installation)
1.10+. Ensure the Docker daemon is running and can be contacted (try `docker
ps`).  Some of the Kubernetes components need to run as root, which normally
works fine with docker.


### Clone the repository

In order to run kubernetes you must have the kubernetes code on the local machine. Cloning this repository is sufficient.

```$ git clone --depth=1 https://github.com/kubernetes/kubernetes.git```

The `--depth=1` parameter is optional and will ensure a smaller download.

### Starting the cluster

In a separate tab of your terminal, run the following (since one needs sudo access to start/stop Kubernetes daemons, it is easier to run the entire script as root):

```sh
cd kubernetes
hack/local-up-cluster.sh
```

This will build and start a lightweight local cluster, consisting of a master
and a single node. Type Control-C to shut it down.

If you've already compiled the Kubernetes components, then you can avoid rebuilding them with this script by using the `-O` flag.

```sh
./hack/local-up-cluster.sh -O
```

You can use the cluster/kubectl.sh script to interact with the local cluster. hack/local-up-cluster.sh will
print the commands to run to point kubectl at the local cluster.


### Running a container

Your cluster is running, and you want to start running containers!

You can now use any of the cluster/kubectl.sh commands to interact with your local setup.

```sh
cluster/kubectl.sh get pods
cluster/kubectl.sh get services
cluster/kubectl.sh get replicationcontrollers
cluster/kubectl.sh run my-nginx --image=nginx --replicas=2 --port=80


## begin wait for provision to complete, you can monitor the docker pull by opening a new terminal
  sudo docker images
  ## you should see it pulling the nginx image, once the above command returns it
  sudo docker ps
  ## you should see your container running!
  exit
## end wait

## introspect Kubernetes!
cluster/kubectl.sh get pods
cluster/kubectl.sh get services
cluster/kubectl.sh get replicationcontrollers
```


### Running a user defined pod

Note the difference between a [container](https://kubernetes.io/docs/user-guide/containers/)
and a [pod](https://kubernetes.io/docs/user-guide/pods/). Since you only asked for the former, Kubernetes will create a wrapper pod for you.
However you cannot view the nginx start page on localhost. To verify that nginx is running you need to run `curl` within the docker container (try `docker exec`).

You can control the specifications of a pod via a user defined manifest, and reach nginx through your browser on the port specified therein:

```sh
cluster/kubectl.sh create -f test/fixtures/doc-yaml/user-guide/pod.yaml
```

Congratulations!

### Troubleshooting

#### I cannot reach service IPs on the network.

Some firewall software that uses iptables may not interact well with
kubernetes.  If you have trouble around networking, try disabling any
firewall or other iptables-using systems, first.  Also, you can check
if SELinux is blocking anything by running a command such as `journalctl --since yesterday | grep avc`.

By default the IP range for service cluster IPs is 10.0.*.* - depending on your
docker installation, this may conflict with IPs for containers.  If you find
containers running with IPs in this range, edit hack/local-cluster-up.sh and
change the service-cluster-ip-range flag to something else.

#### I cannot create a replication controller with replica size greater than 1!  What gives?

You are running a single node setup.  This has the limitation of only supporting a single replica of a given pod.  If you are interested in running with larger replica sizes, we encourage you to try the local vagrant setup or one of the cloud providers.

#### I changed Kubernetes code, how do I run it?

```sh
cd kubernetes
make
hack/local-up-cluster.sh
```

#### kubectl claims to start a container but `get pods` and `docker ps` don't show it.

One or more of the Kubernetes daemons might've crashed. Tail the logs of each in /tmp.

#### The pods fail to connect to the services by host names

To start the DNS service, you need to set the following variables:

```sh
KUBE_ENABLE_CLUSTER_DNS=true
KUBE_DNS_SERVER_IP="10.0.0.10"
KUBE_DNS_DOMAIN="cluster.local"
```

To know more on DNS service you can check out the [docs](http://kubernetes.io/docs/admin/dns/).
