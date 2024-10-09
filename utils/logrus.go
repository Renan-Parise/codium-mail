package utils

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
)

var logging = logrus.New()

func InitElasticAPM() {
	logrus.AddHook(&apmlogrus.Hook{})
}

func InitLogger() {
	logging.Out = os.Stdout

	logging.SetLevel(logrus.InfoLevel)

	logging.SetFormatter(&logrus.JSONFormatter{})

	logging.AddHook(&apmlogrus.Hook{})
}

func GetLogger() *logrus.Logger {
	return logging
}
