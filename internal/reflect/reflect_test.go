package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	ID    int64
	Name  string `reorm:"index"`
	Color string
}

type CarWithoutIndex struct {
	Name string
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
	assert.Equal(t, 1, len(fields.GetIndexes()))

	m := map[string]interface{}{
		"id":  1,
		"foo": "bar",
	}
	fields, err = GetFields(m)
	assert.NoError(t, err)
	assert.Equal(t, 1, fields.PrimaryKey)
	assert.Equal(t, 2, len(fields.Values))
	fields, err = GetFields(&m)
	assert.NoError(t, err)
	assert.Equal(t, 1, fields.PrimaryKey)
	assert.Equal(t, 2, len(fields.Values))
	_, err = GetFields(&map[string]interface{}{
		"foo": "bar",
	})
	assert.Error(t, err)

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

func TestGetType(t *testing.T) {
	var s string
	s = "foobar"
	assert.Equal(t, "", GetType(s))

	assert.Equal(t, "*Car", GetType(&Car{
		Name: "foobar",
	}))
	assert.Equal(t, "Car", GetType(Car{
		Name: "foorbar",
	}))
}

func TestIsAvailableForFind(t *testing.T) {
	assert.False(t, IsAvailableForFind("a"))
	assert.False(t, IsAvailableForFind(1))
	assert.False(t, IsAvailableForFind(nil))

	var cars []Car
	assert.True(t, IsAvailableForFind(&cars))
}

func TestNoError(t *testing.T) {
	n := noIDError{err: "test"}
	assert.Equal(t, n.Error(), "test")
}

func TestFieldsNonDefaultIndex(t *testing.T) {
	c := &CarWithoutIndex{
		Name: "foobar",
	}

	_, err := GetFields(c)
	assert.Error(t, err)
}
