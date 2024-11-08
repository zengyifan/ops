package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetupLogger(logFilePath string) io.Writer {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file %s for output: %s", logFilePath, err)
		return nil
	}

	Logger = logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{} // 设置为JSON格式
	multiWriter := io.MultiWriter(logFile, os.Stderr)
	Logger.Out = multiWriter
	return Logger.Out
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args)
}

func Errorln(args ...interface{}) {
	Logger.Errorln(args)
}
