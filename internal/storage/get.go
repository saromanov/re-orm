package storage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/saromanov/re-orm/internal/models"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Get provides getting of the saved data by request
func Get(client *redis.Client, req, data interface{}) error {
	return get(client, req, data, true)
}

// Last provides finding of the last element in the array
func Last(client *redis.Client, req, data interface{}) error {
	return get(client, req, data, false)
}

// GetByID provides getting data by id
func GetByID(client *redis.Client, name string, ID interface{}, data interface{}) error {
	return getByKey(client, fmt.Sprintf("id:%v:%v", name, ID), data)
}

// general method for get value
func get(client *redis.Client, req, data interface{}, asc bool) error {
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
	return getByIndex(client, fields, asc, data)
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

// getByIndex provides getting of value by the index
// NOTE: At this moment, it takes only first attribute
func getByIndex(client *redis.Client, fields *models.Search, asc bool, data interface{}) error {
	for name, value := range fields.Fields {
		valueStr := strings.ToLower(fmt.Sprintf("%v", value))
		members, err := client.SMembers(valueStr).Result()
		if err != nil {
			return fmt.Errorf("unable to get members by the name: %s", name)
		}
		if len(members) == 0 {
			return fmt.Errorf("unable to find members by the name: %s", name)
		}
		member := members[0]
		if !asc {
			member = members[len(members)-1]
		}
		return getByKey(client, member, data)
	}
	return nil
}
