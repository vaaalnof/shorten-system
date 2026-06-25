package validation

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var shortCodeRegex = regexp.MustCompile(
	`^[a-zA-Z0-9_-]+$`,
)

func ShortCode(
	fl validator.FieldLevel,
) bool {

	field := fl.Field()

	if field.Kind() == reflect.Ptr {

		if field.IsNil() {
			return true
		}

		field = field.Elem()
	}

	value, ok := field.Interface().(string)

	if !ok {
		return false
	}

	if value == "" {
		return true
	}

	return shortCodeRegex.MatchString(
		value,
	)
}
