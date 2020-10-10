package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T) {
	err := Exec(".", "echo", "test")
	assert.NoError(t, err)
}
