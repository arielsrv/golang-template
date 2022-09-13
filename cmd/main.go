package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-template/internal/application"
	"github.com/golang-template/internal/infrastructure/handlers"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	pingService := application.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)
	app.Add(fiber.MethodGet, "/ping", pingHandler.Ping())

	port, host := getAddress()
	address := fmt.Sprintf("%s:%s", host, port)

	log.Print(address)
	log.Fatalln(app.Listen(address))
}

func getAddress() (string, string) {
	port := getPort()
	host := getHost()
	return port, host
}

func getHost() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	return host
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
