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
	return getByIndex(client, fields, data)
}

// First provides finding of the first element in the array
func First(client *redis.Client, req, data interface{}) error {
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

// getByIndex provides getting of value by the index
func getByIndex(client *redis.Client, fields *models.Search, data interface{}) error {
	for name, value := range fields.Fields {
		valueStr := strings.ToLower(fmt.Sprintf("%v", value))
		fmt.Println("VALUESTR: ", valueStr)
		members, err := client.SMembers(valueStr).Result()
		if err != nil {
			return fmt.Errorf("unable to get members by the name: %s", name)
		}
		if len(members) == 0 {
			return fmt.Errorf("unable to find members by the name: %s", name)
		}
		parentID, err := client.HGet(members[0], "index").Result()
		if err != nil {
			return fmt.Errorf("unable to get parentID: %v", err)
		}
		fmt.Println("PARENTID: ", parentID)
		return getByKey(client, parentID, data)
	}
	return nil
}
