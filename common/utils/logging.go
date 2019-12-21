package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type ServiceLogger struct {
}

func NewServiceLogger() ServiceLogger {
	return ServiceLogger{}
}

func (s ServiceLogger) Log(v ...interface{}) {
	logrus.Infoln(v...)
}

func (s ServiceLogger) Logf(format string, v ...interface{}) {
	logrus.Infoln(fmt.Sprintf(format, v...))
}
