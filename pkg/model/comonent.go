package model

// Component represents a Dapr component
type Component struct {
	ComponentType string `yaml:"ComponentType"`
	ComponentName string `yaml:"ComponentName"`
}

// GetType returns the name of the component type
func (c *Component) GetType() string {
	return c.ComponentType
}

// GetName returns the name of the component
func (c *Component) GetName() string {
	return c.ComponentName
}

// Componentable defines the component interface
type Componentable interface {
	GetType() string
	GetName() string
}

// StateComponentTypes lists all supported components
func StateComponentTypes() []string {
	return []string{
		"aerospike",
		"azure.blobstorage",
		"azure.cosmosdb",
		"azure.tablestorage",
		"cassandra",
		"cloudstate.crdt",
		"consul",
		"couchbase",
		"etcd",
		"gcp.firestore",
		"hazelcast",
		"memcached",
		"mongodb",
		"postgresql",
		"redis",
		"rethinkdb",
		"sqlserver",
		"zookeeper",
	}
}

// SecretComponentTypes lists all supported components
func SecretComponentTypes() []string {
	return []string{
		"aws.secretmanager",
		"azure.keyvault",
		"gcp.secretmanager",
		"hashicorp.vault",
		"kubernetes",
		"local.env",
		"local.file",
	}
}
