package validation

import (
	"net/url"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ShortURL(
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

	u, err := url.ParseRequestURI(
		value,
	)

	if err != nil {
		return false
	}

	return u.Scheme == "http" ||
		u.Scheme == "https"
}
