
# MustHaveConfig Operator

A small Kubernetes operator written in Go using Kubebuilder.  
This project was built to learn how reconciliation loops work and how custom resources can control real cluster objects.

## What it does

It introduces a custom resource named **MustHaveConfig**, which lets you declare that a certain ConfigMap should always exist with a specific key/value pair.

Example CR:

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
````

When this resource is created, the operator will:

* Look for a ConfigMap named `my-app-config` in the target namespace
* Create it if it's missing
* Ensure `data[spec.key]` always equals `spec.value`
* Update the CR status to show whether the sync was successful

If the ConfigMap already exists but with a different value, the operator updates it automatically.

## Why I built this

I wanted to understand the basics of Kubernetes operators:

* How custom resources are defined in Go
* How reconciliation loops enforce desired state
* Working with controller-runtime APIs and client interactions

This is not meant to be production ready — just a learning exercise.

## How to run (later)

The operator compiles, and next steps will be:

* Running it inside a local `kind` or `minikube` cluster
* Applying sample MustHaveConfig resources
* Watching ConfigMaps update automatically

## Tech used

* Go
* Kubebuilder
* Controller Runtime
* Kubernetes API
* WSL (Ubuntu)
