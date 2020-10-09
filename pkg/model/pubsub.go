package model

// Pubsub represents PubSub component
type Pubsub struct {
	Component
	TopicName string `yaml:"TopicName"`
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
