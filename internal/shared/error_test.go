package shared_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/internal/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestNewError(t *testing.T) {
	actual := shared.NewError(http.StatusInternalServerError, "nil reference")
	assert.NotNil(t, actual)
	assert.Equal(t, http.StatusInternalServerError, actual.StatusCode)
	assert.Equal(t, "nil reference", actual.Message)
	assert.Equal(t, "nil reference", actual.Error())
}

func TestErrorHandler(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := shared.ErrorHandler(ctx, errors.New("internal server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError shared.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "internal server error", apiError.Message)
}

func TestErrorHandler_FiberError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := shared.ErrorHandler(ctx, fiber.NewError(http.StatusInternalServerError, "internal server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError shared.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "internal server error", apiError.Message)
}

func TestErrorHandler_ApiError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := shared.ErrorHandler(ctx, shared.NewError(http.StatusInternalServerError, "internal server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError shared.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "internal server error", apiError.Message)
}
