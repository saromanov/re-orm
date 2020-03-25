package models

// Data provides inserting of the values
type Data struct {
	PrimaryKey interface{}
	Name       string
	Values     map[string]interface{}
	Indexes    map[string]string
}

// NewData provides initialization of the data
func NewData() *Data {
	return &Data{
		Values:  make(map[string]interface{}),
		Indexes: make(map[string]string),
	}
}

// AddValue provides adding value to the result
func (d *Data) AddValue(key string, value interface{}) {
	d.Values[key] = value
}

// AddIndex provides adding index to the result
func (d *Data) AddIndex(key, value string) {
	d.Indexes[key] = value
}

func (d *Data) GetIndexes() map[string]string {
	return d.Indexes
}
