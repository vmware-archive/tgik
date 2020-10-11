. ./config.sh

scp ubuntu@${CP}:.kube/config kubeconfig
export KUBECONFIG=$(pwd)/kubeconfig
kubectl config set clusters.kubernetes.server https://127.0.0.1:6443
kubectl config set clusters.kubernetes.tls-server-name 10.96.0.1

ssh -N -L 6443:localhost:6443 ubuntu@$CP &
