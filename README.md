# strimzi-client-go

Strimzi client-go package for use with kubernetes controllers.

This is a **work in progress** and hopefully the generation of this package will be fully automated
soon.

At the moment, the partially-automated process is:
1. Download the [strimzi CRD yaml](https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.21.1/strimzi-crds-0.21.1.yaml)
2. Edit the openAPI schemas to remove any 'oneOf' statements, since the GoLang type generator we
use does not handle those.
3. Edit the top-level type definitions so that they implement the runtime.Object interfaces and
add kubebuilder annotations. You may also see some of the "misc properties" type attributes defined
as `type <name> map[string]interface{}` -- these must be changed to `type <name> map[string]string`
4. Convert int types to int32: ```sed -i 's/int `json/int32 `json'/g *```
4. Finally run `controller-gen object paths=./...` is run to generate`DeepCopy` implementations

