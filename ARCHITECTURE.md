# Architecture

## Overview

strimzi-client-go is a Go library that provides Kubernetes Custom Resource Definition (CRD) type
definitions for the [Strimzi][strimzi] Kafka operator. It enables Go programs to interact with
Strimzi-managed Kafka resources using strongly-typed clients instead of unstructured objects.

The library implements the [runtime.Object][runtime-object] and [DeepCopyObject][deepcopy]
interfaces required by the Kubernetes API machinery, allowing these types to be used with
controller-runtime clients, dynamic informers, and standard Kubernetes tooling.

## API Groups

The library organizes CRD types into two Kubernetes API groups, mirroring the upstream Strimzi
operator structure:

### core.strimzi.io/v1beta2

Located in `apis/core.strimzi.io/v1beta2/`. Contains infrastructure-level resources:

| Type | Description |
|------|-------------|
| StrimziPodSet | Manages pod sets for Strimzi components |

### kafka.strimzi.io/v1beta2

Located in `apis/kafka.strimzi.io/v1beta2/`. Contains Kafka workload resources:

| Type | Description |
|------|-------------|
| Kafka | Core Kafka cluster definition |
| KafkaBridge | HTTP bridge for Kafka |
| KafkaConnect | Kafka Connect cluster |
| KafkaConnector | Individual connector within Kafka Connect |
| KafkaMirrorMaker | MirrorMaker 1 replication |
| KafkaMirrorMaker2 | MirrorMaker 2 replication |
| KafkaNodePool | Node pool configuration for Kafka brokers |
| KafkaRebalance | Cruise Control rebalance proposals |
| KafkaTopic | Kafka topic management |
| KafkaUser | Kafka user and ACL management |

## Type Structure

Each CRD follows a consistent pattern:

```go
// Root type — implements runtime.Object
type Kafka struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec              *apiextensions.JSON `json:"spec,omitempty"`
    Status            *apiextensions.JSON `json:"status,omitempty"`
}

// List type — implements runtime.Object
type KafkaList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Kafka `json:"items"`
}
```

Key characteristics:

- **Spec and Status use `*apiextensions.JSON`** rather than strongly-typed structs. This avoids
  maintaining a full Go translation of the Strimzi CRD schema while still enabling typed
  client operations (create, get, list, watch, update, delete).
- **TypeMeta and ObjectMeta** are embedded following standard Kubernetes conventions.
- **List types** wrap a slice of the root type with `metav1.ListMeta`.
- **Kubebuilder markers** (`+kubebuilder:object:root=true`) on root and list types drive code
  generation for `runtime.Object` interface compliance.

## Scheme Registration

Each API group package contains a `groupversion_info.go` file that registers types with the
Kubernetes runtime scheme:

```go
var (
    SchemeGroupVersion = schema.GroupVersion{Group: "kafka.strimzi.io", Version: "v1beta2"}
    SchemeBuilder      = &scheme.Builder{GroupVersion: SchemeGroupVersion}
    AddToScheme        = SchemeBuilder.AddToScheme
)
```

Consumers call `AddToScheme` to register all types in the group with their client's scheme, enabling
typed client operations.

## Type Generation Flow

CRD types are not written by hand. The generation process is:

1. **Source CRD YAML** from the upstream Strimzi operator release
2. **crd-codegen** generates initial Go type files from the CRD YAML
3. **Manual edits** adjust generated output (package names, imports)
4. **convert.py** post-processes the generated files (formatting, field adjustments)
5. **controller-gen** generates `zz_generated.deepcopy.go` files containing `DeepCopyObject`
   implementations

The `zz_generated.deepcopy.go` files in each API group directory are fully generated and should
not be edited by hand.

## Directory Layout

```
apis/
├── core.strimzi.io/
│   └── v1beta2/
│       ├── groupversion_info.go      # Scheme registration
│       ├── StrimziPodSet.go          # CRD type definition
│       └── zz_generated.deepcopy.go  # Generated DeepCopyObject
└── kafka.strimzi.io/
    └── v1beta2/
        ├── groupversion_info.go      # Scheme registration
        ├── Kafka.go                  # One file per CRD type
        ├── KafkaBridge.go
        ├── KafkaConnect.go
        ├── KafkaConnector.go
        ├── KafkaMirrorMaker.go
        ├── KafkaMirrorMaker2.go
        ├── KafkaNodePool.go
        ├── KafkaRebalance.go
        ├── KafkaTopic.go
        ├── KafkaUser.go
        └── zz_generated.deepcopy.go  # Generated DeepCopyObject
convert.py       # Post-processing script for generated types
main.go          # Trivial demo / build verification
go.mod           # Module: github.com/RedHatInsights/strimzi-client-go
```

[strimzi]: https://strimzi.io
[runtime-object]: https://pkg.go.dev/k8s.io/apimachinery/pkg/runtime#Object
[deepcopy]: https://pkg.go.dev/k8s.io/apimachinery/pkg/runtime#Object
