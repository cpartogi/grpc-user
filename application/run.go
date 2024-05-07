package application

import (
	"errors"
	"user-service/config"
)

func (app *ServiceApp) Run(cfg *config.Config) error {

	switch app.ServiceMode {
	case "grpc":
		return grpcRun(cfg)(app)
	}
	return errors.New("unrecognized service mode")
}
