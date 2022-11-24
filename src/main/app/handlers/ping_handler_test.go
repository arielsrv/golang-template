package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/src/main/app/server"

	"github.com/src/main/app/handlers"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PingHandlerSuite struct {
	suite.Suite
	app         *server.App
	pingHandler handlers.IPingHandler
	pingService *MockPingService
}

func (suite *PingHandlerSuite) SetupTest() {
	suite.pingService = new(MockPingService)
	suite.pingHandler = handlers.NewPingHandler(suite.pingService)
	suite.app = server.New()
	suite.app.Add(http.MethodGet, "/ping", suite.pingHandler.Ping)
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
	suite.pingService.On("Ping").Return("pong")

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
