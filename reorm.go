package reorm

import (
	"github.com/go-redis/redis"
)

// ReOrm provides implementation of the Redis ORM
type ReOrm struct {
	client *redis.Client
}

// New initialize Redis Orm
func New(c *Config) *ReOrm {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       0,
	})
	return &ReOrm{
		client: client,
	}
}
