package handlers

import (
	"github.com/golang-template/internal/application"
)

type PingHandler struct {
	PingService application.IPingService `inject:",type"`
}

// Ping godoc
// @Summary     Check if the instance is online
// @Description Ping
// @Tags        Check
// @Success     200
// @Produce     plain
// @Success     200 {string} string "pong"
// @Router      /ping [get]
func (handler PingHandler) Ping() string {
	return handler.PingService.
		Ping()
}
