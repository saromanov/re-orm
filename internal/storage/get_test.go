package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSimple(t *testing.T) {
	a := &Animal{
		ID:    1,
		Title: "Dog",
		Name:  "Bob",
		Color: "Black",
		Type:  1,
	}
	_, err := Save(client, a)
	assert.NoError(t, err)

	var resp Animal
	assert.NoError(t, Get(client, &Animal{ID: 1}, &resp))
	assert.Equal(t, resp.ID, 1)
	assert.Equal(t, resp.Name, "Bob")
}

func TestGetWIthoutIndex(t *testing.T) {
	a := &Animal{
		ID:   1,
		Name: "Dog",
	}
	_, err := Save(client, a)
	assert.NoError(t, err)

	var resp Animal
	assert.Error(t, Get(client, &Animal{Name: "Dog"}, &resp))
}
