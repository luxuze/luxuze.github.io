```yml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: sa-micro-services
  namespace: ns-micro
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: micro-registry
  namespace: ns-micro
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - list
      - patch
      - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: micro-registry
  namespace: ns-micro
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: micro-registry
subjects:
  - kind: ServiceAccount
    name: sa-micro-services
    namespace: ns-micro
```
