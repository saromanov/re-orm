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
	assert.Equal(t, IsAvailableForSave(5), UndefinedSaveType)
	assert.Equal(t, IsAvailableForSave("a"), UndefinedSaveType)
	assert.Equal(t, IsAvailableForSave(2.5), UndefinedSaveType)
	assert.Equal(t, IsAvailableForSave([]string{"aaa"}), UndefinedSaveType)

	assert.Equal(t, IsAvailableForSave(map[string]interface{}{"a": "b"}), MapSaveType)
	assert.Equal(t, IsAvailableForSave(&Car{ID: 1, Name: "foobar"}), StructSaveType)
}

func TestGetFields(t *testing.T) {
	c := &Car{
		ID:   1,
		Name: "foobar",
	}

	fields, err := GetFields(c)
	assert.NoError(t, err)
	assert.NotNil(t, fields)
	assert.Equal(t, fields.Name, "Car")
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
