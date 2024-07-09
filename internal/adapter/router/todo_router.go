package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	server "github.com/himitery/fiber-todo/internal"
	"github.com/himitery/fiber-todo/internal/adapter/router/request"
	"github.com/himitery/fiber-todo/internal/core/port"
)

type TodoRouter struct {
	Router      fiber.Router
	TodoUsecase port.TodoUsecase
}

func NewtodoRouter(httpServer *server.HttpServer, todoUsecase port.TodoUsecase) {
	todoRouter := TodoRouter{
		Router:      httpServer.Server.Group("/api/todo"),
		TodoUsecase: todoUsecase,
	}

	todoRouter.Init()
}

func (router TodoRouter) Init() {
	router.Router.Get("/list", router.GetList)
	router.Router.Get("/:id", router.GetById)
	router.Router.Post("/new", router.Create)
	router.Router.Patch("/:id", router.Update)
	router.Router.Delete("/:id", router.Delete)
}

// @Tags        Todo
// @Summary		Todo 목록 조회
// @Produce		json
// @Success		200		{object}	[]domain.Todo
// @Failure	    500		{object}	response.ErrorRes
// @Router		/api/todo/list		[get]
func (router TodoRouter) GetList(ctx fiber.Ctx) error {
	res, err := router.TodoUsecase.GetList()
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Tags        Todo
// @Summary		Todo 조회
// @Produce		json
// @Param       id		path		string				true	"id"
// @Success		200		{object}	domain.Todo
// @Failure	    404		{object} 	response.ErrorRes
// @Router		/api/todo/{id}		[get]
func (router TodoRouter) GetById(ctx fiber.Ctx) error {
	res, err := router.TodoUsecase.GetOne(uuid.MustParse(ctx.Params("id")))
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Tags        Todo
// @Summary		Todo 생성
// @Accept		json
// @Produce		json
// @Param		request body 		request.CreateTodoReq	true	"CreateTodoReq"
// @Success		200		{object}	domain.Todo
// @Failure	    500		{object} 	response.ErrorRes
// @Router		/api/todo/new		[post]
func (router TodoRouter) Create(ctx fiber.Ctx) error {
	req := new(request.CreateTodoReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	res, err := router.TodoUsecase.Create(req.ToPortReq())
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Tags        Todo
// @Summary		Todo 수정
// @Accept		json
// @Produce		json
// @Param       id		path		string					true	"id"
// @Param		request body 		request.CreateTodoReq	true	"UpdateTodoReq"
// @Success		200		{object}	domain.Todo
// @Failure	    404		{object} 	response.ErrorRes
// @Router		/api/todo/{id}		[patch]
func (router TodoRouter) Update(ctx fiber.Ctx) error {
	req := new(request.UpdateTodoReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	res, err := router.TodoUsecase.Update(uuid.MustParse(ctx.Params("id")), req.ToPortReq())
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Tags        Todo
// @Summary		Todo 삭제
// @Produce		json
// @Param       id		path		string				true	"id"
// @Success		200		{object}	domain.Todo
// @Failure	    404		{object} 	response.ErrorRes
// @Router		/api/todo/{id}		[delete]
func (router TodoRouter) Delete(ctx fiber.Ctx) error {
	res, err := router.TodoUsecase.Delete(uuid.MustParse(ctx.Params("id")))
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}
