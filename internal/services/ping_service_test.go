package services_test

import (
	"github.com/golang-template/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPingService_Ping(t *testing.T) {
	pingService := services.NewPingService()
	actual := pingService.Ping()

	assert.Equal(t, "pong", actual)
}
