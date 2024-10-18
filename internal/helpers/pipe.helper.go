package helpers

import (
	"fmt"
	"reflect"
)

func Pipe(input any, output any) error {
	inputVal := reflect.ValueOf(input)
	outputVal := reflect.ValueOf(output)

	if inputVal.Kind() != reflect.Ptr || outputVal.Kind() != reflect.Ptr {
		return fmt.Errorf("les paramètres doivent être des pointeurs vers des structures")
	}

	inputVal = inputVal.Elem()
	outputVal = outputVal.Elem()

	if inputVal.Kind() != reflect.Struct || outputVal.Kind() != reflect.Struct {
		return fmt.Errorf("les paramètres doivent être des structures")
	}

	for i := 0; i < inputVal.NumField(); i++ {
		fieldName := inputVal.Type().Field(i).Name

		outField := outputVal.FieldByName(fieldName)

		if outField.IsValid() && outField.CanSet() {
			inField := inputVal.Field(i)

			if inField.Kind() == reflect.Struct && outField.Kind() == reflect.Struct {
				err := Pipe(inField.Addr().Interface(), outField.Addr().Interface())
				if err != nil {
					return err
				}
			} else if inField.Type() == outField.Type() {
				outField.Set(inField)
			}
		}
	}

	return nil
}
