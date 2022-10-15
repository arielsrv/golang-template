package handlers_test

import (
	"github.com/arielsrv/golang-toolkit/webserver/api"
	"github.com/golang-template/internal/handlers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingHandlerSuite struct {
	suite.Suite
	pingHandler handlers.IPingHandler
	pingService *MockPingService
	app         *api.Application
}

func (suite *PingHandlerSuite) SetupTest() {
	suite.pingService = new(MockPingService)
	suite.pingHandler = handlers.NewPingHandler(suite.pingService)
	suite.app = new(api.Application)
	suite.app.Register(http.MethodGet, "/ping", suite.pingHandler.Ping)
	suite.app.Build()
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
	response, err := suite.app.FiberApp.Test(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.Equal(http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	suite.NoError(err)
	suite.NotNil(body)

	suite.Equal("pong", string(body))
}
