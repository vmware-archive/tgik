#!/usr/bin/env bash
kind version >/dev/null
if [[ $? != 0 ]]; then
  echo "You need kind"
fi

footloose version >/dev/null
if [[ $? != 0 ]] ; then
  echo "you need footloose"
fi

ansible-playbook --version > /dev/null
if [[ $? != 0 ]] ; then
  echo "go get ansible :)"
fi

echo "let's set this stuff up"

mkdir -p /tmp/etcd-lab/{shared,etcd-cache}

footloose create
cd ansible
ansible-playbook playbook.yaml
cd ../
kind create cluster --config=kind/kind-lb0.yaml

echo "now you have a cluster with external etcd setup!"
