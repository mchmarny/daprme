package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLangs(t *testing.T) {
	for _, l := range GetLangs() {
		c, err := MakeConfigurable(l)
		assert.NoError(t, err)
		assert.NotNil(t, c)
		pc := c.GetProjectConfig()
		assert.NotNil(t, pc)
		assert.NotEmpty(t, pc.Main)
	}
}
