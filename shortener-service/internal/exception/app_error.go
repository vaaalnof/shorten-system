package exception

import "github.com/gofiber/fiber/v2"

type AppError struct {
	StatusCode int               `json:"-"`
	Message    string            `json:"message"`
	Errors     map[string]string `json:"errors,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(
	statusCode int,
	message string,
) *AppError {

	return &AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func BadRequest(
	message string,
) *AppError {

	return New(
		fiber.StatusBadRequest,
		message,
	)
}

func Unauthorized(
	message string,
) *AppError {

	return New(
		fiber.StatusUnauthorized,
		message,
	)
}

func Forbidden(
	message string,
) *AppError {

	return New(
		fiber.StatusForbidden,
		message,
	)
}

func NotFound(
	message string,
) *AppError {

	return New(
		fiber.StatusNotFound,
		message,
	)
}

func Conflict(
	message string,
) *AppError {

	return New(
		fiber.StatusConflict,
		message,
	)
}

func TooManyRequests(
	message string,
) *AppError {

	return New(
		fiber.StatusTooManyRequests,
		message,
	)
}

func UnprocessableEntity(
	message string,
) *AppError {

	return New(
		fiber.StatusUnprocessableEntity,
		message,
	)
}

func Internal(
	message string,
) *AppError {

	return New(
		fiber.StatusInternalServerError,
		message,
	)
}

func ValidationError(
	errors map[string]string,
) *AppError {

	return &AppError{
		StatusCode: fiber.StatusBadRequest,
		Message:    "validation failed",
		Errors:     errors,
	}
}
