package storage

import (
	"fmt"

	"github.com/saromanov/re-orm/internal/reflect"
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
