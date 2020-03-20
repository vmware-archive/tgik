.DEFAULT_GOAL:=help
.ONESHELL:

SHELL := /bin/bash
KUBECONFIG := $(cat ./resources/kpath)
KINDCFG := ./resources/kind.yaml
CLUSTERCTL := ./clusterctl.yaml
CV3 := capiv
CV3IMG := ./resources/images.list

C?=$(CV3)
export KUBECONFIG


help:  ## Display this help
	echo $(KUBECONFIG)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

#KUBERNETES STUFF
.PHONY: images.list
images.list:
	clusterctl --config clusterctl.yaml init -i docker --list-images \
	> $(CV3IMG)

.PHONY: cache_capi
cache_k8s:
	cat $(CV3IMG) | xargs -n3 -I {} docker pull {}


.PHONY: init_cv3
init_cv3:
	kind get clusters | grep -qa $(CV3) \
	|| kind create cluster --name=$(CV3) --config=$(KINDCFG)

.PHONY: load_cv3
load_cv3: init_cv3
	cat $(CV3IMG) | xargs -I {} kind load docker-image {} --name $(CV3)

.PHONY: load
load:
	cat $(CV3IMG) | xargs -I {} kind load docker-image {} --name $(C)

.PHONY: clusterctl_init
clusterctl:
	clusterctl init -v10 --config=$(CLUSTERCTL) -i docker

define kconfigs =
	rm -f ${PWD}/resources/kpath
	for k in $(kind get clusters)
		do 
			kind get kubeconfig --name=$k > ./resources/${k}-kubeconfig
 			echo -n ${PWD}/resources/${k}-kubeconfig:>>./resources/kpath
		done
endef

.PHONY: combined-kubeconfig
combined-kubeconfig:
	$(value kconfigs)

.PHONY: clean_all
clean_all:
	kind get clusters | xargs -i kind delete cluster --name={} || exit 0
	rm -f ./resources/*kubeconfig

.PHONY: clean_docker
clean_docker:
	docker ps -q | xargs docker rm -f 
