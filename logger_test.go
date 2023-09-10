package belajargolanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T)  {
	logger := logrus.New()

	logger.Println("Hello Logger")
	// time="2023-09-11T01:42:28+07:00" level=info msg="Hello Logger"
	fmt.Println("Hello world")
	// Hello world
}

func TestLevel(t *testing.T)  {
	logger := logrus.New()

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")

// time="2023-09-11T01:46:29+07:00" level=info msg="This is Info"
// time="2023-09-11T01:46:29+07:00" level=warning msg="This is Warn"
// time="2023-09-11T01:46:29+07:00" level=error msg="This is Error"
// time="2023-09-11T01:46:29+07:00" level=fatal msg="This is Fatal"
// FAIL    github.com/abrarnaim015/belajar-golang-logging  0.564s
}

func TestLoggingLevel(t *testing.T)  {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // <- set level

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")

	// time="2023-09-11T01:48:43+07:00" level=trace msg="This is Trace"
	// time="2023-09-11T01:48:43+07:00" level=debug msg="This is Debug"
	// time="2023-09-11T01:48:43+07:00" level=info msg="This is Info"
	// time="2023-09-11T01:48:43+07:00" level=warning msg="This is Warn"
	// time="2023-09-11T01:48:43+07:00" level=error msg="This is Error"
	// time="2023-09-11T01:48:43+07:00" level=fatal msg="This is Fatal"
	// FAIL    github.com/abrarnaim015/belajar-golang-logging  0.651s
}