package main

import (
	"context"
	"fmt"
	_ "github.com/golang-template/docs"
	"github.com/golang-template/internal/app"
	"github.com/golang-template/internal/handlers"
	"github.com/golang-template/internal/services"
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
	app := fx.New(
		fx.Provide(services.NewPingService),
		fx.Provide(handlers.NewPingHandler),
		fx.Provide(NewHandlers),
		fx.Invoke(Start),
		// fx.WithLogger(
		//	func() fxevent.Logger {
		//		return fxevent.NopLogger
		//	},
		// ),
	)
	app.Run()
}

type Handlers struct {
	pingHandler handlers.IPingHandler
}

func NewHandlers(pingHandler handlers.IPingHandler) *Handlers {
	return &Handlers{
		pingHandler: pingHandler,
	}
}

func Start(lifecycle fx.Lifecycle, handlers *Handlers) *app.App {
	app := server.New()
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			app.Add(http.MethodGet, "/ping", handlers.pingHandler.Ping)

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
			go app.Start(address) //nolint:nolintlint,errcheck
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
	return app
}
