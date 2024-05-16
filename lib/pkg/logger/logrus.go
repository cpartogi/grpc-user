package logger

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func Log(ctx context.Context, funcName, errMsg string, req, res interface{}) *logrus.Logger {

	var requestID string
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	md, _ := metadata.FromIncomingContext(ctx)

	if len(md["requestid"]) > 0 {
		requestID = md["requestid"][0]
	} else {
		requestID = uuid.New().String()
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
