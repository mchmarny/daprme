package prompt

import (
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// ForService collects service info
func ForService() (*model.Service, error) {
	s := &model.Service{}

	// service name
	name, err := ForString("Service name: ", "myService")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	s.Name = name

	return s, nil
}
