package reflect

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/saromanov/re-orm/internal/models"
)

// SaveType defines input type for saving
type SaveType int

const (
	UndefinedSaveType SaveType = iota + 1
	StructSaveType
	MapSaveType
)

type noIDError struct {
	err string
}

func (e *noIDError) Error() string {
	return e.err
}

// IsAvailableForSave provides check if input data is available for save
func IsAvailableForSave(d interface{}) SaveType {
	if isStruct(d) {
		return StructSaveType
	} else if isMap(d) {
		return MapSaveType
	}
	return UndefinedSaveType
}

// GetFields provides getting fields from the struct
func GetFields(d interface{}) (*models.Data, error) {
	saveType := IsAvailableForSave(d)
	if saveType == UndefinedSaveType {
		return nil, fmt.Errorf("unable to save provided data")
	}

	return getFields(d, saveType)
}

// GetFullFields provides getting non empty fields from the struct
func GetFullFields(d interface{}) (*models.Search, error) {
	saveType := IsAvailableForSave(d)
	if saveType == UndefinedSaveType {
		return nil, fmt.Errorf("unable to save provided data")
	}

	return getFullFields(d), nil
}

// MakeStructType provides making of the struct type for response(find)
func MakeStructType(d interface{}) interface{} {
	return reflect.New(reflect.TypeOf(d).Elem()).Interface()
}

// getFields returns name of fields from the structure
func getFields(d interface{}, st SaveType) (*models.Data, error) {
	s := reflect.ValueOf(d)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
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
				resp.AddIndex(fmt.Sprintf("%v:%v:%v", dataType.Field(i).Name, resp.Name, f.Interface()), saveKey)
			}
		}
	}
	if resp.ID == nil {
		return nil, &noIDError{err: "id is not defined"}
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
	resp.Name = getName(fmt.Sprintf("%T", d))
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Interface() != nil && !f.IsZero() {
			resp.Fields[typeOfT.Field(i).Name] = fmt.Sprintf("%s:%s:%v", strings.ToLower(typeOfT.Field(i).Name), resp.Name, f.Interface())
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
