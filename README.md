# strimzi-client-go

Go library providing Kubernetes CRD type definitions for the [Strimzi][strimzi] Kafka operator.
Enables typed Go clients for managing Strimzi resources on Kubernetes and OpenShift clusters.

## Features

Provides `runtime.Object`-compliant Go types for all 11 Strimzi CRDs across two API groups:

**core.strimzi.io/v1beta2**

- StrimziPodSet

**kafka.strimzi.io/v1beta2**

- Kafka
- KafkaBridge
- KafkaConnect
- KafkaConnector
- KafkaMirrorMaker
- KafkaMirrorMaker2
- KafkaNodePool
- KafkaRebalance
- KafkaTopic
- KafkaUser

## Installation

```bash
go get github.com/RedHatInsights/strimzi-client-go
```

## Usage

Register the Strimzi types with your client's scheme and use typed operations:

```go
package main

import (
    "context"
    "fmt"

    kafkav1beta2 "github.com/RedHatInsights/strimzi-client-go/apis/kafka.strimzi.io/v1beta2"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/client/config"
    "k8s.io/apimachinery/pkg/runtime"
)

func main() {
    scheme := runtime.NewScheme()
    _ = kafkav1beta2.AddToScheme(scheme)

    cfg, _ := config.GetConfig()
    c, _ := client.New(cfg, client.Options{Scheme: scheme})

    kafkaList := &kafkav1beta2.KafkaList{}
    _ = c.List(context.Background(), kafkaList)

    for _, kafka := range kafkaList.Items {
        fmt.Printf("Kafka cluster: %s/%s\n", kafka.Namespace, kafka.Name)
    }
}
```

## Type Generation

CRD types are generated from upstream Strimzi CRD YAML definitions, not written by hand.

### Prerequisites

- [crd-codegen][crd-codegen] -- generates Go types from CRD YAML
- [controller-gen][controller-gen] -- generates DeepCopyObject implementations
- Python 3 -- runs `convert.py` post-processing

### Generation Process

1. Download CRD YAML files from the target [Strimzi release][strimzi-releases]
2. Run crd-codegen to generate initial Go types:
   ```bash
   crd-codegen --apis-dir ./apis <crd-yaml-files>
   ```
3. Manually adjust package names and imports in generated files
4. Run the post-processing script:
   ```bash
   python3 convert.py
   ```
5. Generate DeepCopyObject implementations:
   ```bash
   controller-gen object paths=./apis/...
   ```

## Development

### Build

```bash
go build ./...
```

### Vet

```bash
go vet ./...
```

### Regenerate DeepCopy

```bash
controller-gen object paths=./apis/...
```

## Contributing

See [CONTRIBUTING.md][contributing] for guidelines.

## License

[Apache License 2.0][license]

[strimzi]: https://strimzi.io
[crd-codegen]: https://github.com/RedHatInsights/crd-codegen
[controller-gen]: https://pkg.go.dev/sigs.k8s.io/controller-tools/cmd/controller-gen
[strimzi-releases]: https://github.com/strimzi/strimzi-kafka-operator/releases
[contributing]: CONTRIBUTING.md
[license]: LICENSE
