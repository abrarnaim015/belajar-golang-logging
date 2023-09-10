package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T)  {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	
	entry := logrus.NewEntry(logger)

	entry.WithField("userName", "Naim")
	entry.Info("Hello Entry")

	// {"level":"info","msg":"Hello Entry","time":"2023-09-11T02:19:07+07:00"}
}