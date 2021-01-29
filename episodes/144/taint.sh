# kubectl taint nodes $(hostname) node-role.kubernetes.io/master:NoSchedule-
# kubectl taint nodes clusterapi-peri-md-0-windows-containerd-5dc857845c-c98sv os=windows:NoSchedule-

kubectl taint nodes clusterapi-peri-md-0-windows-containerd-5dc857845c-c98sv os=windows:NoSchedule
