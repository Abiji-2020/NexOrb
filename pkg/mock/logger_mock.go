package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) LogError(msg string) {
	m.Called(msg)
}

func (m *MockLogger) LogInfo(msg string) {
	m.Called(msg)
}

func (m *MockLogger) LogDebug(msg string) {
	m.Called(msg)
}

func (m *MockLogger) LogWarn(msg string) {
	m.Called(msg)
}	
