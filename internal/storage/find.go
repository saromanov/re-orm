package storage

import (
	"fmt"
	"strings"

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

	if len(fields.Fields) == 0 {
		return fmt.Errorf("input data is not provided")
	}

	return find(client, fields, d, resp)
}

func find(client *redis.Client, s *models.Search, d interface{}, resp interface{}) error {
	for k, v := range s.Fields {
		fmt.Println("KEY: ", fmt.Sprintf("%v_%v", k, v))
		key := strings.ToLower(fmt.Sprintf("%v_%v", k, v))
		members, err := client.SMembers(key).Result()
		if err != nil {
			return fmt.Errorf("unable to find members: %v", err)
		}
		fmt.Println("MEMBERS: ", members)
	}
	return nil
}
