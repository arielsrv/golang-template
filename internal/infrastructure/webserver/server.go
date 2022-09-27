package webserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"log"
	"os"
)

type Builder struct {
	server Server
}

func NewBuilder() *Builder {
	return &Builder{
		server: Server{App: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		})},
	}
}

func (s *Builder) Build() Server {
	return s.server
}

type Server struct {
	*fiber.App
}

func (server *Server) UseSwagger() {
	server.Add(fiber.MethodGet, "/swagger/*", swagger.HandlerDefault)
}

func (server *Server) UseRecover() {
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
}

func (server *Server) UseRequestID() {
	server.Use(requestid.New())
}

func (server *Server) UseLogger() {
	server.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
}

func (server *Server) AddRoute(method string, url string, predicate func() string) {
	server.Add(method, url, SendString(predicate))
}

func SendString(predicate func() string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		result := predicate()
		return ctx.SendString(result)
	}
}

func (server *Server) Start() {
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
	log.Fatal(server.Listen(address))
}
