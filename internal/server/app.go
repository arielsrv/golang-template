package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"net/http"
)

type App struct {
	*fiber.App
	config Config
	routes list.List
}

func Fiber(ctx *Context) *fiber.Ctx {
	return ctx.Current.(*fiber.Ctx)
}

func (app *App) Start(addr string) error {
	return app.Build().Listen(addr)
}

func (app *App) Build() *App {
	for node := app.routes.Front(); node != nil; node = node.Next() {
		route, converted := node.Value.(Route)
		if !converted {
			log.Fatalf("Cannot parse route.")
		}
		app.Add(route.Verb, route.Path, func(ctx *fiber.Ctx) error {
			return route.Action(&Context{Current: ctx})
		})
	}
	return app
}

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

type Config struct {
	Recovery  bool
	Swagger   bool
	RequestID bool
	Logger    bool
}

type Route struct {
	Path   string
	Verb   string
	Action func(context *Context) error
}

type Context struct {
	Current interface{}
}

func (app *App) Register(verb string, path string, action func(context *Context) error) *App {
	app.routes.PushBack(Route{
		Path:   path,
		Verb:   verb,
		Action: action,
	})

	return app
}
