# strimzi-client-go

Strimzi client-go package for use with kubernetes controllers.

This is a **work in progress** and hopefully the generation of this package will be fully automated
soon.

At the moment, the partially-automated process is:

1. Download the [strimzi CRD yaml](https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.24.0/strimzi-crds-0.24.0.yaml)

1. Edit the openAPI schemas to remove any 'oneOf' statements, since the GoLang type generator we
use does not handle those.

1. Use [crd-codegen](https://github.com/RedHatInsights/crd-codegen) to generate types from the CRD YAML.

2. Edit the root type definitions in each file so that they implement the runtime.Object interfaces, add a "List" type for the root object, and add kubebuilder annotations. For example, the root "Kafka" object within Kafka.go will initially look like this:

   ```golang
   package v1beta2

   import "encoding/json"
   import "fmt"
   import "reflect"

   type Kafka struct {
       // The specification of the Kafka and ZooKeeper clusters, and Topic Operator.
       Spec *KafkaSpec `json:"spec,omitempty" yaml:"spec,omitempty" mapstructure:"spec,omitempty"`

       // The status of the Kafka and ZooKeeper clusters, and Topic Operator.
       Status *KafkaStatus `json:"status,omitempty" yaml:"status,omitempty" mapstructure:"status,omitempty"`
   }
   ```

   and it will need to be updated to look like this:

   ```golang
   package v1beta2

   import (
       "encoding/json"
       "fmt"
       "reflect"

       metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

       apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
   )

   // +kubebuilder:object:root=true
   // +kubebuilder:subresource:status
   // Kafka
   type Kafka struct {
       metav1.TypeMeta   `json:",inline"`
       metav1.ObjectMeta `json:"metadata,omitempty"`

       // The specification of the Kafka and ZooKeeper clusters, and Topic Operator.
       Spec *KafkaSpec `json:"spec,omitempty" yaml:"spec,omitempty" mapstructure:"spec,omitempty"`

       // The status of the Kafka and ZooKeeper clusters, and Topic Operator.
       Status *KafkaStatus `json:"status,omitempty" yaml:"status,omitempty" mapstructure:"status,omitempty"`
   }

   // +kubebuilder:object:root=true
   // KafkaList contains a list of instances.
   type KafkaList struct {
       metav1.TypeMeta `json:",inline"`
       metav1.ListMeta `json:"metadata,omitempty"`

       // A list of Kafka objects.
       Items []Kafka `json:"items,omitempty"`
   }
   ```

3. Some unstructured/"free-form" fields may be defined as `type <name> map[string]interface{}` and these must be changed to type `apiextensions.JSON`. The python script `convert.py` can be ran against each file to help with this.

4. Convert int types to int32: ```sed -i 's/int `json/int32 `json'/g *```

5. Finally run `controller-gen object paths=./...` is run to generate `DeepCopy` implementations
