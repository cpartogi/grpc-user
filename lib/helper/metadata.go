package helper

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func CheckRequestID(ctx context.Context) context.Context {

	md, _ := metadata.FromIncomingContext(ctx)

	if len(md["requestid"]) > 0 {
		return ctx
	} else {
		requestID := uuid.New().String()
		md.Append("requestid", requestID)
		return metadata.NewIncomingContext(ctx, md)
	}
}

func CheckUserId(ctx context.Context) (userId string, err error) {
	md, _ := metadata.FromIncomingContext(ctx)

	if len(md["userid"]) > 0 {
		userId = md["userid"][0]
		return
	} else {
		return userId, errors.New("failed get metadata")
	}
}
