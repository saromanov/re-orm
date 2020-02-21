package storage

import (
	"github.com/go-redis/redis"
)

// Delete provides deleting data by id
func DeleteByID(client *redis.Client, ID interface{}) error {
	return nil
}
