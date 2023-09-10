package belajargolanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	// code di sini logicnya
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T)  {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.AddHook(&SampleHook{})

	logger.Info("Hello Naim")
	// hook tidak terpanggil karna tidak termasuk dari levelnya
	// {"level":"info","msg":"Hello Naim","time":"2023-09-11T02:27:41+07:00"}
	logger.Warn("Hello Naim")
	// Sample Hook warning Hello Naim
	// {"level":"warning","msg":"Hello Naim","time":"2023-09-11T02:27:41+07:00"}
	logger.Error("Hello Naim")
	// Sample Hook error Hello Naim
	// {"level":"error","msg":"Hello Naim","time":"2023-09-11T02:27:41+07:00"}
}