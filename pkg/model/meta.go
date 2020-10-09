package model

// Meta represents app metadata
type Meta struct {
	Name       string `yaml:"Name"`
	Protocol   string `yaml:"Protocol"`
	Port       int    `yaml:"Port"`
	UsesClient bool   `yaml:"UsesClient"`
}
