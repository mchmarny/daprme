package model

import (
	"regexp"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	// HTTPProtocol indicates HTTP Dapr protocol
	HTTPProtocol = "HTTP"
	// GRPCProtocol indicates gRPC Dapr protocol
	GRPCProtocol = "gRPC"
)

// App represents app state
type App struct {
	Name       string       `yaml:"Name"`
	Protocol   string       `yaml:"Protocol"`
	Port       int          `yaml:"Port"`
	Pubsubs    []*Pubsub    `yaml:"Pubsubs"`
	Bindings   []*Binding   `yaml:"Bindings"`
	Services   []*Service   `yaml:"Services"`
	UsesClient bool         `yaml:"UsesClient"`
	Components []*Component `yaml:"Components"`
}

// Marshal serializes App to YAML
func (a *App) Marshal() ([]byte, error) {
	return yaml.Marshal(a)
}

// String serializes App to YAML string
func (a *App) String() string {
	b, err := a.Marshal()
	if err != nil {
		return "Unable to parse application"
	}
	return string(b)
}

// Unmarshal deserializes bytes into an App
func Unmarshal(in []byte) (*App, error) {
	var a App
	err := yaml.Unmarshal([]byte(in), &a)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling data")
	}
	return &a, nil
}

// ToCodeSafeString removes non-alpha characters
func ToCodeSafeString(val string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	return reg.ReplaceAllString(val, "")
}
