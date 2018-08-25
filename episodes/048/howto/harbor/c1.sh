helm install --name harbor --namespace registry  --set externalPort='',persistence.enabled=false,externalDomain=harbor.c1.k8s.work,email.host=172.18.0.1,email.from="c1 <admin-c1@harbor.com>" .
