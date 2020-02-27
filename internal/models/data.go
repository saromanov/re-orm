package models

// Data provides inserting of the values
type Data struct {
	ID      interface{}
	Name    string
	Values  map[string]interface{}
	Indexes []string
}

func NewData() *Data {
	return &Data{
		Values:  make(map[string]interface{}),
		Indexes: []string{},
	}
}

// AddValue provides adding value to the result
func (d *Data) AddValue(key string, value interface{}) {
	d.Values[key] = value
}

// AddIndex provides adding index to the result
func (d *Data) AddIndex(key string) {
	d.Indexes = append(d.Indexes, key)
}

func (d *Data) GetIndexes() []string {
	return d.Indexes
}
