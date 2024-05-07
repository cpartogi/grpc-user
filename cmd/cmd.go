package cmd

import (
	"user-service/config"

	"user-service/application"

	"github.com/urfave/cli"
)

func runCommand(cfg *config.Config) func(*cli.Context) error {
	return func(c *cli.Context) error {
		app, err := application.Setup(cfg, c)
		if err != nil {
			return err
		}

		route.SetupRoute(cfg, app)
		return app.Run(cfg)
	}
}
