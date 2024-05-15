package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

const (
	requestIDKey = "requestID"
)

func GetContext(ctx context.Context) (c context.Context) {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return ctx
	}

	requestID := md["requestid"]

	return context.WithValue(ctx, requestIDKey, requestID[0])
}

func GetLogger(ctx context.Context, funcName, errMsg string, req, res interface{}) *logrus.Logger {

	var requestID string
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		requestID = md["requestid"][0]
	}

	// requestID, ok := ctx.Value(requestIDKey).(string)
	// if !ok {
	// 	fmt.Println("No Request ID in http request")
	// }

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
