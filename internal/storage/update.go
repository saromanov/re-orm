package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Update provides updating of the data
func Update(client *redis.Client, id, req interface{}) error {
	return update(client, id, req)
}

func update(client *redis.Client, req, rst interface{}) error {
	var resp interface{}
	err := get(client, req, &resp, true)
	if err != nil {
		return errors.Wrap(err, "unable to get value")
	}

	id, ok := resp.(map[string]interface{})
	if !ok {
		return nil
	}
	if err := client.Do("DEL", fmt.Sprintf("id_%v", id)).Err(); err != nil {
		return errors.Wrap(err, "unable to find by the key")
	}

	fields, err := reflect.GetFullFields(rst)
	if err != nil {
		return err
	}

	fmt.Println("FIELDS: ", fields)
	return nil
}
