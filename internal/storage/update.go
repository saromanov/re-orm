package storage

import (
	"fmt"
	ref "reflect"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/saromanov/re-orm/internal/reflect"
)

// Update provides updating of the data
func Update(client *redis.Client, query, req interface{}) error {
	return update(client, query, req)
}

func update(client *redis.Client, req, rst interface{}) error {

	if req == nil {
		return fmt.Errorf("request attribute is empty")
	}
	if rst == nil {
		return fmt.Errorf("response attribute is empty")
	}

	if reflect.IsAvailableForSave(req) == reflect.UndefinedSaveType {
		return fmt.Errorf("unable to validate request data")
	}
	if reflect.IsAvailableForSave(rst) == reflect.UndefinedSaveType {
		return fmt.Errorf("unable to validate data for response")
	}
	resp := reflect.MakeStructType(req)
	err := get(client, req, &resp, true)
	if err != nil {
		return errors.Wrap(err, "unable to get value")
	}
	elemData := ref.ValueOf(resp).Elem()
	id := elemData.FieldByName("ID").Interface()
	if err := client.Do("DEL", fmt.Sprintf("id_%v", id)).Err(); err != nil {
		return errors.Wrap(err, "unable to find by the key")
	}

	return getFullFieldsAndSave(client, elemData, rst)
}

func getFullFieldsAndSave(client *redis.Client, elemData ref.Value, rst interface{}) error {
	fields, err := reflect.GetFullFields(rst)
	if err != nil {
		return err
	}
	if len(fields.Fields) == 0 {
		return errors.Wrap(err, "fields is not defined")
	}

	for key, value := range fields.Fields {
		elemData.FieldByName(key).Set(ref.ValueOf(value))
	}

	return saveFullFields(client, elemData)
}

func saveFullFields(client *redis.Client, elemData ref.Value) error {
	if _, err := Save(client, elemData.Interface()); err != nil {
		return errors.Wrap(err, "unable to save new data")
	}
	return nil
}
