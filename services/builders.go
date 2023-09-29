package services

import (
	"fmt"
	"log"

	"github.com/thanxuanvinh/mesage/commons"
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

// Build builds a service with named builder and the provided config parser
func Build(serviceName, builderName string, parseConfig commons.ConfigParser) (service.Service, error) {
	if parseConfig == nil {
		return nil, fmt.Errorf("nil config parser for service [%s]", serviceName)
	}

	builder, ok := builderRegistry[builderName]
	if !ok || builder == nil {
		return nil, fmt.Errorf("non-existent service builder [%s]", builderName)
	}

	config := builder.NewConfig()
	if err := parseConfig(config); err != nil {
		return nil, fmt.Errorf("failed to buil service [%s] config with builder [%s]: [%w]", serviceName, builderName, err)
	}

	s, err := builder.Build(serviceName, config)
	if err != nil {
		return nil, fmt.Errorf("failed to buil service [%s] with builder [%s]: [%w]", serviceName, builderName, err)
	}

	return s, nil
}
