package model

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
		"aws.dynamodb",
		"aws.kinesis",
		"aws.s3",
		"aws.sns",
		"aws.sqs",
		"azure.blobstorage",
		"azure.cosmosdb",
		"azure.eventgrid",
		"azure.eventhubs",
		"azure.servicebusqueues",
		"azure.signalr",
		"azure.storagequeues",
		"cron",
		"gcp.bucket",
		"gcp.pubsub",
		"http",
		"influx",
		"kafka",
		"mqtt",
		"postgres",
		"rabbitmq",
		"redis",
		"rethinkdb.statechange",
		"twilio.sendgrid",
		"twilio.sms",
		"twitter",
	}
}
