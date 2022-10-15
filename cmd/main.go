package main

import (
	"fmt"
	"github.com/arielsrv/golang-toolkit/webserver/api"
	_ "github.com/golang-template/docs"
	"github.com/golang-template/internal/handlers"
	"github.com/golang-template/internal/services"
	"log"
	"net/http"
	"os"
)

// @title       Golang Template API
// @version     1.0
// @description This is a sample swagger for Golang Template API
// @BasePath    /
func main() {
	app := &api.Application{
		UseRecovery:  true,
		UseRequestID: true,
		UseLogger:    true,
		UseSwagger:   true,
	}

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	app.Register(http.MethodGet, "/ping", pingHandler.Ping)

	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://%s:%s/ping in the browser", host, port)
	log.Fatal(app.Start(address))
}
