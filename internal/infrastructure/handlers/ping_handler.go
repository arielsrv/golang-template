package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-template/internal/application"
)

type PingHandler struct {
	pingService application.IPingService
}

func NewPingHandler(pingService application.IPingService) *PingHandler {
	return &PingHandler{
		pingService: pingService,
	}
}

func (handler PingHandler) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		result := handler.pingService.Ping()
		return ctx.SendString(result)
	}
}
