package main

import (
	"github.com/himitery/fiber-todo/config"
	"github.com/himitery/fiber-todo/config/database"
	server "github.com/himitery/fiber-todo/internal"
	"github.com/himitery/fiber-todo/internal/adapter/persistence"
	"github.com/himitery/fiber-todo/internal/adapter/router"
	"github.com/himitery/fiber-todo/internal/core/application"
	"go.uber.org/fx"
)

func App(conf *config.Config) *fx.App {
	return fx.New(
		fx.Provide(
			append(
				providers(),
				func() *config.Config { return conf },
			)...,
		),
		fx.Invoke(
			invokers()...,
		),
	)
}

func providers() []interface{} {
	return []interface{}{
		server.NewHttpServer,

		// Database
		database.NewDatabase,

		// Service
		application.NewTodoService,

		// Repository
		persistence.NewTodoPersistence,
	}
}

func invokers() []interface{} {
	return []interface{}{
		server.Api,

		// Router
		router.NewSwaggerRouter,
		router.NewHealthRouter,
		router.NewtodoRouter,
	}
}
