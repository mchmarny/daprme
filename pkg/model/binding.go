package model

// Binding represents Binding component
type Binding struct {
	Component
}

// InputBindingComponentTypes lists all supported components
func InputBindingComponentTypes() []string {
	return []string{
		"aws.sqs",
		"aws.kinesis",
		"azure.eventhubs",
		"kafka",
		"mqtt",
		"rabbitmq",
		"azure.servicebusqueues",
		"azure.storagequeues",
		"gcp.pubsub",
		"kubernetes",
		"azure.eventgrid",
		"twitter",
		"cron",
		"rethinkdb.statechange",
	}
}

// OutputBindingComponentTypes lists all supported components
func OutputBindingComponentTypes() []string {
	return []string{
		"aws.sqs",
		"aws.sns",
		"aws.kinesis",
		"azure.eventhubs",
		"aws.dynamodb",
		"azure.cosmosdb",
		"gcp.bucket",
		"http",
		"kafka",
		"mqtt",
		"rabbitmq",
		"redis",
		"aws.s3",
		"azure.blobstorage",
		"azure.servicebusqueues",
		"azure.storagequeues",
		"gcp.pubsub",
		"azure.signalr",
		"twilio.sms",
		"twilio.sendgrid",
		"azure.eventgrid",
		"cron",
		"twitter",
		"influx",
		"postgres",
	}
}
