package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func GetLogger(ctx context.Context, funcName, errMsg string, req, res interface{}) *logrus.Logger {

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		fmt.Println("No Request ID in http request")
	}

	if errMsg != "" {
		log.WithFields(logrus.Fields{
			"requestID": requestID,
			"function":  funcName,
			"request":   req,
			"resp":      res,
		}).Error(errMsg)
	} else {
		log.WithFields(logrus.Fields{
			"requestID": requestID,
			"function":  funcName,
			"request":   req,
			"resp":      res,
		}).Info("Success")
	}

	log.SetReportCaller(true)
	log.Out = os.Stderr
	return log

}
