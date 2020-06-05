package unique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	assert.NotEqual(t, 0, len(GenerateID()))
}
