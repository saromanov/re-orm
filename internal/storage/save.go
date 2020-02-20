package storage

import (
	"fmt"

	"github.com/saromanov/re-orm/internal/reflect"
	"github.com/saromanov/re-orm/internal/serializer/json"
)

// Save provides saving of the object
func Save(d interface{}) (string, error) {
	if ok := reflect.IsAvailableForSave(d); !ok {
		return "", fmt.Errorf("save: input valus is a not struct")
	}

	fields, err := reflect.GetFields(d)
	if err != nil {
		return "", fmt.Errorf("unable to get fields from provided data: %v", err)
	}

	fmt.Println("FIELDS: ", fields)
	return "", nil
}

// save provides saving of the model
func save(ID interface{}, d interface{}) error {
	ser := json.Serializer{}
	key := fmt.Sprintf("id_%v", ID)
	result, err := ser.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to marshal data with id %v: %v", d.ID, err)
	}
	err = r.client.Do("SET", key, string(result)).Err()
	if err != nil {
		return fmt.Errorf("unable to set data: %v", err)
	}
	return nil
}
