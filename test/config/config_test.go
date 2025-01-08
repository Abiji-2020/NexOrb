package config_test 

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Abiji-2020/NexOrb/pkg/mock"
	"github.com/Abiji-2020/NexOrb/config"
)

func TestNewConfig(t *testing.T) {
	// Initialize the logger
	mocklog := mock.MockLog()

	mockdb := mock.InitTestDB()

	appConfig := &config.AppConfig{
		Router: mock.TestRouter(),
		Database: mockdb,
		Logger: mocklog,
		ServerPort: "8080",
	}

	assert.NotNil(t, appConfig, "NewConfig should return a valid AppConfig instance")
	assert.NotNil(t, appConfig.Router, "AppConfig should initialize with a valid Router instance")
	assert.NotNil(t, appConfig.Database, "AppConfig should initialize with a valid Database connection")
	assert.NotNil(t, appConfig.Logger, "AppConfig should initialize with a valid Logger instance")
	assert.Equal(t, "8080", appConfig.ServerPort, "AppConfig should be initialized with the specified server port '8080'")
}