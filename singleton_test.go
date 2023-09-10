package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T)  {
	logrus.Info("Hello World") // 1
	logrus.Warn("Hello World") // 2

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello World") // 3
	logrus.Warn("Hello World") // 4

	// jika langsung di panggil logrus nya maka akan di buat secara global
	// tidak terlalu di sarankan, tp jika memang ingin di set globat mangga
	
	// time="2023-09-11T02:32:19+07:00" level=info msg="Hello World"  // 1
	// time="2023-09-11T02:32:19+07:00" level=warning msg="Hello World" // 2
	// {"level":"info","msg":"Hello World","time":"2023-09-11T02:32:19+07:00"} // 3
	// {"level":"warning","msg":"Hello World","time":"2023-09-11T02:32:19+07:00"} // 4
}