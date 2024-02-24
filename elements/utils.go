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

func setField(
	element interface{},
	fieldName string,
	value interface{},
) error {
	val := reflect.ValueOf(element)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fmt.Println(val.Kind())
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("element must be a struct or a pointer to a struct")
	}

	fieldVal := val.FieldByName(fieldName)

	if fieldVal.IsValid() && fieldVal.CanSet() {
		return setValue(fieldVal, value)
	}

	// Assuming BasicElement is a struct field, not an interface or pointer
	basicElVal := val.FieldByName("BasicElement")

	if !basicElVal.IsValid() || !basicElVal.CanSet() {
		return fmt.Errorf("cannot found or set BasicElement on %v", reflect.TypeOf(element))
	}

	fieldVal = basicElVal.FieldByName(fieldName)
	fmt.Println(fieldVal, fieldVal.IsValid(), fieldVal.CanSet())
	if fieldVal.IsValid() && fieldVal.CanSet() {
		return setValue(fieldVal, value)
	}

	return fmt.Errorf("cannot found or set %s", fieldName)
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
	value := reflect.ValueOf(s)
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
