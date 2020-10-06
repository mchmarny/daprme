package model

// Protocol represents Dapr protocol
type Protocol int

const (
	// HTTP indicates HTTP Dapr protocol
	HTTP Protocol = iota
	// GRPC indicates gRPC Dapr protocol
	GRPC
)

func (s Protocol) String() string {
	return [...]string{"HTTP", "GRPC"}[s]
}
