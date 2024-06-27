package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidBraces(t *testing.T) {
	assert := assert.New(t)
	// Testing valid strings
	assert.Equal(true, ValidBraces("(){}[]"))
	assert.Equal(true, ValidBraces("({[]})"))

	// Testing invalid strings
	assert.Equal(false, ValidBraces("(}"))
	assert.Equal(false, ValidBraces("({[)]"))
}
