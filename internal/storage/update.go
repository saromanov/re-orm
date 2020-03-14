package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// Update provides updating of the data
func Update(client *redis.Client, id, req interface{}) error {
	return update(client, id, req)
}

func update(client *redis.Client, id, req interface{}) error {
	if err := client.Do("DEL", fmt.Sprintf("id_%v", id)).Err(); err != nil {
		return errors.Wrap(err, "unable to find by the key")
	}

	return nil
}
