package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/himitery/fiber-todo/config"
	server "github.com/himitery/fiber-todo/internal"
	"github.com/himitery/fiber-todo/internal/adapter/router/request"
	"github.com/himitery/fiber-todo/internal/adapter/router/response"
	"github.com/himitery/fiber-todo/internal/core/port"
)

type AuthRouter struct {
	Conf        *config.Config
	Router      fiber.Router
	AuthUsecase port.AuthUsecase
	JwtUsecase  port.JwtUsecase
}

func NewAuthRouter(
	conf *config.Config,
	httpServer *server.HttpServer,
	authUsecase port.AuthUsecase,
	jwtUsecase port.JwtUsecase,
) {
	router := AuthRouter{
		Conf:        conf,
		Router:      httpServer.Server.Group("/api/auth"),
		AuthUsecase: authUsecase,
		JwtUsecase:  jwtUsecase,
	}

	router.Init()
}

func (router AuthRouter) Init() {
	router.Router.Post("/login", router.signIn)
	router.Router.Post("/new", router.signUp)
	router.Router.Post("/renew", router.reIssue)

	// todo: udpate password with authentication
	router.Router.Use(JwtHandler(router.Conf))
	router.Router.Patch("/", router.updatePassword)
}

// @Tags        Auth
// @Summary		로그인
// @Accept		json
// @Produce		json
// @Param		request body 		request.SignInReq	true	"SignInReq"
// @Success		200		{object}	response.TokenRes
// @Failure	    500		{object} 	response.ErrorRes
// @Router		/api/auth/login		[post]
func (router AuthRouter) signIn(ctx fiber.Ctx) error {
	req := new(request.SignInReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	auth, err := router.AuthUsecase.SignIn(port.SignInReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	res, err := router.JwtUsecase.Generate(auth)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTokenRes(res))
}

// @Tags        Auth
// @Summary		회원가입
// @Accept		json
// @Produce		json
// @Param		request body 		request.SignUpReq	true	"SignUpReq"
// @Success		200		{object}	response.TokenRes
// @Failure	    500		{object} 	response.ErrorRes
// @Router		/api/auth/new		[post]
func (router AuthRouter) signUp(ctx fiber.Ctx) error {
	req := new(request.SignUpReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	auth, err := router.AuthUsecase.SignUp(port.SignUpReq{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		return err
	}

	res, err := router.JwtUsecase.Generate(auth)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTokenRes(res))
}

// @Tags        Auth
// @Summary		토큰 재발행
// @Accept		json
// @Produce		json
// @Param		request body 		request.ReIssueReq	true	"ReIssueReq"
// @Success		200		{object}	response.TokenRes
// @Failure	    500		{object} 	response.ErrorRes
// @Router		/api/auth/renew	[post]
func (router AuthRouter) reIssue(ctx fiber.Ctx) error {
	req := new(request.ReIssueReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	claims, err := router.JwtUsecase.Parse(req.RefreshToken)
	if err != nil {
		return err
	}

	auth, err := router.AuthUsecase.ReIssue(port.ReIssueReq{
		Id: uuid.MustParse(claims.Sub),
	})
	if err != nil {
		return err
	}

	res, err := router.JwtUsecase.Generate(auth)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTokenRes(res))
}

// @Tags        Auth
// @Summary		비밀번호 변경
// @Accept		json
// @Produce		json
// @Param		request body 		request.UpdatePasswordReq	true	"UpdatePasswordReq"
// @Success		200		{object}	response.TokenRes
// @Failure	    401		{object} 	response.ErrorRes
// @Router		/api/auth	[patch]
// @Security 	ApiKeyAuth
func (router AuthRouter) updatePassword(ctx fiber.Ctx) error {
	req := new(request.UpdatePasswordReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	id := ctx.Locals("auth").(string)

	auth, err := router.AuthUsecase.UpdatePassword(port.UpdatePasswordReq{
		Id:              uuid.MustParse(id),
		AlreadyPassword: req.AlreadyPassword,
		NewPassword:     req.NewPassword,
	})
	if err != nil {
		return err
	}

	res, err := router.JwtUsecase.Generate(auth)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.NewTokenRes(res))
}
