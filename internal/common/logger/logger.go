package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	l.SetLevel(logrus.DebugLevel)
	return &Logger{l}
}
