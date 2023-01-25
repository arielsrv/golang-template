package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/src/main/app/config"
	"github.com/src/main/app/config/env"
	"github.com/src/main/app/handlers"
	"github.com/src/main/app/server"
	"github.com/src/main/app/services"
)

func Run() error {
	app := server.New(server.Config{
		Recovery:  true,
		Swagger:   true,
		RequestID: true,
		Logger:    true,
	})

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	server.RegisterHandler(pingHandler)

	server.Register(http.MethodGet, "/ping", server.Resolve[handlers.PingHandler]().Ping)

	host := config.String("HOST")
	if env.IsEmpty(host) && !env.IsDev() {
		host = "0.0.0.0"
	} else {
		host = "127.0.0.1"
	}

	port := config.String("PORT")
	if env.IsEmpty(port) {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://%s:%s/ping in the browser", host, port)

	return app.Start(address)
}
