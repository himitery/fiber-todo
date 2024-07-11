package response

import (
	"time"

	"github.com/himitery/fiber-todo/internal/core/domain"
)

type TodoRes struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"Content"`
}

func NewTodoRes(todo domain.Todo) *TodoRes {
	return &TodoRes{
		Id:        todo.Id.String(),
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		Title:     todo.Title,
		Content:   todo.Content,
	}
}
