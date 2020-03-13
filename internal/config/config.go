// Package config defines inner representation of the config
package config

// Config defines inner config
type Config struct {
	// KeyPrefix returns prefix where data will be saved
	// for example "prefix_id"
	KeyPrefix string
	SetType   SetType
}