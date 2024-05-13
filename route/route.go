package route

import (
	"user-service/application"
	"user-service/config"

	"user-service/pb/user"

	userhandler "user-service/domain/user/handler/grpc"
	userrepo "user-service/domain/user/repo"
	userusecase "user-service/domain/user/usecase"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SetupRoute(cfg *config.Config, app *application.ServiceApp) {
	switch app.ServiceMode {
	case "grpc":
		GrpcRoute(app.GrpcServer,
			app.DbRead,
			app.DbWrite,
			app.GrpcClientConn, cfg, app.Log)
	}
}

func GrpcRoute(
	grpcServer *grpc.Server,
	dbRead *pg.DB,
	dbWrite *pg.DB,
	clientConnection map[string]*grpc.ClientConn,
	config *config.Config,
	log *logrus.Logger,
) {
	// Handler initiation
	UserServer := userhandler.NewHandler(
		userusecase.NewUserUsecase(
			userrepo.NewUserRepo(dbRead), config,
		),
	)

	user.RegisterUserServiceServer(grpcServer, UserServer)
}
