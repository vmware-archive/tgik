

helm repo add harbor https://helm.goharbor.io

kubectl create namespace harbor

helm install harbor -n harbor harbor/harbor --values values-harbor.yaml
