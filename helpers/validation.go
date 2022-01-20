package helpers

import validator "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validator.New().Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
