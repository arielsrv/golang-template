# golang-template
[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-97.2%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

## Developer tools
- [Golang Lint](https://golangci-lint.run/)
- [Golang Task](https://taskfile.dev/)
- [Golang Dependencies Update](https://github.com/oligot/go-mod-upgrade)

### For macOs
```shell
$ brew install go-task/tap/go-task
$ brew install golangci-lint
$ go install github.com/oligot/go-mod-upgrade@latest
```

## building

./task build

## running

./task run

## lint (.golangci.yml)

./task lint

## test

./task test

## coverage

./task coverage

## upgrade packages

./task download upgrade

## example request
```text
$ curl 'http://localhost:8080/ping' --verbose
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
