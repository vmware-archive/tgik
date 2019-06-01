#!/bin/bash -xe
for i in kind-control-plan{e,e2,e3}
do docker cp kubeadm ${i}:/usr/bin/kubeadm
done
