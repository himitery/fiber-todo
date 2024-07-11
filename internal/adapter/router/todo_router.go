package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/himitery/fiber-todo/config"
	server "github.com/himitery/fiber-todo/internal"
	"github.com/himitery/fiber-todo/internal/adapter/router/request"
	"github.com/himitery/fiber-todo/internal/adapter/router/response"
	"github.com/himitery/fiber-todo/internal/core/domain"
	"github.com/himitery/fiber-todo/internal/core/port"
	"github.com/samber/lo"
)

type TodoRouter struct {
	Conf        *config.Config
	Router      fiber.Router
	TodoUsecase port.TodoUsecase
}

func NewtodoRouter(conf *config.Config, httpServer *server.HttpServer, todoUsecase port.TodoUsecase) {
	todoRouter := TodoRouter{
		Conf:        conf,
		Router:      httpServer.Server.Group("/api/todo"),
		TodoUsecase: todoUsecase,
	}

	todoRouter.Init()
}

func (router TodoRouter) Init() {
	router.Router.Use(JwtHandler(router.Conf))

	router.Router.Get("/list", router.getList)
	router.Router.Get("/:id", router.getById)
	router.Router.Post("/new", router.create)
	router.Router.Patch("/:id", router.update)
	router.Router.Delete("/:id", router.delete)
}

// @Tags        Todo
// @Summary		Todo 목록 조회
// @Produce		json
// @Success		200		{object}	[]response.TodoRes
// @Failure	    500		{object}	response.ErrorRes
// @Router		/api/todo/list		[get]
// @Security 	ApiKeyAuth
func (router TodoRouter) getList(ctx fiber.Ctx) error {
	authId := ctx.Locals("auth").(string)

	res, err := router.TodoUsecase.GetList(uuid.MustParse(authId))
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(lo.Map(res, func(todo domain.Todo, _ int) response.TodoRes {
		return *response.NewTodoRes(todo)
	}))
}

// @Tags        Todo
// @Summary		Todo 조회
// @Produce		json
// @Param       id		path		string				true	"id"
// @Success		200		{object}	response.TodoRes
// @Failure	    404		{object} 	response.ErrorRes
// @Router		/api/todo/{id}		[get]
// @Security 	ApiKeyAuth
func (router TodoRouter) getById(ctx fiber.Ctx) error {
	authId := ctx.Locals("auth").(string)

	res, err := router.TodoUsecase.GetOne(uuid.MustParse(authId), uuid.MustParse(ctx.Params("id")))
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTodoRes(res))
}

// @Tags        Todo
// @Summary		Todo 생성
// @Accept		json
// @Produce		json
// @Param		request body 		request.CreateTodoReq	true	"CreateTodoReq"
// @Success		200		{object}	response.TodoRes
// @Failure	    500		{object} 	response.ErrorRes
// @Router		/api/todo/new		[post]
// @Security 	ApiKeyAuth
func (router TodoRouter) create(ctx fiber.Ctx) error {
	authId := ctx.Locals("auth").(string)

	req := new(request.CreateTodoReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	res, err := router.TodoUsecase.Create(uuid.MustParse(authId), req.ToPortReq())
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTodoRes(res))
}

// @Tags        Todo
// @Summary		Todo 수정
// @Accept		json
// @Produce		json
// @Param       id		path		string					true	"id"
// @Param		request body 		request.CreateTodoReq	true	"UpdateTodoReq"
// @Success		200		{object}	response.TodoRes
// @Failure	    404		{object} 	response.ErrorRes
// @Router		/api/todo/{id}		[patch]
// @Security 	ApiKeyAuth
func (router TodoRouter) update(ctx fiber.Ctx) error {
	authId := ctx.Locals("auth").(string)

	req := new(request.UpdateTodoReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	res, err := router.TodoUsecase.Update(uuid.MustParse(authId), uuid.MustParse(ctx.Params("id")), req.ToPortReq())
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTodoRes(res))
}

// @Tags        Todo
// @Summary		Todo 삭제
// @Produce		json
// @Param       id		path		string				true	"id"
// @Success		200		{object}	response.TodoRes
// @Failure	    404		{object} 	response.ErrorRes
// @Router		/api/todo/{id}		[delete]
// @Security 	ApiKeyAuth
func (router TodoRouter) delete(ctx fiber.Ctx) error {
	authId := ctx.Locals("auth").(string)

	res, err := router.TodoUsecase.Delete(uuid.MustParse(authId), uuid.MustParse(ctx.Params("id")))
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTodoRes(res))
}
