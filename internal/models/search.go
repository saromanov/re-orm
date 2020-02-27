package models

// Search defines searching of the data by the fields
type Search struct {
	Name   string
	Fields map[string]interface{}
}
