package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	a := &Animal{
		ID:    1,
		Title: "Dog",
		Name:  "Bob",
		Color: "Black",
		Type:  1,
	}
	resp, err := Save(client, a)
	assert.NoError(t, err)
	assert.Equal(t, resp, "")

	assert.NoError(t, Update(client, 1, &Animal{
		Title: "John",
	}))
}
