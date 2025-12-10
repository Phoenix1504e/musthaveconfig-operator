# MustHaveConfig Operator

This is a simple Kubernetes operator written in Go using controller-runtime.  
I built this project to understand how Kubernetes operators work at a practical level — from CRDs and reconciliation to drift correction and status updates.

The operator ensures that a specific ConfigMap **must always exist** with a given key/value pair, as defined by a custom resource.

---

## What this operator does

The operator introduces a custom resource called **MustHaveConfig**.

When a `MustHaveConfig` resource is created or updated, the operator:

- Ensures a ConfigMap with the same name exists  
- Creates the ConfigMap if it does not exist  
- Ensures the ConfigMap contains the specified key/value  
- Automatically corrects any drift if the ConfigMap is modified manually  
- Updates the CR status to reflect whether the resource is in sync  

When the custom resource is deleted, the associated ConfigMap is automatically cleaned up.

---

## Example Custom Resource

```yaml
apiVersion: ops.aditya.dev/v1alpha1
kind: MustHaveConfig
metadata:
  name: demo-config
  namespace: default
spec:
  namespace: default
  key: message
  value: "Hello from Operator"
```
