package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {

	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {

		name := strings.Split(
			field.Tag.Get("json"),
			",",
		)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return validate
}
