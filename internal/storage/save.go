package storage

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/models"
	"github.com/saromanov/re-orm/internal/reflect"
	"github.com/saromanov/re-orm/internal/serializer/json"
)

// Save provides saving of the object
func Save(client *redis.Client, d interface{}) (string, error) {
	if ok := reflect.IsAvailableForSave(d); !ok {
		return "", fmt.Errorf("save: input values is a not struct or map")
	}

	fields, err := reflect.GetFields(d)
	if err != nil {
		return "", fmt.Errorf("unable to get fields from provided data: %v", err)
	}

	if err := save(client, fields, d); err != nil {
		return "", fmt.Errorf("unable to save data: %v", err)
	}
	return "", nil
}

// save provides saving of the model
func save(client *redis.Client, fields *models.Data, d interface{}) error {
	ser := json.Serializer{}
	key := fmt.Sprintf("id:%v:%v", fields.Name, fields.ID)
	result, err := ser.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to marshal data with id %v", err)
	}
	err = client.Do("SET", key, string(result)).Err()
	if err != nil {
		return fmt.Errorf("unable to set data: %v", err)
	}
	if err := saveIndexes(client, fields, key); err != nil {
		return fmt.Errorf("unable to create index: %v", err)
	}
	return nil
}

// saveIndexes provides saving of indexes
func saveIndexes(client *redis.Client, fields *models.Data, parentID string) error {
	indexes := fields.GetIndexes()
	if len(indexes) == 0 {
		return nil
	}

	for key := range indexes {
		if err := client.HSet(key, "index", parentID).Err(); err != nil {
			return fmt.Errorf("unable to create index %s: %v", key, err)
		}
		key = strings.ToLower(key)
		if ok, _ := client.SIsMember(key, parentID).Result(); ok {
			continue
		}
		if err := client.SAdd(key, parentID).Err(); err != nil {
			return fmt.Errorf("unable to add index %s %s: %v", key, parentID, err)
		}
	}
	return nil
}
