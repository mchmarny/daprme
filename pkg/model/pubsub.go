package model

// Pubsub represents PubSub component
type Pubsub struct {
	ComponentType string `yaml:"ComponentType"`
	ComponentName string `yaml:"ComponentName"`
	TopicName     string `yaml:"TopicName"`
}

// PubsubComponentTypes lists all supported components
func PubsubComponentTypes() []string {
	return []string{
		"redis",
		"natsstreaming",
		"azure.eventhubs",
		"azure.servicebus",
		"rabbitmq",
		"hazelcast",
		"gcp.pubsub",
		"kafka",
		"snssqs",
		"mqtt",
		"pulsar",
	}
}
