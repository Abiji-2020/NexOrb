package config_test 

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Abiji-2020/NexOrb/pkg/mock"
	"github.com/Abiji-2020/NexOrb/config"
)

func TestNewConfig(t *testing.T) {
	// Initialize the logger
	mocklog := &mock.MockLogger{}

	mockdb := mock.InitTestDB()

	appConfig := &config.AppConfig{
		Router: mock.TestRouter(),
		DB: mockdb,
		Logger: mocklog,
		ServerPort: "8080",
	}

	assert.NotNil(t, appConfig, "AppConfig should not be nil")
	assert.NotNil(t, appConfig.Router,"Router should not be nil")
	assert.NotNil(t, appConfig.DB,"DB should not be nil")
	assert.NotNil(t, appConfig.Logger,"Logger should not be nil")
	assert.Equal(t, "8080", appConfig.ServerPort, "ServerPort should be 8080")
}