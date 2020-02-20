package reflect

import (
	"fmt"
	"reflect"

	"github.com/saromanov/re-orm/internal/models"
)

// IsAvailableForSave provides check if input data is available for save
func IsAvailableForSave(d interface{}) bool {
	return isStruct(d)
}

// GetFields provides getting fields from the struct
func GetFields(d interface{}) (*models.Data, error) {
	if ok := IsAvailableForSave(d); !ok {
		return nil, fmt.Errorf("unable to save provided data")
	}

	return getFields(d)
}

// getFields returns name of fields from the structure
func getFields(d interface{}) (*models.Data, error) {
	s := reflect.ValueOf(d).Elem()
	typeOfT := s.Type()
	resp := &models.Data{}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
		if typeOfT.Field(i).Name == "ID" {
			resp.ID = f.Interface()
		}
	}
	if resp.ID == nil {
		return nil, fmt.Errorf("id is not defined")
	}
	return resp, nil
}

// isStruct provides checking if input data is a struct
func isStruct(d interface{}) bool {
	switch reflect.ValueOf(d).Kind() {
	case reflect.Struct:
		return true
	case reflect.Ptr:
		return reflect.ValueOf(d).Type().Elem().Kind() == reflect.Struct
	}
	return false
}
