package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

type formatterHook struct {
	writer    io.Writer
	formatter logrus.Formatter
	levels    []logrus.Level
}

func (f *formatterHook) Fire(entry *logrus.Entry) error {
	data, err := f.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = f.writer.Write(data)
	return err
}

func (f *formatterHook) Levels() []logrus.Level {
	return f.levels
}

func Setup() {
	startUpTimeStr := time.Now().Format("2006-01-02-150405")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	loggerFile := filepath.Join(wd, "logs", fmt.Sprintf("%s.log", startUpTimeStr))

	fileHook := &formatterHook{
		writer: &lumberjack.Logger{
			Filename:   loggerFile,
			MaxBackups: 7,
			MaxAge:     30,
			LocalTime:  true,
			Compress:   true,
		},
		formatter: &logrus.JSONFormatter{},
		levels:    logrus.AllLevels,
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.AddHook(fileHook)
	logrus.SetLevel(logrus.DebugLevel)
}
