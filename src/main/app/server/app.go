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
)

var handlers = make(map[string]any)
var routes []Route

type App struct {
	*fiber.App
	config Config
}

type Route struct {
	Verb   string
	Path   string
	Action func(ctx *fiber.Ctx) error
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

func Register(verb string, path string, action func(ctx *fiber.Ctx) error) {
	route := &Route{
		Verb:   verb,
		Path:   path,
		Action: action,
	}
	routes = append(routes, *route)
}

func RegisterHandler(handler any) {
	key := getType(handler)
	handlers[key] = handler
}

func Resolve[T any](_ ...T) *T {
	args := make([]T, 1)
	key := getType(args[0])
	return handlers[key].(*T)
}

func getType(value any) string {
	name := reflect.TypeOf(value)
	if name.Kind() == reflect.Ptr {
		name = name.Elem()
	}

	return name.String()
}
