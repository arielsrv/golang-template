# golang-template
[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

## Installation instructions

* Install [Golang Lint](https://golangci-lint.run/)
* Install [Golang Task](https://taskfile.dev/)
* Install [Golang Dependencies Update](https://github.com/oligot/go-mod-upgrade) (optional)

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

## curl
```text
$ curl 'http://localhost:8080/ping'
pong
```
