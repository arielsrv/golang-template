package shared

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Details    string `json:"details,omitempty"`
}

func NewError(statusCode int, message string) *Error {
	return &Error{StatusCode: statusCode, Message: message}
}

func (e Error) Error() string {
	return e.Message
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	e := &Error{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
		Details:    err.Error(),
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		return ctx.Status(fiberError.Code).SendString(fiberError.Error())
	}

	var apiError *Error
	if errors.As(err, &apiError) {
		e.StatusCode = apiError.StatusCode
		e.Message = apiError.Message
		ctx.Status(e.StatusCode) //nolint:nolintlint,errcheck
	}

	return ctx.Status(e.StatusCode).JSON(&Error{
		StatusCode: e.StatusCode,
		Message:    e.Message,
		Details:    e.Details,
	})
}
