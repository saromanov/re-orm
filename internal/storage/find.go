package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/models"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Find provides finding of data by filling data on the fields
func Find(client *redis.Client, d interface{}, resp interface{}) error {
	if ok := reflect.IsAvailableForSave(d); !ok {
		return fmt.Errorf("Find: input values is able to find")
	}

	fields, err := reflect.GetFullFields(d)
	if err != nil {
		return fmt.Errorf("unable to get fields from provided data: %v", err)
	}

	if len(fields.Values) == 0 {
		return fmt.Errorf("input data is not provided")
	}

	return find(client, fields, d, resp)
}

func find(client *redis.Client, fields *models.Data, d interface{}, resp interface{}) error {
	return nil
}
