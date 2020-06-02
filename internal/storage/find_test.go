package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	ID   int
	Name string
}

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

	_, err = Find(client, 123)
	assert.Error(t, err)

	_, err = Find(client, &Animal{})
	assert.Error(t, err)

	_, err = Find(client, &Car{})
	assert.Error(t, err)
}
