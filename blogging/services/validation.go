package services

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a struct and returns a map of field errors, or nil if valid.
func ValidateStruct(s interface{}) (map[string]string, error) {
	err := validate.Struct(s)
	if err == nil {
		return nil, nil
	}

	// If it's a validation error, convert it to a map
	if verrs, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, verr := range verrs {
			errors[verr.Field()] = "Failed on the '" + verr.Tag() + "' tag"
		}
		return errors, err
	}

	// Other kinds of errors
	return nil, err
}
