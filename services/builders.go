package services

import (
	"log"

	"github.com/thanxuanvinh/mesage/services/service"
)

var builderRegistry map[string]service.Builder = map[string]service.Builder{}

// ClearBuilders clears registered builders
func ClearBuilders() {
	builderRegistry = map[string]service.Builder{}
}

// RegisterBuilderOrDie registers a service builder or dies.
// This is expected to be called by init functions of builder implementations.
func RegisterBuilderOrDie(name string, builder service.Builder) {
	if _, ok := builderRegistry[name]; ok {
		log.Panicf("service builder name collision [%s]", name)
	}
	if builder == nil {
		log.Panicf("nil service builder [%s]", name)
	}
	builderRegistry[name] = builder
}
