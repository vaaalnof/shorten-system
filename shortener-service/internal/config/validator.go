package config

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
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
