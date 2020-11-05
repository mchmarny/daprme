package prompt

import (
	"github.com/dapr-templates/daprme/pkg/model"
)

// ForService collects service info
func ForService() *model.Service {
	s := &model.Service{}
	s.Name = ForString("Operation name: ", "myMethod")
	return s
}
