package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	ID    int64
	Name  string
	Colot string
}

func TestData(t *testing.T) {
	d := NewData()
	d.AddIndex("foo", "bar")
	d.AddIndex("bar", "foo")
	indexes := d.GetIndexes()
	assert.Equal(t, len(indexes), 2)
	assert.Equal(t, indexes["foo"], "bar")
	assert.Equal(t, indexes["bar"], "foo")
}
