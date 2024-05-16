package helper

import (
	"context"

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
