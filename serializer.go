package reorm

import json "github.com/pquerna/ffjson/ffjson"

// Serializer defines interface for serialization of structs
type Serializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

type JSONSerializer struct {
}

func(j *JSONSerializer) Marshal(d interface{}) error {
	return json.Marshal(d)
}

func(j *JSONSerializer) Unmarshal(data []byte, res interface{}) error {
	return json.Unmarshal(data, res)
}

