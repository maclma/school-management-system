package logger

import (
	"io"
	"os"
	"school-management-system/internal/config"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(cfg *config.Config) *logrus.Logger {
	log := logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	// Set log format
	if cfg.LogFormat == "json" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     cfg.AppEnv == "development",
		})
	}

	// Set output
	log.SetOutput(os.Stdout)

	// Add file output in production
	if cfg.AppEnv == "production" {
		file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(mw)
		}
	}

	Log = log
	return log
}

func GetLogger() *logrus.Logger {
	if Log == nil {
		log := logrus.New()
		log.SetLevel(logrus.InfoLevel)
		log.SetFormatter(&logrus.TextFormatter{})
		Log = log
	}
	return Log
}
