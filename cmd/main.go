package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-template/internal/handlers"
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

	app.Add(fiber.MethodGet, "/ping", handlers.
		NewPingHandler().
		Ping())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	log.Printf("%s:%s", host, port)
	log.Fatalln(app.Listen(fmt.Sprintf("%s:%s", host, port)))
}
