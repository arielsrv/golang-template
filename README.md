# golang-template

[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-89.0%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=arielsrv_golang-template&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=arielsrv_golang-template)

## Developer tools

- [Golang Lint](https://golangci-lint.run/)
- [Golang Task](https://taskfile.dev/)
- [Golang Dependencies Update](https://github.com/oligot/go-mod-upgrade)

### For macOs

- [Homebrew](https://brew.sh/index_es)

```shell
brew install go-task/tap/go-task
brew install golangci-lint
go install github.com/oligot/go-mod-upgrade@latest
```

## template

### main

```go
package main

import (
	"log"

	"github.com/src/main/app"
	_ "github.com/src/resources/docs"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### bootstrap

```go
package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		NewRelic:  false,
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
```

```shell
go test ./... -bench=.
```

````text
goos: darwin
goarch: arm64
pkg: github.com/internal/handlers
BenchmarkPingHandler_Ping-8        22664             53260 ns/op
````

## building

```shell
task build
```

## running (from output binary)

```shell
task run
```

## lint [included rules](.golangci.yml)

```shell
task lint
```

## test

```shell
task test
```

## coverage

```shell
task coverage
```

## upgrade packages

```shell
task download upgrade
```

## swagger [docs](/docs)

```shell
task swagger
```

## example request

```shell
curl 'http://localhost:8080/ping' --verbose
```

```text
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /ping HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.85.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 13 Sep 2022 11:57:44 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 4
< X-Request-Id: e9f18d4a-6a5f-46c1-bef2-880a5c78535d
<
* Connection #0 to host localhost left intact
pong
```

## example error response

```shell
curl 'http://localhost:8080/ping' --verbose
```

```text
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /pets HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.86.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 500 Internal Server Error
< Date: Wed, 02 Nov 2022 08:43:44 GMT
< Content-Type: application/json
< Content-Length: 50
< X-Request-Id: e6d61deb-0bbf-40fe-882a-9b246a72194b
<
* Connection #0 to host localhost left intact
{"status_code":500,"message":"unhealthy instance"}
```

## Configuration

Environment configuration is based on **Archaius Config**, you should use a similar folder
structure.
*SCOPE* env variable in remote environment is required

```
└── config
    ├── config.yml (shared config)
    └── dev
        └── config.yml (for local development)
    └── prod (for remote environment)
        └── config.yml (base config)
        └── {environment}.config.yml (base config)
```

The SDK provides a simple configuration hierarchy

* resources/config/config.properties (shared config)
* resources/config/{environment}/config.properties (override shared config by environment)
* resources/config/{environment}/{scope}.config.properties (override env and shared config by scope)

example *test.pets-api.internal.com*

```
└── config
    ├── config.yml                              3th (third)
    └── dev
        └── config.yml                          <ignored>
    └── prod
        └── config.yml (base config)            2nd (second)
        └── test.config.yml (base config)       1st (first)
```

* 1st (first)   prod/test.config.yml
* 2nd (second)  prod/config.yml
* 3th (third)   config.yml

```
2022/11/20 13:24:26 INFO: Two files have same priority. keeping
    /resources/config/prod/test.config.yml value
2022/11/20 13:24:26 INFO: Configuration files:
    /resources/config/prod/test.config.yml,
    /resources/config/prod/config.yml,
    /resources/config/config.yml
2022/11/20 13:24:26 INFO: invoke dynamic handler:FileSource
2022/11/20 13:24:26 INFO: enable env source
2022/11/20 13:24:26 INFO: invoke dynamic handler:EnvironmentSource
2022/11/20 13:24:26 INFO: archaius init success
2022/11/20 13:24:26 INFO: ENV: prod, SCOPE: test
2022/11/20 13:24:26 INFO: create new watcher
2022/11/20 13:24:26 Listening on port 8080
2022/11/20 13:24:26 Open http://127.0.0.1:8080/ping in the browser
```
