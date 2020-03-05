package storage

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// Get provides getting data by id
func Get(client *redis.Client, name string, ID interface{}, data interface{}) error {
	return get(client, name, ID, data)
}

func get(client *redis.Client, name string, ID interface{}, data interface{}) error {
	return getByKey(client, fmt.Sprintf("id:%v:%v", name, ID), data)
}

func getByKey(client *redis.Client, name string, data interface{}) error {
	objStr, err := client.Do("GET", name).String()
	if err != nil {
		return errors.Wrap(err, "unable to find by the key")
	}
	b := []byte(objStr)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return errors.Wrap(err, "unable to find by the key")
	}
	return nil
}
