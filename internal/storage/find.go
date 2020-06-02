package storage

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/models"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Find provides finding of data by filling data on the fields
func Find(client *redis.Client, d interface{}) ([]interface{}, error) {
	if ok := reflect.IsAvailableForFind(d); !ok {
		return nil, fmt.Errorf("Find: input data have unsupported format")
	}

	fields, err := reflect.GetFullFields(d)
	if err != nil {
		return nil, fmt.Errorf("Find: unable to get fields from provided data: %v", err)
	}

	if len(fields.Fields) == 0 {
		return nil, fmt.Errorf("Find: input data is not provided")
	}
	return find(client, fields, d)
}

// general method for finding data
func find(client *redis.Client, s *models.Search, d interface{}) ([]interface{}, error) {
	result := []interface{}{}
	for _, v := range s.Fields {
		if v == nil {
			continue
		}
		key := strings.ToLower(fmt.Sprintf("%v", v))
		members, err := client.SMembers(key).Result()
		if err != nil {
			return nil, fmt.Errorf("unable to find members: %v", err)
		}

		for _, m := range members {
			dataResp := reflect.MakeStructType(d)
			if err := getByKey(client, m, &dataResp); err != nil {
				return nil, fmt.Errorf("unable to get by the key: %v", err)
			}
			result = append(result, dataResp)
		}
	}
	return result, nil
}
