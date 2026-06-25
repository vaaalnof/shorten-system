package exception

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validation(
	err error,
) error {

	var validationErrors validator.ValidationErrors

	if !errors.As(
		err,
		&validationErrors,
	) {
		return BadRequest(
			"validation failed",
		)
	}

	errorsMap := make(
		map[string]string,
	)

	for _, e := range validationErrors {

		field := e.Field()

		switch e.Tag() {

		case "required":

			errorsMap[field] =
				field + " is required"

		case "shorturl":

			errorsMap[field] =
				field + " must be a valid shorturl"

		case "min":

			errorsMap[field] =
				field +
					" minimum length is " +
					e.Param()

		case "max":

			errorsMap[field] =
				field +
					" maximum length is " +
					e.Param()

		default:

			errorsMap[field] =
				"invalid value"
		}
	}

	return ValidationError(
		errorsMap,
	)
}
