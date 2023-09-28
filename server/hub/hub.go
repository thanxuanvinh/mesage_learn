package hub

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/thanxuanvinh/mesage/commons"
	"github.com/thanxuanvinh/mesage/io/hub"
)

// IOConfig for hub IOs
type IOConfig struct {
	InputTransformation  map[string]string `yaml:"ioInputTransformation"`
	OutputTransformation map[string]string `yaml:"ioOutputTransformation"`
}

// IO definition
type IO struct {
	handle               commons.Handle
	inputTransformation  map[string]string
	outputTransformation map[string]string
}

// Config for hub
type Config struct {
	Color string `yaml:"color"`
	Taste string `yaml:"taste"`
}

func (c *Config) get(key string) interface{} {
	switch key {
	case "color":
		return c.Color
	case "taste":
		return c.Taste
	default:
		return ""
	}
}

// Hub implementation
type Hub struct {
	name       string
	config     *Config
	ioRegistry map[string]*IO

	mu      sync.Mutex
	started bool
	closed  bool
	wg      sync.WaitGroup
}

// NewConfig implementation for IO registeration
func (h *Hub) NewConfig() interface{} {
	return new(IOConfig)
}

// Register implementation
func (h *Hub) Register(info *hub.IORegistrationInfo) error {
	// Start should be called after all Register calls
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.started {
		return fmt.Errorf("register is disallowed after the hub is started")
	}

	name := fmt.Sprintf("%s|%s", info.Executor, info.Name)
	if _, ok := h.ioRegistry[name]; ok {
		return fmt.Errorf("io name collision: %s", name)
	}

	config, ok := info.Config.(*IOConfig)
	if !ok {
		return fmt.Errorf("imcompatible io config")
	}

	log.Printf("new io %s registered to hub %s", name, h.name)
	h.ioRegistry[name] = &IO{
		info.Handle,
		config.InputTransformation,
		config.OutputTransformation}
	return nil
}

// Start implementation
func (h *Hub) Start() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.started {
		return fmt.Errorf("hub already started")
	}
	h.started = true
	return nil
}

// ShutDown implementation
func (h *Hub) ShutDown(context.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.closed = true

	h.wg.Wait() // Wait until all ongoing IOs are completed
}

// Name implementaion
func (h *Hub) Name() string {
	return h.name
}

// Type implementation
func (*Hub) Type() string {
	return builderName
}
