package handlers_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang-template/internal/handlers"
	"github.com/golang-template/internal/server"
)

func BenchmarkPingHandler_Ping(b *testing.B) {
	pingService := new(MockPingService)
	pingHandler := handlers.NewPingHandler(pingService)
	app := server.New(server.Config{
		Logger: false,
	})
	app.Register(http.MethodGet, "/ping", pingHandler.Ping)

	pingService.On("Ping").Return("pong")

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		response, err := app.Test(request)
		if err != nil || response.StatusCode != http.StatusOK {
			log.Print("f[" + strconv.Itoa(i) + "] Status != OK (200)")
		}
	}
}
