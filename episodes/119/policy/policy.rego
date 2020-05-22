package main

warn[msg] {
  input.apiVersion == "v1"
  input.kind == "List"
  obj := input.items[_]
  msg := _warn with input as obj
}

warn[msg] {
  input.apiVersion != "v1"
  input.kind != "List"
  msg := _warn
}

# Based on https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md

# Ingress resources will no longer be served from extensions/v1beta1 in v1.20. Migrate use to the networking.k8s.io/v1beta1 API, available since v1.14.
_warn = msg {
  input.apiVersion == "extensions/v1beta1"
  input.kind == "Ingress"
  msg := sprintf("%s/%s: API extensions/v1beta1 for Ingress is deprecated from Kubernetes 1.20, use networking.k8s.io/v1beta1 instead.", [input.kind, input.metadata.name])
}

# All resources will no longer be served from rbac.authorization.k8s.io/v1alpha1 and rbac.authorization.k8s.io/v1beta1 in 1.20. Migrate to use rbac.authorization.k8s.io/v1 instead
_warn = msg {
  apis := ["rbac.authorization.k8s.io/v1alpha1", "rbac.authorization.k8s.io/v1beta1"]
  input.apiVersion == apis[_]
  msg := sprintf("%s/%s: API %s is deprecated from Kubernetes 1.20, use rbac.authorization.k8s.io/v1 instead.", [input.kind, input.metadata.name, input.apiVersion])
}

# Ingress resources will no longer be served from extensions/v1beta1 in v1.20. Migrate use to the networking.k8s.io/v1beta1 API, available since v1.14.
_warn = msg {
  input.apiVersion == "extensions/v1beta1"
  input.kind == "Ingress"
  msg := sprintf("%s/%s: API extensions/v1beta1 for Ingress is deprecated from Kubernetes 1.20, use networking.k8s.io/v1beta1 instead.", [input.kind, input.metadata.name])
}