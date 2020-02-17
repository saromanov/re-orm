package reorm

import "github.com/saromanov/re-orm/internal/serializer"

// Config defines configuration for the project
type Config struct {
	Address    string
	Password   string
	Serializer serializer.Serializer
}
