package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func (b *Handler) SendString(ctx *fiber.Ctx, body string) error {
	if body == "" {
		ctx.Status(http.StatusNotFound)
	}

	return ctx.SendString(body)
}

func (b *Handler) SendJSON(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(data)
}
