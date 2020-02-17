package reorm

import (
	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/storage"
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

// Save provides saving of the data. Also, it returns stored id
func (r *ReOrm) Save(data interface{}) (string, error) {
	return storage.Save(data)
}
