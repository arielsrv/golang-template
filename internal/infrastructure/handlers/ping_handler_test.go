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

type PingHandlerSuite struct {
	suite.Suite
	app         *fiber.App
	pingHandler handlers.IPingHandler
	pingService *MockPingService
}

func (suite *PingHandlerSuite) SetupTest() {
	suite.pingService = new(MockPingService)
	suite.pingHandler = handlers.NewPingHandler(suite.pingService)
	suite.app = fiber.New()
	suite.app.Get("/ping", suite.pingHandler.Ping())
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(PingHandlerSuite))
}

type MockPingService struct {
	mock.Mock
}

func (mock *MockPingService) Ping() string {
	args := mock.Called()
	return args.Get(0).(string)
}

func (suite *PingHandlerSuite) TestPingHandler_Ping() {
	suite.pingService.
		On("Ping").
		Return("pong")

	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	response, err := suite.app.Test(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.Equal(http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	suite.NoError(err)
	suite.NotNil(body)

	suite.Equal("pong", string(body))
}
