package io

import (
	"fmt"
	"io"
	"log"

	"github.com/thanxuanvinh/mesage/commons"
	"github.com/thanxuanvinh/mesage/io/hub"
	"gopkg.in/yaml.v2"
)

// Manager manages the collection of station IO hubs
type Manager struct {
	hubs map[string]hub.Hub
}

func buildHub(configReader io.Reader) (hub.Hub, error) {
	if configReader == nil {
		return nil, fmt.Errorf("nil hub config Reader")
	}

	//Parse config from reader
	var metadata commons.ConfigMetadata

	dec := yaml.NewDecoder(configReader)
	dec.SetStrict(true)
	if err := dec.Decode(&metadata); err != nil {
		return nil, fmt.Errorf("failed to parse IO hub config %w", err)
	}

	//Build the hub
	log.Printf("building io hub [%s] with type[%s]\n", metadata.Name, metadata.Type)
	h, err := BuildHub(metadata.Name, metadata.Type, dec.Decode)
	if err != nil {
		return nil, fmt.Errorf("failed to build IO hub: %w", err)
	}

	// Expect EOF - one hub config per reader
	if err := dec.Decode(&metadata); err != io.EOF {
		return nil, fmt.Errorf("unexpected config input, expecting EOF")
	}

	return h, nil
}

// NewManager returns a new hub manager
func NewManager(configReaders <-chan io.Reader) (*Manager, error) {
	if configReaders == nil {
		return nil, fmt.Errorf("nil hub config reader channel")
	}

	man := &Manager{
		hubs: make(map[string]hub.Hub),
	}
	for r := range configReaders {
		h, err := buildHub(r)
		if err != nil {
			return nil, err
		}

		name := h.Name()
		if _, ok := man.hubs[name]; ok {
			return nil, fmt.Errorf("IO hub name collision [%s]", name)
		}
		man.hubs[name] = h
	}
	return man, nil
}
