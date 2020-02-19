package reflect

import "reflect"

// IsAvailableForSave provides check if input data is available for save
func IsAvailableForSave(d interface{}) bool {
	return isStruct(d)
}

// isStruct provides checking if input data is a struct
func isStruct(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Struct
}
