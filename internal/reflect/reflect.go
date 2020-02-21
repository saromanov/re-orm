package reflect

import (
	"fmt"
	"reflect"

	"github.com/saromanov/re-orm/internal/models"
)

// IsAvailableForSave provides check if input data is available for save
func IsAvailableForSave(d interface{}) bool {
	return isStruct(d) || isMap(d)
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
	resp := models.NewData()
	resp.Name = reflect.ValueOf(d).String()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if typeOfT.Field(i).Name == "ID" {
			resp.ID = f.Interface()
		} else {
			resp.Values[typeOfT.Field(i).Name] = f.Interface()
		}
	}
	if resp.ID == nil {
		return nil, fmt.Errorf("id is not defined")
	}
	fmt.Println("RESSS: ", resp.Name)
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

func isMap(d interface{}) bool {
	switch reflect.ValueOf(d).Kind() {
	case reflect.Map:
		return true
	case reflect.Ptr:
		return reflect.ValueOf(d).Type().Elem().Kind() == reflect.Map
	}
	return false
}
