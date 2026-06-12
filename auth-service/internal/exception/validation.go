package exception

import (
	"github.com/go-playground/validator/v10"
)

func Validation(
	err error,
) error {

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return BadRequest(
			"validation failed",
		)
	}

	errors := make(
		map[string]string,
	)

	for _, e := range validationErrors {

		field := e.Field()

		switch e.Tag() {

		case "required":
			errors[field] = field + " is required"

		case "email":
			errors[field] = "invalid email format"

		case "min":
			errors[field] = field +
				" minimum length is " +
				e.Param()

		default:
			errors[field] = "invalid value"
		}
	}

	return ValidationError(
		errors,
	)
}
