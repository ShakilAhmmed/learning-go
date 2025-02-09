package helpers

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

var validate = validator.New()

func FormatValidation(structObj interface{}) map[string]string {
	validationErrors := make(map[string]string)
	err := validate.Struct(structObj)
	if err == nil {
		return nil // No validation errors
	}

	structType := reflect.TypeOf(structObj)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem() // Dereference pointer
	}

	for _, err := range err.(validator.ValidationErrors) {
		modelField, _ := structType.FieldByName(err.Field())
		field := modelField.Tag.Get("json")

		// If no JSON tag, fallback to field name
		if field == "" || field == "-" {
			field = err.Field()
		}

		switch err.Tag() {
		case "required":
			validationErrors[field] = field + " is required"
		case "min":
			validationErrors[field] = field + " must be at least " + err.Param() + " characters long"
		case "max":
			validationErrors[field] = field + " must be at most " + err.Param() + " characters long"
		case "email":
			validationErrors[field] = "Invalid email format"
		case "gte":
			validationErrors[field] = field + " must be at least " + err.Param()
		case "lte":
			validationErrors[field] = field + " must be at most " + err.Param()
		default:
			validationErrors[field] = "Invalid value for " + field
		}
	}
	return validationErrors
}
