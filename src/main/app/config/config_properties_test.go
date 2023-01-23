package config_test

import (
	"testing"

	"github.com/src/main/app/config"

	"github.com/stretchr/testify/assert"
)

func TestGetProperty(t *testing.T) {
	actual := config.String("value")
	assert.Equal(t, "shared", actual)
}

func TestGetProperty_Err(t *testing.T) {
	actual := config.String("missing")
	assert.Equal(t, "", actual)
}

func TestGetIntProperty_Err(t *testing.T) {
	actual := config.Int("missing")
	assert.Equal(t, 0, actual)
}
