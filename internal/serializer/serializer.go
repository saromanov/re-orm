package reorm

// Serializer defines interface for serialization of structs
type Serializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}
