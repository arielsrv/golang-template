package handlers_test

import (
	"github.com/arielsrv/golang-toolkit/webserver/api"
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
	app := new(api.Application)
	app.Register(http.MethodGet, "/ping", pingHandler.Ping)
	app.Build()

	pingService.On("Ping").Return("pong")

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		response, err := app.FiberApp.Test(request)
		if err != nil || response.StatusCode != http.StatusOK {
			log.Print("f[" + strconv.Itoa(i) + "] Status != OK (200)")
		}
	}
}
