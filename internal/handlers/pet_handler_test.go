package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/internal/handlers"
	"github.com/internal/model"
	"github.com/internal/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/suite"
)

type PetHandlerSuite struct {
	suite.Suite
	app        *fiber.App
	petHandler handlers.IPetHandler
}

func (suite *PetHandlerSuite) SetupTest() {
	suite.petHandler = handlers.NewPetHandler()
	suite.app = fiber.New(fiber.Config{
		ErrorHandler: shared.ErrorHandler,
	})
	suite.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	suite.app.Add(http.MethodGet, "/pets", suite.petHandler.GetAll)
	suite.app.Add(http.MethodGet, "/pets/:petID", suite.petHandler.GetPetByID)
	suite.app.Add(http.MethodPost, "/pets", suite.petHandler.Create)
}

func TestPetHandlerSuite(t *testing.T) {
	suite.Run(t, new(PetHandlerSuite))
}

func (suite *PetHandlerSuite) TestPetHandler_GetPetByID() {
	request := httptest.NewRequest(http.MethodGet, "/pets/1", nil)
	response, err := suite.app.Test(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.Equal(http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	suite.NoError(err)
	suite.NotNil(body)

	var model model.PetModel
	err = json.Unmarshal(body, &model)
	suite.NoError(err)

	suite.NotNil(model)
	suite.Equal(int64(1), model.ID)
	suite.Equal("Rin Tin Tin", model.Name)
}

func (suite *PetHandlerSuite) TestPetHandler_GetAll() {
	request := httptest.NewRequest(http.MethodGet, "/pets", nil)
	response, err := suite.app.Test(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.Equal(http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	suite.NoError(err)
	suite.NotNil(body)

	var apiError shared.Error
	err = json.Unmarshal(body, &apiError)
	suite.NoError(err)

	suite.NotNil(apiError)
	suite.Equal(http.StatusNotFound, apiError.StatusCode)
	suite.Equal("no pets found", apiError.Message)
}

func (suite *PetHandlerSuite) TestPetHandler_Create() {
	request := httptest.NewRequest(http.MethodPost, "/pets", nil)
	response, err := suite.app.Test(request)
	suite.NoError(err)
	suite.NotNil(response)
	suite.Equal(http.StatusInternalServerError, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	suite.NoError(err)
	suite.NotNil(body)

	var apiError shared.Error
	err = json.Unmarshal(body, &apiError)
	suite.NoError(err)

	suite.NotNil(apiError)
	suite.Equal(http.StatusInternalServerError, apiError.StatusCode)
	suite.Equal("runtime error: integer divide by zero", apiError.Message)
}
