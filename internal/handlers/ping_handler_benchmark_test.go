package handlers_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-template/internal/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func BenchmarkPingHandler_Ping(b *testing.B) {
	pingService := new(MockPingService)
	pingHandler := handlers.NewPingHandler(pingService)
	app := fiber.New()
	app.Get("/ping", pingHandler.Ping)

	pingService.On("Ping").Return("pong")

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		response, err := app.Test(request)
		if err != nil || response.StatusCode != http.StatusOK {
			log.Print("f[" + strconv.Itoa(i) + "] Status != OK (200)")
		}
	}
}
