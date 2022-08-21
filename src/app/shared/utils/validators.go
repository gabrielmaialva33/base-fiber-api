package utils

import (
	"base-fiber-api/src/database"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Field       string
	Value       string
	Param       string
}

var validate *validator.Validate

func ValidateStruct(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate = validator.New()

	validate.RegisterValidation("unique-field", Unique)

	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Field = strings.ToLower(err.Field())
			element.Value = err.Value().(string)
			element.Param = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func Unique(fl validator.FieldLevel) bool {
	var count int64

	model := fl.Top().Interface()
	field := underscore(fl.StructFieldName())
	value := fl.Field().String()
	fmt.Fprintf(os.Stdout, "Unique field: %s, value: %s\n", field, value)

	database.DB.Model(model).Where(field+" = ?", value).Count(&count)

	return count == 0
}
