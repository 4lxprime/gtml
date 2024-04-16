package elements

import (
	"fmt"
	"reflect"
)

func hasField(
	s interface{},
	fieldName string,
) bool {
	value := reflect.ValueOf(s)
	field := value.FieldByName(fieldName)
	return field.IsValid()
}

func setFieldValue(
	element interface{},
	fieldName string,
	value interface{},
) error {
	val := reflect.ValueOf(element).Elem()

	field := val.FieldByName(fieldName)
	basicField := val.FieldByName("BasicElement")

	if !field.IsValid() {
		if basicField.IsValid() && basicField.CanSet() {
			basicFieldName := basicField.FieldByName(fieldName)
			if basicFieldName.IsValid() && basicFieldName.CanSet() {
				return setValue(basicField, value)
			}
		}

		customField := val.FieldByName("Custom")
		if customField.IsValid() && customField.CanSet() {
			customFieldName := customField.FieldByName(fieldName)
			if customFieldName.IsValid() && customFieldName.CanSet() {
				return setValue(basicField, value)
			}
		}

		return fmt.Errorf("field %s doesn't exist", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot edit field %s", fieldName)
	}

	return setValue(field, value)
}

func setValue(
	fieldVal reflect.Value,
	value interface{},
) error {
	switch fieldVal.Kind() {
	case reflect.String:
		if v, ok := value.(string); ok {
			fieldVal.SetString(v)
			return nil
		}

	case reflect.Bool:
		if v, ok := value.(bool); ok {
			fieldVal.SetBool(v)
			return nil
		}

	case reflect.Int64:
		if v, ok := value.(int64); ok {
			fieldVal.SetInt(v)
			return nil
		}

	case reflect.Func:
		if v, ok := value.(EventHandler); ok {
			fieldVal.Set(reflect.ValueOf(v))
			return nil
		}

	default:
		return fmt.Errorf("wrong value type for the field")
	}

	return fmt.Errorf("unsupported field type")
}

func fieldsToMap(s interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	value := reflect.ValueOf(s).Elem()
	t := value.Type()

	var addFields func(reflect.Value, reflect.Type)
	addFields = func(v reflect.Value, t reflect.Type) {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" {
				continue
			}

			fieldValue := v.Field(i)
			key := field.Name

			if fieldValue.IsZero() {
				continue
			}

			if fieldValue.Kind() == reflect.Struct {
				addFields(fieldValue, fieldValue.Type())

			} else {
				fields[key] = fieldValue.Interface()
			}
		}
	}

	addFields(value, t)
	return fields
}
