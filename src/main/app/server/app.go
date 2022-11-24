package server

import (
	"log"
	"net/http"
	"reflect"

	"github.com/src/main/app/server/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	"runtime"
)

var routes = make(map[string]func(ctx *fiber.Ctx) error)

type App struct {
	*fiber.App
	config Config
}

func (app *App) Start(addr string) error {
	return app.Listen(addr)
}

func New(config ...Config) *App {
	app := &App{
		App: fiber.New(fiber.Config{
			DisableStartupMessage: true,
			ErrorHandler:          errors.ErrorHandler,
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
		log.Println("Swagger enabled")
		app.Add(http.MethodGet, "/swagger/*", swagger.HandlerDefault)
	}

	return app
}

type Config struct {
	Recovery  bool
	Swagger   bool
	RequestID bool
	Logger    bool
}

func RegisterHandler(action func(ctx *fiber.Ctx) error) {
	name := getFuncName(action)
	routes[name] = action
}

func Use(action func(ctx *fiber.Ctx) error) func(ctx *fiber.Ctx) error {
	name := getFuncName(action)
	return routes[name]
}

func getFuncName(action func(ctx *fiber.Ctx) error) string {
	return runtime.FuncForPC(reflect.ValueOf(action).Pointer()).Name()
}
