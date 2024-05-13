package application

import (
	"context"
	"user-service/config"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
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
	Config         *config.Config
	Log            *logrus.Logger
}
