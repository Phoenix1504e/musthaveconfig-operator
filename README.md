# MustHaveConfig Kubernetes Operator

A custom Kubernetes Operator built using Go and Kubebuilder to enforce and manage required configuration resources within a Kubernetes cluster.

This project demonstrates hands-on experience with Kubernetes extensibility, Custom Resource Definitions (CRDs), and controller-based reconciliation logic.

---

## Overview

The **MustHaveConfig Operator** introduces a custom Kubernetes resource that represents mandatory configuration requirements for workloads running in a cluster.

The operator continuously watches the cluster state and reconciles it to ensure that required configuration objects exist and remain consistent with the desired specification.

This project focuses on understanding how Kubernetes controllers work internally rather than providing a generic example.

---

## Architecture

The operator is built using standard Kubernetes controller patterns:

- **Custom Resource Definition (CRD)**  
  Defines the `MustHaveConfig` resource and its schema.

- **Controller (Reconciler)**  
  Watches for changes to `MustHaveConfig` resources and reconciles the actual cluster state.

- **Controller Runtime (Kubebuilder)**  
  Provides abstractions for managing watches, reconciliation loops, and client interactions.

---

## Custom Resource

### Kind
```
MustHaveConfig
```

### Example Manifest

```yaml
apiVersion: config.example.com/v1alpha1
kind: MustHaveConfig
metadata:
  name: example-config
spec:
  foo: "required-value"
```

---

## Reconciliation Flow

1. Kubernetes API server detects a change to a `MustHaveConfig` resource
2. The controller’s reconcile loop is triggered
3. Desired state is compared with actual cluster state
4. Missing or inconsistent configuration is corrected
5. The cluster converges toward the declared desired state

---

## Project Structure

```
.
├── api/
│   └── v1alpha1/
│       └── musthaveconfig_types.go   # CRD schema
├── controllers/
│   └── musthaveconfig_controller.go  # Reconciliation logic
├── config/                            # Kubernetes manifests
├── main.go                            # Operator entrypoint
├── Makefile
└── README.md
```

---

## Key Technical Learnings

- Kubernetes controller and reconciliation patterns
- Designing and versioning Custom Resource Definitions (CRDs)
- Using Kubebuilder and controller-runtime
- Managing desired vs actual state in Kubernetes
- Debugging controller behavior and API interactions
- Understanding Kubernetes API machinery and runtime objects

---

## Why Operators?

Kubernetes Operators allow teams to:

- Extend Kubernetes APIs safely
- Encode operational knowledge into software
- Automate configuration enforcement
- Maintain cluster consistency declaratively

This pattern is widely used for managing databases, infrastructure components, and platform-level services.

---

## Possible Enhancements

- Status subresource updates
- Validation webhooks
- Support for multiple configuration targets
- Metrics and observability
- Policy-based enforcement logic

---

## Status

This operator is functional and designed as a learning-focused implementation of Kubernetes controller patterns using Go and Kubebuilder.

The project emphasizes understanding Kubernetes internals and real-world operator development workflows.
