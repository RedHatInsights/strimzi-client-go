# strimzi-client-go

Strimzi client-go package for use with kubernetes controllers.

This is a **work in progress** and hopefully the generation of this package will be fully automated
soon.

At the moment, the partially-automated process is:
1. Download the [strimzi CRD yaml](https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.24.0/strimzi-crds-0.24.0.yaml)
2. Edit the openAPI schemas to remove any 'oneOf' statements, since the GoLang type generator we
use does not handle those.
3. Use [crd-codegen](https://github.com/RedHatInsights/crd-codegen) to generate types from the CRD YAML.
4. Edit the top-level type definitions so that they implement the runtime.Object interfaces and
add kubebuilder annotations. You may also see some of the "misc properties" type attributes defined
as `type <name> map[string]interface{}` -- these must be changed to type `apiextensions.JSON` ([example](https://github.com/RedHatInsights/strimzi-client-go/commit/c6a1bf7c1dd6299e58a82ab1748b91036ac56e8d)). The python script `convert.py` can be used to help with this.
5. Convert int types to int32: ```sed -i 's/int `json/int32 `json'/g *```
6. Finally run `controller-gen object paths=./...` is run to generate `DeepCopy` implementations

