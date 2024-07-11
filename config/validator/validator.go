package validator

import (
	"regexp"

	goValidator "github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *goValidator.Validate
}

func NewStructValidator() *StructValidator {
	validate := goValidator.New()
	validate.RegisterValidation("password", func(fl goValidator.FieldLevel) bool {
		input := fl.Field().String()
		if len(input) < 8 {
			return false
		}

		if !regexp.MustCompile(`^[a-zA-Z!@#$%^&*()]$`).MatchString(string(input[0])) {
			return false
		}

		if !regexp.MustCompile(`\d`).MatchString(input) || !regexp.MustCompile(`[a-zA-Z]`).MatchString(input) || !regexp.MustCompile(`[!@#$%^&*()$]`).MatchString(input) {
			return false
		}

		return true
	})

	return &StructValidator{
		validate: validate,
	}
}

func (structValidator *StructValidator) Validate(out any) error {
	return structValidator.validate.Struct(out)
}
