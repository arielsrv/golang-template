package server

import (
	"net/http"

	"github.com/internal/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	"regexp"
)

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
			ErrorHandler:          shared.ErrorHandler,
		}),
		config: Config{
			Recovery:  true,
			Swagger:   true,
			RequestID: true,
			Logger:    true,
			Cors:      false,
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

	if app.config.Cors {
		regex := regexp.MustCompile(`(^herokuapp\.com|^azurewebsites\.com)`)
		app.Add(http.MethodOptions, "/*", func(ctx *fiber.Ctx) error {
			origin := ctx.GetReqHeaders()["Origin"]
			matched := regex.MatchString(origin)
			if matched {
				ctx.Response().Header.Add("Vary", "Accept,Accept-Encoding,Origin,Access-Control-Request-Headers")
				ctx.Response().Header.Add("Cache-Control", "max-age=0")
				ctx.Response().Header.Add("Access-Control-Allow-Origin", origin)
				ctx.Response().Header.Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,HEAD")
				ctx.Response().Header.Add("Access-Control-Allow-Credentials", "true")
				return ctx.SendStatus(http.StatusNoContent)
			}
			return ctx.SendStatus(http.StatusForbidden)
		})
	}

	return app
}

type Config struct {
	Recovery  bool
	Swagger   bool
	RequestID bool
	Logger    bool
	Cors      bool
}
