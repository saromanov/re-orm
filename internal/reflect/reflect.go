package reflect

import "reflect"

// IsStruct provides checking if input data is a struct
func IsStruct(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Struct
}
