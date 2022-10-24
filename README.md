# golang-template

[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

## Developer tools

- [Golang Lint](https://golangci-lint.run/)
- [Golang Task](https://taskfile.dev/)
- [Golang Dependencies Update](https://github.com/oligot/go-mod-upgrade)

### For macOs

```shell
brew install go-task/tap/go-task
brew install golangci-lint
go install github.com/oligot/go-mod-upgrade@latest
```

## template

```go

package main

import (
	_ "github.com/docs" // only for Swagger
	"github.com/internal/handlers"
	"github.com/internal/server"
	"github.com/internal/services"
	"log"
	"net/http"
)

func main() {
	app := server.New(server.Config{
		Recovery:  true,
		Swagger:   false,
		RequestID: true,
		Logger:    true,
	})

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	app.Add(http.MethodGet, "/ping", pingHandler.Ping)

	log.Fatal(app.Start("localhost:8080"))
}
```

## benchmark

```go
package handlers_test

import (
	_ "github.com/docs" // only for swagger
	"github.com/internal/handlers"
	"github.com/internal/server"
	"github.com/internal/services"
	"log"
	"net/http"
)

func main() {
	app := server.New(server.Config{
		Recovery:  true,
		Swagger:   false,
		RequestID: true,
		Logger:    true,
	})

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	app.Add(http.MethodGet, "/ping", pingHandler.Ping)

	log.Fatal(app.Start("localhost:8080"))
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
