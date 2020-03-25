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

type AnimalExtend struct {
	ID    string
	Title string
	Name  string `reorm:"index"`
	Color string
	Type  int
	Sound Sound
}

type Sound struct {
	Message string
}

type Music struct {
	ID int
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

	a2 := &AnimalExtend{
		ID:    "aaa",
		Title: "Dog",
		Name:  "Bob",
		Color: "Black",
		Type:  1,
		Sound: Sound{
			Message: "Data",
		},
	}
	resp, err = Save(client, a2)
	assert.NoError(t, err)
	assert.Equal(t, resp, "")

	resp, err = Save(client, map[string]interface{}{
		"id":   2,
		"name": "bob",
	})
	assert.NoError(t, err)
	assert.Equal(t, resp, "")

	var respData Animal
	assert.NoError(t, Get(client, &Animal{ID: 1}, &respData))
	assert.Equal(t, respData.ID, 1)
	assert.Equal(t, respData.Name, "Bob")

	var respData2 AnimalExtend
	assert.NotEqual(t, respData2.ID, "")
	assert.NoError(t, Get(client, &AnimalExtend{Name: "Bob"}, &respData2))
	assert.Equal(t, respData2.Name, "Bob")
	assert.Equal(t, respData2.Sound.Message, "Data")
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
	_, err = Save(client, 1)
	assert.EqualError(t, err, errNotAvailableForSave.Error())
	_, err = Save(client, "A")
	assert.EqualError(t, err, errNotAvailableForSave.Error())
	_, err = Save(client, []string{"a", "b"})
	assert.EqualError(t, err, errNotAvailableForSave.Error())
}
