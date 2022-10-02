package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-template/internal/application"
)

type IPingHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type PingHandler struct {
	pingService application.IPingService
}

func NewPingHandler(pingService application.IPingService) *PingHandler {
	return &PingHandler{
		pingService: pingService,
	}
}

// Ping godoc
// @Summary     Check if the instance is online
// @Description Ping
// @Tags        Check
// @Success     200
// @Produce     plain
// @Success     200 {string} string "pong"
// @Router      /ping [get]
func (handler PingHandler) Ping(ctx *fiber.Ctx) error {
	result := handler.pingService.Ping()
	return ctx.SendString(result)
}
