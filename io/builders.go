package io

import (
	"fmt"
	"log"

	"github.com/thanxuanvinh/mesage/commons"
	"github.com/thanxuanvinh/mesage/io/hub"
)

var hubBuilderRegistry map[string]hub.Builder = map[string]hub.Builder{}

// BuildHub builds an IO hub with the named builder and the provided configParser.
func BuildHub(hubName, builderName string, parseConfig commons.ConfigParser) (hub.Hub, error) {
	if parseConfig == nil {
		return nil, fmt.Errorf("nil config parser for hub [%s]", hubName)
	}

	builder, ok := hubBuilderRegistry[builderName]
	if !ok || builder == nil {
		return nil, fmt.Errorf("non-existent IO hub builder [%s]", builderName)
	}

	config := builder.NewConfig()
	if err := parseConfig(config); err != nil {
		return nil, fmt.Errorf("error parsing hub [%s] config with builder [%s]: %w", hubName, builderName, err)
	}

	hub, err := builder.Build(hubName, config)
	if err != nil {
		return nil, fmt.Errorf("failed to build hub [%s] with builder [%s]: %w", hubName, builderName, err)
	}
	return hub, nil
}

// RegisterHubBuilderOrDie registers a hub builder or dies.
// This is expected to be called by init functions of builder implementations.
func RegisterHubBuilderOrDie(name string, builder hub.Builder) {
	if _, ok := hubBuilderRegistry[name]; ok {
		log.Panicf("IO hub builder name collision [%s]", name)
	}
	if builder == nil {
		log.Panicf("nil IO hub builder [%s]", name)
	}
	hubBuilderRegistry[name] = builder
}
