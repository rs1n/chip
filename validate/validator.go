package validate

import (
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

type (
	IValidator interface {
		Validate() ValidationErrors
	}

	Validate struct {
		*validator.Validate
		TranslateError
	}
)

func NewValidate(translateError TranslateError) *Validate {
	// Create a base validator and use json name to represent error's namespace.
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	if translateError == nil {
		translateError = GetErrorMessageFor // Default translator.
	}
	result := &Validate{
		Validate:       validate,
		TranslateError: translateError,
	}
	return result
}

func (v *Validate) ValidateStruct(s interface{}) error {
	err := v.Validate.Struct(s)
	if err != nil {
		return v.getValidationErrors(err)
	}

	customValidator, ok := s.(IValidator)
	if ok {
		customErrors := customValidator.Validate()
		if customErrors != nil {
			return customErrors
		}
	}
	return nil
}

func (v *Validate) getValidationErrors(err error) ValidationErrors {
	validationErrors := ValidationErrors{}
	for _, err := range err.(validator.ValidationErrors) {
		namespace := err.Namespace()
		tag := err.Tag()
		param := err.Param()
		validationErrors[namespace] = v.TranslateError(tag, param)
	}
	return validationErrors
}
