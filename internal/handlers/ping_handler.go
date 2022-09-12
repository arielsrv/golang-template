package handlers

import "github.com/gofiber/fiber/v2"

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (handler PingHandler) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	}
}
