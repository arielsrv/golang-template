package main

import (
	_ "github.com/golang-template/docs"
	"github.com/golang-template/internal/application"
	"github.com/golang-template/internal/common/container"
	"github.com/golang-template/internal/infrastructure/handlers"
	"github.com/golang-template/internal/infrastructure/webserver"
	"net/http"
)

// @title       Golang Template API
// @version     1.0
// @description This is a sample swagger for Golang Template API
// @BasePath    /
func main() {
	server := webserver.NewBuilder().
		Build()

	server.UseRecover()
	server.UseRequestID()
	server.UseLogger()
	server.UseSwagger()

	container.Register[application.IPingService](application.NewPingService())
	pingHandler := container.RegisterHandler(new(handlers.PingHandler))
	server.AddRoute(http.MethodGet, "/ping", pingHandler.Ping)

	server.Start()
}
