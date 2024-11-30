package utils_validation

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}, customMessages map[string]string) map[string]string {
	errors := make(map[string]string)
	err := validate.Struct(s)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			key := e.StructField() + "." + e.Tag()
			if msg, exists := customMessages[key]; exists {
				errors[e.StructField()] = msg
			} else {
				errors[e.StructField()] = e.Error()
			}
		}
	}

	return errors
}
