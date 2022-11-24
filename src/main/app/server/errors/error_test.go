package errors_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	errors2 "github.com/src/main/app/server/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestNewError(t *testing.T) {
	actual := errors2.NewError(http.StatusInternalServerError, "nil reference")
	assert.NotNil(t, actual)
	assert.Equal(t, http.StatusInternalServerError, actual.StatusCode)
	assert.Equal(t, "nil reference", actual.Error())
}

func TestErrorHandler(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := errors2.ErrorHandler(ctx, errors.New("api server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError errors2.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "api server error", apiError.Error())
}

func TestErrorHandler_FiberError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := errors2.ErrorHandler(ctx, fiber.NewError(http.StatusInternalServerError, "api server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError errors2.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "api server error", apiError.Error())
}

func TestErrorHandler_ApiError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := errors2.ErrorHandler(ctx, errors2.NewError(http.StatusInternalServerError, "api server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError errors2.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "api server error", apiError.Error())
}
