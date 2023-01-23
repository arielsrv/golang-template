package env_test

import (
	"testing"

	"github.com/src/main/app/config/env"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	actual := env.IsEmpty("")
	assert.True(t, actual)
}

func TestIsNotEmpty(t *testing.T) {
	actual := env.IsEmpty("value")
	assert.False(t, actual)
}

func TestGetScope(t *testing.T) {
	t.Setenv("SCOPE", "test")
	actual := env.GetScope()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "test", actual)
}

func TestGetEnv(t *testing.T) {
	actual := env.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "dev", actual)
}

func TestGetEnv_Custom(t *testing.T) {
	t.Setenv("app.env", "staging")
	actual := env.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "staging", actual)
}

func TestGetEnv_Prod(t *testing.T) {
	t.Setenv("SCOPE", "prod")
	actual := env.GetEnv()
	assert.NotEmpty(t, actual)
	assert.Equal(t, "prod", actual)
}

func TestIsDev(t *testing.T) {
	assert.True(t, env.IsDev())
}
