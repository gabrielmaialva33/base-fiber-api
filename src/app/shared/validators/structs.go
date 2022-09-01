package validators

import (
	"base-fiber-api/src/app/shared/utils"
	"base-fiber-api/src/database"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Field       string `json:"field"`
	Value       string `json:"value"`
	Param       string `json:"param"`
}

var validate *validator.Validate

// ValidateStruct validates a struct (all the fields)
func ValidateStruct(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate = validator.New()
	_ = validate.RegisterValidation("unique", Unique)

	if err := validate.Struct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var response ErrorResponse
			response.FailedField = err.StructNamespace()
			response.Tag = err.Tag()
			response.Field = utils.Underscore(err.Field())
			response.Value = err.Value().(string)
			response.Param = err.Param()
			errors = append(errors, &response)
		}
	}

	return errors
}

// ValidatePartialStruct validates a partial struct (only the fields that are present)
func ValidatePartialStruct(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate = validator.New()
	_ = validate.RegisterValidation("unique", Unique)

	var fields []string
	val := reflect.Indirect(reflect.ValueOf(model))
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Interface() != "" {
			fields = append(fields, val.Type().Field(i).Name)
		}
	}

	if err := validate.StructPartial(model, fields...); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var response ErrorResponse
			response.FailedField = err.StructNamespace()
			response.Tag = err.Tag()
			response.Field = utils.Underscore(err.Field())
			response.Value = err.Value().(string)
			response.Param = err.Param()
			errors = append(errors, &response)
		}
	}

	return errors
}

// Unique checks if a field is unique in the database
func Unique(fl validator.FieldLevel) bool {
	var count int64

	model := fl.Top().Interface()
	field := utils.Underscore(fl.StructFieldName())
	value := fl.Field().String()

	database.DB.Model(model).Where(field+" = ?", value).Count(&count)

	return count == 0
}