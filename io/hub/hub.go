package hub

import (
	"context"

	"github.com/thanxuanvinh/mesage/commons"
)

// IORegistrationInfo for hub IOs
type IORegistrationInfo struct {
	Name     string
	Executor string
	Function string
	Handle   commons.Handle
	Config   interface{}
}

// Hub for station IOs
type Hub interface {
	commons.Configer                         // Configer returns new config object for registering station IO
	Register(info *IORegistrationInfo) error // Register registers a station IO through the interface

	Start() error                 // Start starts the IO interface
	ShutDown(ctx context.Context) // ShutDown shuts down the IO interface

	Name() string // Name of the IO interface
	Type() string // Type (builder name) of the IO interface, e.g., OPCUA
}

// Builder for building IO hubs
type Builder interface {
	commons.Configer
	Build(name string, config interface{}) (Hub, error)
}
