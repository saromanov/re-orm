package storage

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Get provides getting of the saved data by request
func Get(client *redis.Client, req, data interface{}) error {

	fields, err := reflect.GetFullFields(req)
	if err != nil {
		return fmt.Errorf("Get: unable to get fields from provided data: %v", err)
	}

	if len(fields.Fields) == 0 {
		return fmt.Errorf("Get: input data is not provided")
	}

	id, ok := fields.Fields["ID"]
	if ok {
		return getByKey(client, fmt.Sprintf("%s", id), data)
	}
	return nil
}

func First(client *redis.Client, re1, data interface{}) error {
	fields, err := reflect.GetFullFields(req)
	if err != nil {
		return fmt.Errorf("Get: unable to get fields from provided data: %v", err)
	}

	if len(fields.Fields) == 0 {
		return fmt.Errorf("Get: input data is not provided")
	}

	id, ok := fields.Fields["ID"]
	if ok {
		return getByKey(client, fmt.Sprintf("%s", id), data)
	}
	return nil
}

// GetByID provides getting data by id
func GetByID(client *redis.Client, name string, ID interface{}, data interface{}) error {
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
		return errors.Wrap(err, "unable to unmarshal data")
	}
	return nil
}
