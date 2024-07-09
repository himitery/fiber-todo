package application

import (
	"github.com/google/uuid"
	"github.com/himitery/fiber-todo/internal/core/domain"
	"github.com/himitery/fiber-todo/internal/core/port"
)

type TodoService struct {
	todoRepository port.TodoRepository
}

func NewTodoService(todoRepository port.TodoRepository) port.TodoUsecase {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (service TodoService) GetList() ([]domain.Todo, error) {
	todos, err := service.todoRepository.Find()
	if err != nil {
		return nil, &port.PortError{Code: 500, Message: err.Error()}
	}

	return todos, nil
}

func (service TodoService) GetOne(id uuid.UUID) (domain.Todo, error) {
	todo, err := service.todoRepository.FindOne(id)
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}

	return todo, nil
}

func (service TodoService) Create(req port.CreateTodoReq) (domain.Todo, error) {
	todo, err := service.todoRepository.Save(&domain.Todo{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 500, Message: err.Error()}
	}

	return todo, nil
}

func (service TodoService) Update(id uuid.UUID, req port.UpdateTodoReq) (domain.Todo, error) {
	todo, err := service.todoRepository.Update(&domain.Todo{
		Id:      id,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}

	return todo, nil
}

func (service TodoService) Delete(id uuid.UUID) (domain.Todo, error) {
	todo, err := service.todoRepository.Delete(id)
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}

	return todo, nil
}
