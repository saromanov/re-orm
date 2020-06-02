package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	a := &Animal{
		ID:    1,
		Title: "Dog",
		Name:  "Bob",
		Color: "Black",
		Type:  1,
	}
	_, err := Save(client, a)
	assert.NoError(t, err)

	res, err := Find(client, &Animal{
		Title: "Dog",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
}
