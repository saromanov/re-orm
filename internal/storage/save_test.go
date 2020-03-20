package storage

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

type Animal struct {
	ID    int
	Title string
	Name  string
	Color string
	Type  int
}

func TestSaveBasic(t *testing.T) {
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
}

func TestInvalidSave(t *testing.T) {
	type Invalid struct {
		Title string
	}
	a := &Invalid{
		Title: "Title",
	}

	_, err := Save(client, a)
	assert.Error(t, err)

}
