package logger

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/rohitnarayan/leaderboard/internal/config"
)

var logger *logrus.Logger

func SetupLogger(config *config.LoggerConfig) {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		log.Fatalf("[setupLogger] error parsing log level, err: %+v", err)
	}

	logger = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		},
	}

	return
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
