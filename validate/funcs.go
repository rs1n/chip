package validate

import (
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

func validatePresense(fl validator.FieldLevel) bool {
	if strings.TrimSpace(fl.Field().String()) == "" {
		return false
	}
	return true
}
