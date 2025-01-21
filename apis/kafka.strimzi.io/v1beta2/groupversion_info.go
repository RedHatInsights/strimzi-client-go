/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package v1beta1 contains API Schema definitions for the kafka.strimzi.io v1beta1 API group
// +kubebuilder:object:generate=true
// +groupName=kafka.strimzi.io
package v1beta2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "kafka.strimzi.io", Version: "v1beta2"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(&Kafka{}, &KafkaList{})
	SchemeBuilder.Register(&KafkaBridge{}, &KafkaBridgeList{})
	SchemeBuilder.Register(&KafkaConnect{}, &KafkaConnectList{})
	SchemeBuilder.Register(&KafkaConnector{}, &KafkaConnectorList{})
	SchemeBuilder.Register(&KafkaMirrorMaker{}, &KafkaMirrorMakerList{})
	SchemeBuilder.Register(&KafkaMirrorMaker2{}, &KafkaMirrorMaker2List{})
	SchemeBuilder.Register(&KafkaNodePool{}, &KafkaNodePoolList{})
	SchemeBuilder.Register(&KafkaRebalance{}, &KafkaRebalanceList{})
	SchemeBuilder.Register(&KafkaTopic{}, &KafkaTopicList{})
	SchemeBuilder.Register(&KafkaUser{}, &KafkaUserList{})
}
