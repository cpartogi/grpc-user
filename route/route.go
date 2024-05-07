package route

import (
	"user-service/application"
	"user-service/config"

	"user-service/pb/user"

	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
)

func SetupRoute(cfg *config.Config, app *application.ServiceApp) {
	switch app.ServiceMode {
	case "grpc":
		GrpcRoute(app.GrpcServer,
			app.DbRead,
			app.DbWrite,
			app.GrpcClientConn)
	}
}

func GrpcRoute(
	grpcServer *grpc.Server,
	dbRead *pg.DB,
	dbWrite *pg.DB,
	clientConnection map[string]*grpc.ClientConn,
) {
	// Handler initiation
	UserServer := userhandler.NewHandler(
		userusecase.NewService(
			vehiclerepo.NewPostgresRepo(log, dbRead, dbWrite),
		), log,
	)

	user.RegisterUserServiceServer(grpcServer, UserServer)
}
