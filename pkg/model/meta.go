package model

// Meta represents app metadata
type Meta struct {
	Name       string `yaml:"Name"`
	Type       string `yaml:"Type"`
	Lang       string `yaml:"Lang"`
	Main       string `yaml:"Main"`
	Port       int    `yaml:"Port"`
	UsesClient bool   `yaml:"UsesClient"`
	Owner      string `yaml:"Owner"`
}
