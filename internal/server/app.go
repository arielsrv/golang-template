package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"net/http"
)

// App Fiber Wrapper
//
// Config enabled by default
type App struct {
	*fiber.App
	config Config
}

// Start server
func (app *App) Start(addr string) error {
	return app.Listen(addr)
}

// New Create a new Fiber Server
// Use Config for disable recovery, swagger, requestID and Logger middlewares.
// All configs are enabled by default
func New(config ...Config) *App {
	app := &App{
		App: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
		config: Config{
			Recovery:  true,
			Swagger:   true,
			RequestID: true,
			Logger:    true,
		},
	}

	if len(config) > 0 {
		app.config = config[0]
	}

	if app.config.Recovery {
		app.Use(recover.New(recover.Config{
			EnableStackTrace: true,
		}))
	}

	if app.config.RequestID {
		app.Use(requestid.New())
	}

	if app.config.Logger {
		app.Use(logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		}))
	}

	if app.config.Swagger {
		app.Add(http.MethodGet, "/swagger/*", swagger.HandlerDefault)
	}

	return app
}

// Config fiber options wrapper
//
// Recovery option is useful to handle panics.
//
// Swagger option enable UI /swagger/*.
//
// RequestID option include unique identifier for incoming request.
//
// Logger option enable request logging
type Config struct {
	Recovery  bool
	Swagger   bool
	RequestID bool
	Logger    bool
}
