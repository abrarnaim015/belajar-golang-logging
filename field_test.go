package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T)  {
	logger := logrus.New()

	logger.WithField("firstName", "Abrar").Info("logging ...")
	logger.WithField("firstName", "Abrar").WithField("lastName", "Naim").Info("logging ...")

	// time="2023-09-11T02:07:50+07:00" level=info msg="logging ..." firstName=Abrar
	// time="2023-09-11T02:07:50+07:00" level=info msg="logging ..." firstName=Abrar lastName=Naim
}

func TestFieldFormaterJson(t *testing.T)  {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("firstName", "Abrar").Info("logging ...")
	logger.WithField("firstName", "Abrar").WithField("lastName", "Naim").Info("logging ...")

	// {"firstName":"Abrar","level":"info","msg":"logging ...","time":"2023-09-11T02:09:35+07:00"}
	// {"firstName":"Abrar","lastName":"Naim","level":"info","msg":"logging ...","time":"2023-09-11T02:09:35+07:00"}
}

func TestFielsdFormaterJson(t *testing.T)  {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"fistName": "Abrar",
		"lastName": "Naim",
	}).Info("logging ...")

	// {"firstName":"Abrar","lastName":"Naim","level":"info","msg":"logging ...","time":"2023-09-11T02:09:35+07:00"}
}