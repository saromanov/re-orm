package storage

import (
	"fmt"
	ref "reflect"
	"strings"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/models"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Find provides finding of data by filling data on the fields
func Find(client *redis.Client, d interface{}, resp interface{}) error {
	if ok := reflect.IsAvailableForSave(d); !ok {
		return fmt.Errorf("Find: input data have unsupported format")
	}

	fields, err := reflect.GetFullFields(d)
	if err != nil {
		return fmt.Errorf("Find: unable to get fields from provided data: %v", err)
	}

	if len(fields.Fields) == 0 {
		return fmt.Errorf("Find: input data is not provided")
	}
	return find(client, fields, d, resp)
}

func find(client *redis.Client, s *models.Search, d interface{}, resp interface{}) error {
	dataResp := reflect.MakeStructType(resp)
	result := ref.MakeSlice(ref.SliceOf(ref.TypeOf(dataResp)), 0, 1)
	for _, v := range s.Fields {
		key := strings.ToLower(fmt.Sprintf("%v", v))
		members, err := client.SMembers(key).Result()
		if err != nil {
			return fmt.Errorf("unable to find members: %v", err)
		}

		for _, m := range members {
			if err := getByKey(client, m, &dataResp); err != nil {
				return fmt.Errorf("unable to get by the key: %v", err)
			}
			result = ref.Append(result, ref.ValueOf(dataResp))
		}
	}
	w := result.Interface()
	if ref.ValueOf(w).Len() == 0 {
		return fmt.Errorf("unable to get response")
	}
	resp = w
	return nil
}
