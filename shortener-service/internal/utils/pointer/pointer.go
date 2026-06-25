package pointer

func Value[T any](
	value *T,
) T {

	var zero T

	if value == nil {
		return zero
	}

	return *value
}
