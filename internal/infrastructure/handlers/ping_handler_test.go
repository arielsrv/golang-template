package handlers_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-template/internal/infrastructure/handlers"
	"github.com/stretchr/testify/mock"
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
	pingService *MockPingService
}

func (suite *PingControllerSuite) SetupTest() {
	suite.pingService = new(MockPingService)
	suite.pingHandler = handlers.NewPingHandler(suite.pingService)
	suite.app = fiber.New()
	suite.app.Get("/ping", suite.pingHandler.Ping())
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(PingControllerSuite))
}

type MockPingService struct {
	mock.Mock
}

func (m *MockPingService) Ping() string {
	args := m.Called()
	result := args.Get(0)
	return result.(string)
}

func (suite *PingControllerSuite) TestPing() {
	suite.pingService.On("Ping").Return("pong")
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
