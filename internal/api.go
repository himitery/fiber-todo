package server

import (
	"context"

	"github.com/gofiber/fiber/v3/log"
	"github.com/himitery/fiber-todo/config"
	"go.uber.org/fx"
)

func Api(
	conf *config.Config,
	lifecycle fx.Lifecycle,
	httpServer *HttpServer,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := httpServer.Listen(conf.Host + ":" + conf.Port); err != nil {
					log.Error(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown()
		},
	})
}
