package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/golang-template/docs"
	"github.com/golang-template/internal/handlers"
	"github.com/golang-template/internal/server"
	"github.com/golang-template/internal/services"
)

// @title       Golang Template API
// @version     1.0
// @description This is a sample swagger for Golang Template API
// @BasePath    /.
func main() {
	app := server.New()

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	petHandler := handlers.NewPetHandler()

	app.Add(http.MethodGet, "/ping", pingHandler.Ping)
	app.Add(http.MethodGet, "/pets", petHandler.GetAll)
	app.Add(http.MethodGet, "/pets/:petID", petHandler.GetPetByID)
	app.Add(http.MethodPost, "/pets", petHandler.Create)

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
