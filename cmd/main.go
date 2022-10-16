package main

import (
	"fmt"
	_ "github.com/golang-template/docs"
	"github.com/golang-template/internal/app"
	"github.com/golang-template/internal/handlers"
	"github.com/golang-template/internal/services"
	"log"
	"net/http"
	"os"
)

// @title       Golang Template App
// @version     1.0
// @description This is a sample swagger for Golang Template App
// @BasePath    /
func main() {
	app := app.New(app.Config{
		Recovery:  true,
		Swagger:   true,
		RequestID: true,
		Logger:    true,
	})

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
