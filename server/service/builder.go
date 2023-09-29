package service

import (
	"fmt"

	"github.com/thanxuanvinh/mesage/services"
	"github.com/thanxuanvinh/mesage/services/service"
)

const builderName = "mock"

type builder int

// NewConfig implementation
func (b *builder) NewConfig() interface{} {
	return new(Config)
}

// Build implementation
func (b *builder) Build(name string, configInterface interface{}) (service.Service, error) {
	config, ok := configInterface.(*Config)
	if !ok {
		return nil, fmt.Errorf("incompatible config")
	}

	return &Service{
		name:   name,
		config: config,
	}, nil
}

func init() {
	services.RegisterBuilderOrDie(builderName, new(builder))
}
