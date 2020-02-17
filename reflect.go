package reorm

import "reflect"

func isStruct(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Struct
}
