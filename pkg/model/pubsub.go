package model

// PubSub represents PubSub component
type PubSub struct {
	Name  string `yaml:"Name"`
	Type  string `yaml:"Type"`
	Topic string `yaml:"Topic"`
}

// GetType returns the name of the component type
func (c *PubSub) GetType() string {
	return c.Type
}

// GetName returns the name of the component
func (c *PubSub) GetName() string {
	return c.Name
}

// PubsubComponentTypes lists all supported components
func PubsubComponentTypes() []string {
	return []string{
		"azure.eventhubs",
		"azure.servicebus",
		"gcp.pubsub",
		"hazelcast",
		"kafka",
		"mqtt",
		"natsstreaming",
		"pulsar",
		"rabbitmq",
		"redis",
		"snssqs",
	}
}
