package model

// Component represents a Dapr component
type Component struct {
	ComponentType string `yaml:"ComponentType"`
	ComponentName string `yaml:"ComponentName"`
}

// StateComponentTypes lists all supported components
func StateComponentTypes() []string {
	return []string{
		"redis",
		"consul",
		"azure.blobstorage",
		"azure.cosmosdb",
		"azure.tablestorage",
		"etcd",
		"cassandra",
		"memcached",
		"mongodb",
		"zookeeper",
		"gcp.firestore",
		"postgresql",
		"sqlserver",
		"hazelcast",
		"cloudstate.crdt",
		"couchbase",
		"aerospike",
		"rethinkdb",
	}
}

// SecretComponentTypes lists all supported components
func SecretComponentTypes() []string {
	return []string{
		"kubernetes",
		"azure.keyvault",
		"hashicorp.vault",
		"aws.secretmanager",
		"gcp.secretmanager",
		"local.file",
		"local.env",
	}
}
