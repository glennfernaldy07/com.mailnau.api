package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(req interface{}) error {
	if err := validate.Struct(req); err != nil {
		return formatValidationError(err.(validator.ValidationErrors))
	}
	return nil
}

func formatValidationError(errors validator.ValidationErrors) error {
	for _, err := range errors {
		switch err.Tag() {
		case "required":
			return fmt.Errorf("properti '%s' diperlukan", err.Field())
		case "oneof":
			return fmt.Errorf("properti '%s' wajib salah satu dari [%s]", err.Field(), err.Param())
		case "max":
			return fmt.Errorf("properti '%s' melebihi nilai maksimal (%s)", err.Field(), err.Param())
		default:
			return fmt.Errorf("properti '%s' tidak valid", err.Field())
		}
	}
	return fmt.Errorf("kesalahan validasi permintaan")
}
