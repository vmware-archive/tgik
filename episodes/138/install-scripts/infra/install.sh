#!/bin/bash

kubectl apply -f https://projectcontour.io/quickstart/contour.yaml

kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.3/cert-manager.crds.yaml
kubectl create namespace cert-manager
kubectl create namespace external-dns
kubectl -n cert-manager create secret generic google-credentials --from-file=/tmp/key.json
kubectl -n external-dns create secret generic google-credentials --from-file=/tmp/key.json

helm repo add jetstack https://charts.jetstack.io
helm repo add bitnami https://charts.bitnami.com/bitnami

helm install cert-manager --namespace cert-manager --values cert-manager.yaml jetstack/cert-manager


helm install mexternal-dns bitnami/external-dns --namespace external-dns --version 3.5.0 --values external-dns.yaml

kubectl apply -f ./cluster-issuer.yaml
