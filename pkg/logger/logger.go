package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func Setup() {
	startUpTimeStr := time.Now().Format("2006-01-02-150405")
	loggerFile := fmt.Sprintf("./logs/%s.log", startUpTimeStr)
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   loggerFile,
		MaxBackups: 7,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{})
}
