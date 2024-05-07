package application

import (
	"context"

	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
)

type ServiceApp struct {
	DbRead         *pg.DB
	DbWrite        *pg.DB
	GrpcServer     *grpc.Server
	Ctx            context.Context
	GrpcClientConn map[string]*grpc.ClientConn
	ServiceName    string
	ServiceMode    string
}
