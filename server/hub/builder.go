package hub

import (
	"fmt"
	"sync"

	"github.com/thanxuanvinh/mesage/io"
	"github.com/thanxuanvinh/mesage/io/hub"
)

const builderName = "mock"

type builder int

// NewConfig inplementation
func (*builder) NewConfig() interface{} {
	return new(Config)
}

// Build implementation
func (*builder) Build(name string, configInterface interface{}) (hub.Hub, error) {
	config, ok := configInterface.(*Config)
	if !ok {
		return nil, fmt.Errorf("incompatible config")
	}

	return &Hub{
		name:       name,
		config:     config,
		ioRegistry: make(map[string]*IO),
		mu:         sync.Mutex{},
		started:    false,
	}, nil
}

func init() {
	io.RegisterHubBuilderOrDie(builderName, new(builder))
}
