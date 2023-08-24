package xvalidator

import (
	"iwexlmsapi/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator(singleInstance *validator.Validate) {
	validate = singleInstance
}

func ValidateStruct(data interface{}) []models.ErrorResponse {
	validationErrors := []models.ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			elem := models.ErrorResponse{}

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
