package writer

import (
	"testing"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestAddComponent(t *testing.T) {
	p := &model.PubSub{}
	p.Name = "awssqs-binding"
	p.Type = "bindings.aws.sqs"

	var i interface{}
	i = p

	c, ok := i.(model.Componentable)
	assert.True(t, ok)
	assert.NotNil(t, c)

	compName := c.GetName()
	assert.NotEmpty(t, compName)

	compType := c.GetType()
	assert.NotEmpty(t, compType)
}

func TestFileExists(t *testing.T) {
	ok := pathExists("./project_test.go")
	assert.True(t, ok)
}
