package port

import (
	"github.com/google/uuid"
	"github.com/himitery/fiber-todo/internal/core/domain"
)

type TodoRepository interface {
	Find() ([]domain.Todo, error)
	FindOne(id uuid.UUID) (domain.Todo, error)
	Save(todo *domain.Todo) (domain.Todo, error)
	Update(todo *domain.Todo) (domain.Todo, error)
	Delete(id uuid.UUID) (domain.Todo, error)
}

type TodoUsecase interface {
	GetList() ([]domain.Todo, error)
	GetOne(id uuid.UUID) (domain.Todo, error)
	Create(req CreateTodoReq) (domain.Todo, error)
	Update(id uuid.UUID, req UpdateTodoReq) (domain.Todo, error)
	Delete(id uuid.UUID) (domain.Todo, error)
}

type CreateTodoReq struct {
	Title   string
	Content string
}

type UpdateTodoReq struct {
	Title   string
	Content string
}
