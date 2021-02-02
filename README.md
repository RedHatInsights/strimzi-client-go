# strimzi-client-go

Strimzi client-go package for use with kubernetes controllers.

This is a **work in progress** and hopefully the generation of this package will be fully automated
soon.

The types for this have been generated using using [crd-codegen](https://github.com/bsquizz/crd-codegen),
then the top-level type defintions have been tweaked to implement the runtime.Object interfaces and
add kubebuilder annotations, and finally `controller-gen object paths=./...` is run to generate the
`DeepCopy` implementations.

