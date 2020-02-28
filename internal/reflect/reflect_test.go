package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	ID    int64
	Name  string
	Colot string
}

func TestIsAvailableForSave(t *testing.T) {
	assert.False(t, IsAvailableForSave(5))
	assert.False(t, IsAvailableForSave("a"))
	assert.False(t, IsAvailableForSave(2.5))
	assert.False(t, IsAvailableForSave([]string{"aaa"}))

	assert.True(t, IsAvailableForSave(map[string]interface{}{"a": "b"}))
	assert.True(t, IsAvailableForSave(&Car{ID: 1, Name: "foobar"}))
}

func TestGetFields(t *testing.T) {
	c := &Car{
		ID:   1,
		Name: "foobar",
	}

	fields, err := GetFields(c)
	assert.NoError(t, err)
	assert.NotNil(t, fields)
	assert.Equal(t, fields.Name, "*reflect.Car")
	assert.Equal(t, fields.ID, c.ID)
	values := fields.Values
	dataName, ok := values["Name"]
	assert.True(t, ok)
	assert.Equal(t, dataName, c.Name)
}

func TestGetFullFields(t *testing.T) {
	c := &Car{
		Name: "foobar",
	}

	fields, err := GetFullFields(c)
	assert.NoError(t, err)
	assert.NotNil(t, fields)
	assert.Equal(t, len(fields.Fields), 1)
}
