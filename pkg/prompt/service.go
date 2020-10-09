package prompt

import (
	"github.com/dapr-templates/daprme/pkg/model"
)

// ForService collects service info
func ForService() (*model.Service, error) {
	s := &model.Service{}

	// service name
	s.Name = ForString("Service name: ", "myService")

	return s, nil
}
