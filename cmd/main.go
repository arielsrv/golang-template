package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arielsrv/golang-toolkit/server"

	_ "github.com/docs"
	"github.com/internal/handlers"
	"github.com/internal/services"
)

// @title       Golang Template API
// @version     1.0
// @description This is a sample swagger for Golang Template API
// @BasePath    /
func main() {
	app := server.New(server.Config{
		Recovery:  true,
		Swagger:   true,
		RequestID: true,
		Logger:    true,
	})

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	server.RegisterHandler(pingHandler.Ping)

	app.Add(http.MethodGet, "/ping", server.Use(handlers.PingHandler{}.Ping))

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
