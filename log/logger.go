package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type CustomFormatter struct{}

func (f CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// fmt.Println(entry.)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	level := ""
	switch entry.Level.String() {
	case "warning":
		level = "WARN "
	case "info":
		level = "INFO "
	case "fatal":
		level = "FATAL"
	}
	message := fmt.Sprintf("%s %s : %s\n", timestamp, level, entry.Message)
	return []byte(message), nil
}

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&CustomFormatter{})
	log.SetOutput(os.Stdout)
	return log
}

var logger *logrus.Logger

func init() {
	logger = NewLogger()
}

func Info(message string, data ...interface{}) {
	logger.Infof(message, data...)
}

func Warn(message string, data ...interface{}) {
	logger.Warnf(message, data...)
}

func Errorf(message string, data ...interface{}) {
	logger.Errorf(message, data...)
}

func Fatal(data ...interface{}) {
	logger.Fatal(data...)
}

func Panic(err error) {
	logger.Panic(err)
}
