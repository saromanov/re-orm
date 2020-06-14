package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializer(t *testing.T) {
	s := &Serializer{}
	result, err := s.Marshal([]byte("foobar"))
	assert.NoError(t, err)
	var r []byte
	assert.NoError(t, s.Unmarshal(result, &r))
	assert.Equal(t, "foobar", string(r))
}
