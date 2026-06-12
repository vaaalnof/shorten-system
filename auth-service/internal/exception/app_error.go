package exception

import "github.com/gofiber/fiber/v2"

type AppError struct {
	StatusCode int
	Message    string
	Errors     map[string]string
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

func TooManyRequests(
	message string,
) *AppError {
	return New(
		fiber.StatusTooManyRequests,
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

func ValidationError(
	errors map[string]string,
) *AppError {
	return &AppError{
		StatusCode: fiber.StatusBadRequest,
		Message:    "validation failed",
		Errors:     errors,
	}
}
