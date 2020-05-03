package reorm

import "github.com/saromanov/re-orm/internal/serializer"

// SetType defines type for the set for index insering
type SetType int

var (
	// ZADD provides definition of the type for inserting
	ZADD SetType = 0
	// SADD provides definition of the type for inserting
	SADD SetType = 1
)

// Config defines configuration for the project
type Config struct {
	Address    string
	Password   string
	Serializer serializer.Serializer
	// KeyPrefix returns prefix where data will be saved
	// for example "prefix_id"
	KeyPrefix string
	SetType   SetType
}
