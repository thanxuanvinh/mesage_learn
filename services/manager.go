package services

import (
	"fmt"
	"io"
	"log"

	"github.com/thanxuanvinh/mesage/commons"
	"github.com/thanxuanvinh/mesage/services/service"
	"gopkg.in/yaml.v2"
)

// Manager manages the collection of services
type Manager struct {
	services map[string]service.Service
}

func buildService(configReader io.Reader) (service.Service, error) {
	if configReader == nil {
		return nil, fmt.Errorf("nil service config Reader")
	}

	// Parse config from reader
	var metadata commons.ConfigMetadata
	dec := yaml.NewDecoder(configReader)
	dec.SetStrict(true)
	if err := dec.Decode(&metadata); err != nil {
		return nil, fmt.Errorf("failed to parse service config metadata: %w", err)
	}

	// Build the service
	log.Printf("building service [%s] with type [%s]", metadata.Name, metadata.Type)
	s, err := Build(metadata.Name, metadata.Type, dec.Decode)
	if err != nil {
		return nil, fmt.Errorf("failed to build service: [%w]", err)
	}

	// Expect EOF - one service config per reader
	if err := dec.Decode(&metadata); err != io.EOF {
		return nil, fmt.Errorf("unexpected config input, expecting EOF")
	}

	return s, nil
}

// NewManager returns a new service manager
func NewManager(configReader <-chan io.Reader) (*Manager, error) {
	if configReader == nil {
		return nil, fmt.Errorf("nil service config reader channel")
	}

	man := &Manager{
		services: make(map[string]service.Service),
	}
	for r := range configReader {
		s, err := buildService(r)
		if err != nil {
			return nil, err
		}

		name := s.Name()
		if _, ok := man.services[name]; ok {
			return nil, fmt.Errorf("service name collision [%s]", name)
		}

		man.services[name] = s
	}

	return man, nil

}
