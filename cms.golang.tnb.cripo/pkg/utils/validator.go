package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	// err := validate.Struct(user)
	// if err != nil {
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		fmt.Println(err.Namespace())
	// 		fmt.Println(err.Field())
	// 		fmt.Println(err.StructNamespace())
	// 		fmt.Println(err.StructField())
	// 		fmt.Println(err.Tag())
	// 		fmt.Println(err.ActualTag())
	// 		fmt.Println(err.Kind())
	// 		fmt.Println(err.Type())
	// 		fmt.Println(err.Value())
	// 		fmt.Println(err.Param())
	// 		fmt.Println()
	// 	}
	// 	return
	// }

	return validate
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
