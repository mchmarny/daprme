package model

// App represents app state
type App struct {
	Name     string
	Protocol Protocol
	Port     int
	Pubsubs  []*Pubsub
	Bindings []*Binding
	Services []*Service
}
