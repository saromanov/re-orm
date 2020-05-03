package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/saromanov/re-orm/internal/reflect"
)

// DeleteByID provides deleting data by id
func DeleteByID(client *redis.Client, req, ID interface{}) error {
	return delete(client, req, ID)
}

func delete(client *redis.Client, req, ID interface{}) error {
	fmt.Printf("SSS: %v", reflect.GetType(req))
	if ID == nil {
		return errors.New("ID argument is nil")
	}
	if err := client.Do("DEL", fmt.Sprintf("id_%v", ID)).Err(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to find by the key %v", ID))
	}

	return nil
}
