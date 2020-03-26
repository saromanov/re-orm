package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	ID    int64
	Name  string
	Color string
}

func TestIsAvailableForSave(t *testing.T) {
	assert.Equal(t, IsAvailableForSave(5), UndefinedSaveType)
	assert.Equal(t, IsAvailableForSave("a"), UndefinedSaveType)
	assert.Equal(t, IsAvailableForSave(2.5), UndefinedSaveType)
	assert.Equal(t, IsAvailableForSave([]string{"aaa"}), UndefinedSaveType)

	assert.Equal(t, IsAvailableForSave(map[string]interface{}{"a": "b"}), MapSaveType)
	assert.Equal(t, IsAvailableForSave(&Car{ID: 1, Name: "foobar"}), StructSaveType)
	assert.Equal(t, IsAvailableForSave(Car{ID: 1, Name: "foobar"}), StructSaveType)

	m := map[string]interface{}{"a": "b"}
	assert.Equal(t, IsAvailableForSave(&m), MapSaveType)
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
	assert.Equal(t, fields.PrimaryKey, c.ID)
	values := fields.Values
	dataName, ok := values["Name"]
	assert.True(t, ok)
	assert.Equal(t, dataName, c.Name)

	m := map[string]interface{}{
		"id":  1,
		"foo": "bar",
	}
	fields, err = GetFields(m)
	assert.NoError(t, err)
	assert.Equal(t, 1, fields.PrimaryKey)
	assert.Equal(t, 2, len(fields.Values))

}

func TestGetUnsupportedFields(t *testing.T) {
	fields, err := GetFields(10)
	assert.Error(t, err, errUnsupportedType)
	assert.Nil(t, fields, nil)
}

func TestGetFullFields(t *testing.T) {
	c := &Car{
		Name: "foobar",
	}

	fields, err := GetFullFields(c)
	assert.NoError(t, err)
	assert.NotNil(t, fields)
	assert.Equal(t, len(fields.Fields), 1)

	_, err = GetFullFields(1)
	assert.Error(t, err)
}

func TestGetMapFullFields(t *testing.T) {
	d := map[string]interface{}{
		"id":   1,
		"name": "foobar",
	}
	fields, err := GetFullFields(d)
	assert.NoError(t, err)
	assert.NotNil(t, fields)
}

func TestMakeStructType(t *testing.T) {
	assert.Equal(t, &Car{}, MakeStructType(&Car{}))
}
