package handlers_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-template/internal/handlers"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingControllerSuite struct {
	suite.Suite
	app         *fiber.App
	pingHandler *handlers.PingHandler
}

func (suite *PingControllerSuite) SetupTest() {
	suite.pingHandler = handlers.NewPingHandler()
	suite.app = fiber.New()
	suite.app.Get("/ping", suite.pingHandler.Ping())
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(PingControllerSuite))
}

func (suite *PingControllerSuite) TestPing() {
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	response, err := suite.app.Test(request)
	suite.NotNil(response)
	suite.NoError(err)

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	suite.NotNil(body)
	suite.NoError(err)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("pong", string(body))
}
