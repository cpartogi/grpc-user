package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func GetLogger(funcName, errMsg, requestId string, req, res interface{}) *logrus.Logger {

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	if errMsg != "" {
		log.WithFields(logrus.Fields{
			"requestId": requestId,
			"function":  funcName,
			"request":   req,
			"resp":      res,
		}).Error(errMsg)
	} else {
		log.WithFields(logrus.Fields{
			"requestId": requestId,
			"function":  funcName,
			"request":   req,
			"resp":      res,
		}).Info("Success")
	}

	log.SetReportCaller(true)
	log.Out = os.Stderr
	return log

}
