package nodejs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolang(t *testing.T) {
	c := &NodeJs{}
	pc := c.GetProjectConfig()
	assert.NotNil(t, pc)
	assert.NotEmpty(t, pc.Main)
}
