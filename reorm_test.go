package reorm

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
	Title string `reorm:"index"`
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

	_, err := New(nil)
	assert.Error(t, err)

	r, err := New(&Config{
		Address: "localhost:6379",
	})
	assert.NoError(t, r.Save(a))
}
