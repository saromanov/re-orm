package json

import json "github.com/pquerna/ffjson/ffjson"

// JSONSerializer defines...
type JSONSerializer struct {
}

// Marshal provides serialization of object
func(j *JSONSerializer) Marshal(d interface{}) ([]byte, error) {
	return json.Marshal(d)
}

// Unmarshal provides deserialization of object
func(j *JSONSerializer) Unmarshal(data []byte, res interface{}) error {
	return json.Unmarshal(data, res)
}