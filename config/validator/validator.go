package validator

import (
	goValidator "github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *goValidator.Validate
}

func NewStructValidator() *StructValidator {
	validate := goValidator.New()

	return &StructValidator{
		validate: validate,
	}
}

func (structValidator *StructValidator) Validate(out any) error {
	return structValidator.validate.Struct(out)
}
