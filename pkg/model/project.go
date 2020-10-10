package model

// Project represents app project
type Project struct {
	Main      string
	PortHTTP  int
	PortGRPC  int
	Templates map[string]string
}
