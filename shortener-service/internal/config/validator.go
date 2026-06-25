package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	validationutil "shortener-service/internal/validation"
)

func NewValidator() *validator.Validate {

	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)

	// =====================================================
	// JSON FIELD NAME
	// =====================================================

	validate.RegisterTagNameFunc(
		func(
			field reflect.StructField,
		) string {

			name := strings.Split(
				field.Tag.Get("json"),
				",",
			)[0]

			if name == "-" {
				return ""
			}

			return name
		},
	)

	// =====================================================
	// CUSTOM VALIDATIONS
	// =====================================================

	mustRegisterValidation(
		validate,
		"shortcode",
		validationutil.ShortCode,
	)

	mustRegisterValidation(
		validate,
		"shorturl",
		validationutil.ShortURL,
	)

	return validate
}

func mustRegisterValidation(
	validate *validator.Validate,
	tag string,
	fn validator.Func,
) {

	if err := validate.RegisterValidation(
		tag,
		fn,
	); err != nil {

		panic(
			"failed to register validation: " +
				tag +
				": " +
				err.Error(),
		)
	}
}
