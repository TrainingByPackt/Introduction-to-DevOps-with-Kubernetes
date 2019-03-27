package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {

	logger.SetFormatter(&logrus.JSONFormatter{})
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logger.Warn("LOG_LEVEL is not set!")
	}

	switch logLevel {
	case "DEBUG":
		logger.Level = logrus.DebugLevel
		break
	case "WARNING":
		logger.Level = logrus.WarnLevel
		break
	case "ERROR":
		logger.Level = logrus.ErrorLevel
		break
	default:
		logger.Level = logrus.InfoLevel
	}

	logger.Info("Application is starting..")

	connEnabled := os.Getenv("ENABLE_CONNECTION")
	if connEnabled == "" {
		logger.Warn("ENABLE_CONNECTION is not set!")
	} else if connEnabled != "Yes" && connEnabled != "No" {
		logger.Error("ENABLE_CONNECTION can only be set to [Yes] or [No]!")
	} else {
		logger.Debugf("ENABLE_CONNECTION set to [%v]", connEnabled)
	}

	if connEnabled == "Yes" {
		logger.Info("Application is running..")
		time.Sleep(999999 * time.Second)
	}

	logger.Info("Application is exiting..")
	time.Sleep(10 * time.Second)
}
