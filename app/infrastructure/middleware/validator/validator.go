package validator

import (
	util_error "go_sample/app/utility/error"

	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return util_error.NewErrValidationError(err)
	}
	return nil
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
