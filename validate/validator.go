package validate

import (
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

type Validate struct {
	*validator.Validate
}

func NewValidate() *Validate {
	// Create a base validator and use json name to represent error's namespace.
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	result := &Validate{
		Validate: validate,
	}
	return result
}

func (v *Validate) Struct(s interface{}) error {
	err := v.Validate.Struct(s)
	if err != nil {
		return getValidationErrors(err)
	}
	return nil
}

func getValidationErrors(err error) error {
	validationErrors := ValidationErrors{}
	for _, err := range err.(validator.ValidationErrors) {
		namespace := err.Namespace()
		tag := err.Tag()
		param := err.Param()
		validationErrors[namespace] = GetErrorMessageFor(tag, param)
	}
	return validationErrors
}
