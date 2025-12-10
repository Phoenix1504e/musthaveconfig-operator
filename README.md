# MustHaveConfig Operator

A simple Kubernetes Operator written in Go using Kubebuilder.

## What it does

Defines a custom resource:

```yaml
apiVersion: ops.aditya.dev/v1alpha1
kind: MustHaveConfig
metadata:
  name: my-app-config
  namespace: default
spec:
  namespace: default
  key: message
  value: "Hello from Operator"
