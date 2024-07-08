package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple3And5(t *testing.T) {
	assert := assert.New(t)
	// Test the example with 10 from the original task
	assert.Equal(Multiple3And5(10), 23)
}
