package belajargolanglogging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T)  {
	logger := logrus.New()

	
	if file, err := os.OpenFile("applocation.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666); err == nil {
		logger.SetOutput(file) // masuk ke file applocation.log

		logger.Info("Hello Logging")
		logger.Warn("Hello Logging")
		logger.Error("Hello Logging")
	}
}