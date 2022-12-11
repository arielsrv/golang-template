package errors

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func NewError(statusCode int, message string) *Error {
	return &Error{StatusCode: statusCode, Message: message}
}

func (e Error) Error() string {
	return e.Message
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var e = new(Error)

	var fiberError *fiber.Error
	var apiError *Error

	switch {
	case errors.As(err, &fiberError):
		{
			e.StatusCode = fiberError.Code
			e.Message = fiberError.Message
		}
	case errors.As(err, &apiError):
		{
			e.StatusCode = apiError.StatusCode
			e.Message = apiError.Message
		}
	default:
		e.StatusCode = http.StatusInternalServerError
		e.Message = err.Error()
	}

	ctx.Status(e.StatusCode)

	return ctx.JSON(e)
}
