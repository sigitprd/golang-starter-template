package config

import (
	"fmt"
	"reflect"
)

// Validate checks if required fields are filled.
func (c *Config) Validate() error {
	// Walk through the fields of the struct using reflection
	return validateStruct(reflect.ValueOf(c), "")
}

// validateStruct recursively validates a struct and its nested structs.
func validateStruct(value reflect.Value, prefix string) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil
	}

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		// Get the `required` tag
		required := fieldType.Tag.Get("required")
		fieldName := fieldType.Name

		// Check if the field is required and empty
		if required == "true" && isEmptyValue(field) {
			return fmt.Errorf("field '%s%s' is required but not set", prefix, fieldName)
		}

		// Recursively validate nested structs
		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			err := validateStruct(field, prefix+fieldName+".")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// isEmptyValue checks if a value is the zero value for its type.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	default:
		return false
	}
}
