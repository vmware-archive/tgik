helm install --name harbor --namespace registry  --set externalPort='',persistence.enabled=false,externalDomain=harbor.c2.k8s.work,email.host=172.19.0.1,email.from="c2 <admin-c2@harbor.com>" .
