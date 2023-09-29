package service

// Config struct
type Config struct {
	DefaultMode string `yaml:"defaultMode"`
}

// Service implementation
type Service struct {
	name   string
	config *Config
}

// Name implementation
func (s *Service) Name() string {
	return s.name
}

//  Type implementation
func (s *Service) Type() string {
	return builderName
}
