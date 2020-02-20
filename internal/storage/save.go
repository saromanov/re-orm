package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/models"
	"github.com/saromanov/re-orm/internal/reflect"
	"github.com/saromanov/re-orm/internal/serializer/json"
)

// Save provides saving of the object
func Save(client *redis.Client, d interface{}) (string, error) {
	if ok := reflect.IsAvailableForSave(d); !ok {
		return "", fmt.Errorf("save: input valus is a not struct")
	}

	fields, err := reflect.GetFields(d)
	if err != nil {
		return "", fmt.Errorf("unable to get fields from provided data: %v", err)
	}

	if err := save(client, fields, d); err != nil {
		return "", fmt.Errorf("unable to save data: %v")
	}
	return "", nil
}

// save provides saving of the model
func save(client *redis.Client, fields *models.Data, d interface{}) error {
	ser := json.Serializer{}
	key := fmt.Sprintf("id_%v", fields.ID)
	result, err := ser.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to marshal data with id %v", err)
	}
	err = client.Do("SET", key, string(result)).Err()
	if err != nil {
		return fmt.Errorf("unable to set data: %v", err)
	}
	return nil
}
