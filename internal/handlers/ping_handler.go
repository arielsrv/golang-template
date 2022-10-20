package handlers

import (
	"github.com/golang-template/internal/server"
	"github.com/golang-template/internal/services"
)

type IPingHandler interface {
	Ping(ctx *server.Context) error
}

type PingHandler struct {
	pingService services.IPingService
}

func NewPingHandler(pingService services.IPingService) *PingHandler {
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
// @Router      /ping [get].
func (handler PingHandler) Ping(ctx *server.Context) error {
	result := handler.pingService.Ping()
	return server.Fiber(ctx).SendString(result)
}
