package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T)  {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.Warn("Hello Logging")
	logger.Error("Hello Logging")
	
	// {"level":"info","msg":"Hello Logging","time":"2023-09-11T02:03:59+07:00"}
	// {"level":"warning","msg":"Hello Logging","time":"2023-09-11T02:03:59+07:00"}
	// {"level":"error","msg":"Hello Logging","time":"2023-09-11T02:03:59+07:00"}
}