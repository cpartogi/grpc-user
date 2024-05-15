package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	requestIDKey = "requestID"
)

func GetContext(ctx context.Context) (c context.Context) {

	//check requestID
	_, ok := ctx.Value(requestIDKey).(string)

	if !ok {
		return context.WithValue(ctx, requestIDKey, uuid.New().String())
	}

	return ctx
}

func GetLogger(ctx context.Context, funcName, errMsg string, req, res interface{}) *logrus.Logger {

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	requestID, ok := ctx.Value(requestIDKey).(string)
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
