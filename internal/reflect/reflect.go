package reflect

import (
	"fmt"
	"reflect"
	"strings"

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

// GetFullFields provides getting non empty fields from the struct
func GetFullFields(d interface{}) (*models.Search, error) {
	if ok := IsAvailableForSave(d); !ok {
		return nil, fmt.Errorf("unable to save provided data")
	}

	return getFullFields(d), nil
}

// getFields returns name of fields from the structure
func getFields(d interface{}) (*models.Data, error) {
	s := reflect.ValueOf(d).Elem()
	dataType := s.Type()
	resp := models.NewData()
	resp.Name = getName(fmt.Sprintf("%T", d))
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if dataType.Field(i).Name == "ID" {
			resp.ID = f.Interface()
			continue
		}
		tags := dataType.Field(i).Tag.Get("reorm")
		if isStructField(dataType.Field(i)) {

		} else {
			saveKey := dataType.Field(i).Name
			resp.AddValue(saveKey, f.Interface())
			if strings.Contains(tags, "index") {
				resp.AddIndex(fmt.Sprintf("%s:%s:%v", resp.Name, dataType.Field(i).Name, f.Interface()), saveKey)
			}
		}
	}
	if resp.ID == nil {
		return nil, fmt.Errorf("id is not defined")
	}
	return resp, nil
}

// getting "clear" name from the struct
func getName(name string) string {
	if !strings.Contains(name, ".") {
		return name
	}
	splitter := strings.Split(name, ".")
	if len(splitter) == 1 {
		return name
	}
	return splitter[1]
}

// parseTags provides checks of the tags at the input
// and if it contains any tags, its adding to the result
func parseTags(tags string, data *models.Data) {

}

// check if struct contains struct field
func isStructField(sf reflect.StructField) bool {
	t := sf.Type
	return t.Kind() == reflect.Struct ||
		(t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct)
}

// getFullFields retruns filled fields from the input data
func getFullFields(d interface{}) *models.Search {
	s := reflect.ValueOf(d).Elem()
	typeOfT := s.Type()
	resp := &models.Search{}
	resp.Fields = map[string]interface{}{}
	resp.Name = fmt.Sprintf("%T", d)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Interface() != nil && !f.IsZero() {
			resp.Fields[typeOfT.Field(i).Name] = f.Interface()
		}
	}
	return resp
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
