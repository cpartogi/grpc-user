package application

import (
	"errors"
	"net"
	"strings"
	"user-service/config"

	"user-service/lib/pkg/database/db"

	"log"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func Setup(cfg *config.Config, c *cli.Context) (*ServiceApp, error) {
	app := new(ServiceApp)

	app.ServiceMode = c.String("mode")
	app.ServiceName = c.String("name")
	if err := runInit(
		initDatabase(cfg),
		initApp(cfg),
	)(app); err != nil {
		return app, err
	}

	return app, nil
}

func runInit(appFuncs ...func(*ServiceApp) error) func(*ServiceApp) error {
	return func(app *ServiceApp) error {
		for _, fn := range appFuncs {
			if e := fn(app); e != nil {
				return e
			}
		}
		return nil
	}
}

func initDatabase(cfg *config.Config) func(*ServiceApp) error {
	return func(app *ServiceApp) error {
		dbRead, err := db.NewPostgresDB(&cfg.UserDB)
		if err != nil {
			return err
		}

		app.DbRead = dbRead
		app.DbWrite = dbRead
		return nil
	}
}

func initGrpcServer() func(*ServiceApp) error {
	return func(app *ServiceApp) error {
		g := grpc.NewServer(
			grpc.ChainUnaryInterceptor(),
		)
		app.GrpcServer = g
		return nil
	}
}

func initApp(cfg *config.Config) func(*ServiceApp) error {
	return func(app *ServiceApp) error {
		switch app.ServiceMode {
		case "grpc":
			return initGrpcServer()(app)
		}
		return errors.New("unrecognized mode")
	}
}

func grpcRun(cfg *config.Config) func(*ServiceApp) error {
	return func(app *ServiceApp) error {
		log.Println("Starting service")
		log.Println(strings.ToUpper(cfg.Application.ServiceName) + " SERVICE running on port " + cfg.Application.ServerPort)
		defer func() {
			app.GrpcServer.GracefulStop()
			app.DbRead.Close()
			app.DbWrite.Close()

		}()
		lis, err := net.Listen("tcp", ":"+cfg.Application.ServerPort)
		if err != nil {
			log.Println(err)
			return err
		}
		if err := app.GrpcServer.Serve(lis); err != nil {
			log.Println("failed to serve ", err.Error())
			return err
		}
		return nil
	}
}
