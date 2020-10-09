package model

// Meta represents app metadata
type Meta struct {
	Name       string `yaml:"Name"`
	Type       string `yaml:"Type"`
	Port       int    `yaml:"Port"`
	UsesClient bool   `yaml:"UsesClient"`
}
