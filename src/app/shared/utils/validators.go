package utils

import (
	"base-fiber-api/src/database"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Field       string `json:"field"`
	Value       string `json:"value"`
	Param       string `json:"param"`
}

var validate *validator.Validate

func ValidateStruct(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate = validator.New()
	_ = validate.RegisterValidation("unique", Unique)

	if err := validate.Struct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var response ErrorResponse
			response.FailedField = err.StructNamespace()
			response.Tag = err.Tag()
			response.Field = underscore(err.Field())
			response.Value = err.Value().(string)
			response.Param = err.Param()
			errors = append(errors, &response)
		}
	}

	return errors
}

func Unique(fl validator.FieldLevel) bool {
	var count int64

	model := fl.Top().Interface()
	field := underscore(fl.StructFieldName())
	value := fl.Field().String()

	database.DB.Model(model).Where(field+" = ?", value).Count(&count)

	return count == 0
}
