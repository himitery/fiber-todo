package router

import (
	"github.com/gofiber/fiber/v3"
	server "github.com/himitery/fiber-todo/internal"
)

type HealthRouter struct {
	Router fiber.Router
}

func NewHealthRouter(httpServer *server.HttpServer) {
	router := HealthRouter{
		Router: httpServer.Server.Group("/api/health"),
	}

	router.Init()
}

func (router HealthRouter) Init() {
	router.Router.Get("/", router.ping)
}

// @Tags        Health
// @Summary		Ping
// @Produce		plain
// @Success		200		{object}	string
// @Router		/api/health/		[get]
func (router HealthRouter) ping(ctx fiber.Ctx) error {
	return ctx.Status(200).SendString("ok")
}
