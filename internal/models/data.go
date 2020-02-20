package models

// Data provides inserting of the values
type Data struct {
	ID     interface{}
	Values map[string]interface{}
}

func NewData() *Data {
	return &Data{
		Values: make(map[string]interface{}),
	}
}
