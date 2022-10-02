package application_test

import (
	"github.com/golang-template/internal/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPingService_Ping(t *testing.T) {
	pingService := application.NewPingService()
	actual := pingService.Ping()

	assert.Equal(t, "pong", actual)
}
