package mock

import (
	"github.com/sirupsen/logrus"
	"github.com/Abiji-2020/NexOrb/logger"
)

func MockLog() *logger.Logger {
	mockLog := logrus.New()
	return &logger.Logger{Instance: mockLog}
}