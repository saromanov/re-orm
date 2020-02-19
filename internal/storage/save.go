package storage

import (
	"fmt"

	"github.com/saromanov/re-orm/internal/reflect"
)

// Save provides saving of the object
func Save(data interface{}) (string, error) {
	if ok := reflect.IsAvailableForSave(data); !ok {
		return "", fmt.Errorf("save: input valus is a not struct")
	}
	return "", nil
}
