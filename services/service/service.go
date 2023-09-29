package service

import "github.com/thanxuanvinh/mesage/commons"

// Service for executors
type Service interface {
	Name() string // Name of the service
	Type() string // Type (builder name) of the service
}

// Builder for building services
type Builder interface {
	commons.Configer
	Build(name string, config interface{}) (Service, error)
}
