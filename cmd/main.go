package main

import (
	"log"
	"os"

	"github.com/himitery/fiber-todo/config"
	_ "github.com/himitery/fiber-todo/docs"
	"github.com/urfave/cli/v2"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := &cli.App{
		Name:    Name,
		Usage:   Usage,
		Version: Version,
		Flags:   Flags,
		Action: func(c *cli.Context) error {
			App(config.LoadConfigFile(
				c.String(EnvFlag.Name),
			)).Run()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
