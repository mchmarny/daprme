package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolang(t *testing.T) {
	c := &Golang{}
	pc := c.GetProjectConfig()
	assert.NotNil(t, pc)
	assert.NotEmpty(t, pc.Main)
}
