package gui

import (
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

type WailsLogger struct {
	logger.Logger
}

func (w *WailsLogger) Debug(message string) {
	logrus.Debug(message)
}

func (w *WailsLogger) Error(message string) {
	logrus.Error(message)
}

func (w *WailsLogger) Fatal(message string) {
	logrus.Fatal(message)
}

func (w *WailsLogger) Info(message string) {
	logrus.Info(message)
}

func (w *WailsLogger) Print(message string) {
	logrus.Print(message)
}

func (w *WailsLogger) Trace(message string) {
	logrus.Trace(message)
}

func (w *WailsLogger) Warning(message string) {
	logrus.Warning(message)
}

func NewWailsLogger() *WailsLogger {
	return &WailsLogger{}
}
