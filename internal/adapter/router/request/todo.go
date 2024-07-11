package request

import "github.com/himitery/fiber-todo/internal/core/port"

type CreateTodoReq struct {
	Title   string `json:"title" example:"sample title" validate:"required"`
	Content string `json:"content" example:"sample content"`
}

func (req CreateTodoReq) ToPortReq() port.CreateTodoReq {
	return port.CreateTodoReq{
		Title:   req.Title,
		Content: req.Content,
	}
}

type UpdateTodoReq struct {
	Title   string `json:"title" example:"sample title" validate:"required"`
	Content string `json:"content" example:"sample content" validate:"required"`
}

func (req UpdateTodoReq) ToPortReq() port.UpdateTodoReq {
	return port.UpdateTodoReq{
		Title:   req.Title,
		Content: req.Content,
	}
}
