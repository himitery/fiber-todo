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

func (service TodoService) GetList(authId uuid.UUID) ([]domain.Todo, error) {
	todos, err := service.todoRepository.FindByAuthId(authId)
	if err != nil {
		return nil, &port.PortError{Code: 500, Message: err.Error()}
	}

	return todos, nil
}

func (service TodoService) GetOne(authId uuid.UUID, id uuid.UUID) (domain.Todo, error) {
	todo, err := service.todoRepository.FindById(id)
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}
	if todo.AuthId != authId {
		return domain.Todo{}, &port.PortError{Code: 403, Message: "권한이 없습니다."}
	}

	return todo, nil
}

func (service TodoService) Create(authId uuid.UUID, req port.CreateTodoReq) (domain.Todo, error) {
	todo, err := service.todoRepository.Save(&domain.Todo{
		Title:   req.Title,
		Content: req.Content,
		AuthId:  authId,
	})
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 500, Message: err.Error()}
	}

	return todo, nil
}

func (service TodoService) Update(authId uuid.UUID, id uuid.UUID, req port.UpdateTodoReq) (domain.Todo, error) {
	todo, err := service.todoRepository.FindById(id)
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}
	if todo.AuthId != authId {
		return domain.Todo{}, &port.PortError{Code: 403, Message: "권한이 없습니다."}
	}

	todo, err = service.todoRepository.Update(&domain.Todo{
		Id:      id,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}

	return todo, nil
}

func (service TodoService) Delete(authId uuid.UUID, id uuid.UUID) (domain.Todo, error) {
	todo, err := service.todoRepository.FindById(id)
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 404, Message: "아이템을 찾을 수 없습니다."}
	}
	if todo.AuthId != authId {
		return domain.Todo{}, &port.PortError{Code: 403, Message: "권한이 없습니다."}
	}

	todo, err = service.todoRepository.Delete(todo.Id)
	if err != nil {
		return domain.Todo{}, &port.PortError{Code: 500, Message: err.Error()}
	}

	return todo, nil
}
