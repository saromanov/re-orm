package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// DeleteByID provides deleting data by id
func DeleteByID(client *redis.Client, ID interface{}) error {
	return delete(client, ID)
}

func delete(client *redis.Client, ID interface{}) error {
	if ID == nil {
		return errors.New("ID argument is nil")
	}
	if err := client.Do("DEL", fmt.Sprintf("id_%v", ID)).Err(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to find by the key %v", ID))
	}

	return nil
}
