package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Instance *logrus.Logger
}

func InitLogger() *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)
	return &Logger{Instance: log}
}

func (l *Logger) LogInfo(msg string) {
	l.Instance.Info(msg)
}

func (l *Logger) LogError(msg string) {
	l.Instance.Error(msg)
}

func (l *Logger) LogFatal(msg string) {
	l.Instance.Fatal(msg)
}

func (l *Logger) LogWarn(msg string) {
	l.Instance.Warn(msg)
}

func (l *Logger) LogDebug(msg string) {
	l.Instance.Debug(msg)
}
