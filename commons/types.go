package commons

import "context"

// Configer for hubs, IOs, and services
type Configer interface {
	NewConfig() interface{}
}

// ConfigParse for parsing hub, IO, and service configs
type ConfigParser func(config interface{}) error

//ConfigMetadata for hub, IO, and service configs
type ConfigMetadata struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// Values for hub IO input and output
type Values map[string]interface{}

// Handle function for the hub IO
type Handle func(context.Context, Values) (Values, error)
