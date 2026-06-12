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

			errors[field] =
				field + " is required"

		case "url":

			errors[field] =
				"invalid url format"

		case "uuid":

			errors[field] =
				"invalid uuid format"

		case "min":

			errors[field] =
				field +
					" minimum length is " +
					e.Param()

		case "max":

			errors[field] =
				field +
					" maximum length is " +
					e.Param()

		case "oneof":

			errors[field] =
				field +
					" must be one of: " +
					e.Param()

		default:

			errors[field] =
				"invalid value"
		}
	}

	return ValidationError(
		errors,
	)
}
