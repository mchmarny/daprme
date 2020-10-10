package prompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeSafeStringParsing(t *testing.T) {
	s := codeSafeString("foo.bar-Something")
	assert.NotEmpty(t, s)
	assert.Equal(t, "foobarSomething", s)
}
