package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/simonkimi/minebangumi/tools/dir"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

func Setup() {
	startUpTimeStr := time.Now().Format("2006-01-02-150405")
	loggerFile := filepath.Join(dir.GetTempDir(), "logs", fmt.Sprintf("%s.log", startUpTimeStr))
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   loggerFile,
		MaxBackups: 7,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{})
}
