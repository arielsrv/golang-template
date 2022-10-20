package tests_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-template/internal/server"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingHandlerSuite struct {
	suite.Suite
	app *server.App
}

func (suite *PingHandlerSuite) SetupTest() {
	suite.app = server.New(server.Config{Cors: true})
	suite.app.Add(http.MethodGet, "/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(PingHandlerSuite))
}

func (suite *PingHandlerSuite) TestCorsRequest_Ok() {
	request := httptest.NewRequest(http.MethodOptions, "/ping", nil)
	request.Header.Set("Origin", "herokuapp.com")
	response, err := suite.app.Test(request)

	suite.NoError(err)
	suite.NotNil(response)

	suite.Equal(http.StatusNoContent, response.StatusCode)
	suite.Equal("Accept,Accept-Encoding,Origin,Access-Control-Request-Headers", response.Header.Get("Vary"))
	suite.Equal("max-age=0", response.Header.Get("Cache-Control"))
	suite.Equal("herokuapp.com", response.Header.Get("Access-Control-Allow-Origin"))
	suite.Equal("GET,POST,PUT,DELETE,PATCH,HEAD", response.Header.Get("Access-Control-Allow-Methods"))
	suite.Equal("true", response.Header.Get("Access-Control-Allow-Credentials"))

	body, err := io.ReadAll(response.Body)
	suite.NoError(err)
	suite.NotNil(body)
}

func (suite *PingHandlerSuite) TestCorsRequest_Forbidden() {
	request := httptest.NewRequest(http.MethodOptions, "/ping", nil)
	request.Header.Set("Origin", "domain.com")
	response, err := suite.app.Test(request)

	suite.NoError(err)
	suite.NotNil(response)

	suite.Equal(http.StatusForbidden, response.StatusCode)
}
