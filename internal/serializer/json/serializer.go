package json

import json "github.com/pquerna/ffjson/ffjson"

// Serializer defines...
type Serializer struct {
}

// Marshal provides serialization of object
func (j *Serializer) Marshal(d interface{}) ([]byte, error) {
	return json.Marshal(d)
}

// Unmarshal provides deserialization of object
func (j *Serializer) Unmarshal(data []byte, res interface{}) error {
	return json.Unmarshal(data, res)
}
