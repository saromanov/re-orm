package reflect

import (
	"fmt"
	"reflect"
)

// IsAvailableForSave provides check if input data is available for save
func IsAvailableForSave(d interface{}) bool {
	return isStruct(d)
}

// GetFields provides getting fields from the struct
func GetFields(d interface{}) {
	if ok := isStruct(d); !ok {
		return
	}

}

// getFields returns name of fields from the structure
func getFields(d interface{}) {
	s := reflect.ValueOf(d).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

// isStruct provides checking if input data is a struct
func isStruct(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Struct
}
