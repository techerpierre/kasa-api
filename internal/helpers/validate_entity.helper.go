package helpers

import "reflect"

func ValidateEntity(data any, requirements []string) bool {
	val := reflect.ValueOf(data)
	typeOfData := val.Type()

	var requirmentsAsAny []any

	for _, requirment := range requirements {
		requirmentsAsAny = append(requirmentsAsAny, requirment)
	}

	for i := 0; i < val.NumField(); i++ {
		fieldName := typeOfData.Field(i).Name
		fieldType := typeOfData.Field(i).Type
		if ArrayContains(requirmentsAsAny, fieldName) && fieldType == reflect.TypeOf("") {
			if val.Field(i).Interface() == "" {
				return false
			}
		}
	}

	return false
}
