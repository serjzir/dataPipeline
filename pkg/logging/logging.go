package logging

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func Init() *Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",

	}
	return &Logger{l}
}
