package persistence

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/himitery/fiber-todo/config/database"
	"github.com/himitery/fiber-todo/internal/adapter/persistence/sql"
	"github.com/himitery/fiber-todo/internal/core/domain"
	"github.com/himitery/fiber-todo/internal/core/port"
	"github.com/himitery/fiber-todo/internal/utils"
)

type TodoPersistence struct {
	database *database.Database
}

func NewTodoPersistence(database *database.Database) port.TodoRepository {
	return &TodoPersistence{
		database: database,
	}
}

func (repo TodoPersistence) Find() ([]domain.Todo, error) {
	res, err := repo.database.Queries.GetTodoMany(repo.database.Context)

	return lo.Map(res, func(it sql.Todo, idx int) domain.Todo {
		return repo.mapToDomain(it)
	}), err
}

func (repo TodoPersistence) FindByAuthId(authId uuid.UUID) ([]domain.Todo, error) {
	res, err := repo.database.Queries.GetTodoByAuthId(repo.database.Context, utils.UuidToPGUuid(authId))

	return lo.Map(res, func(it sql.Todo, idx int) domain.Todo {
		return repo.mapToDomain(it)
	}), err
}

func (repo TodoPersistence) FindById(id uuid.UUID) (domain.Todo, error) {
	res, err := repo.database.Queries.GetTodoById(repo.database.Context, utils.UuidToPGUuid(id))

	return repo.mapToDomain(res), err
}

func (repo TodoPersistence) Save(todo *domain.Todo) (domain.Todo, error) {
	res, err := repo.database.Queries.CreateTodo(repo.database.Context, sql.CreateTodoParams{
		AuthID:  utils.UuidToPGUuid(todo.AuthId),
		Title:   todo.Title,
		Content: todo.Content,
	})

	return repo.mapToDomain(res), err
}

func (repo TodoPersistence) Update(todo *domain.Todo) (domain.Todo, error) {
	res, err := repo.database.Queries.UpdateTodo(repo.database.Context, sql.UpdateTodoParams{
		ID:      utils.UuidToPGUuid(todo.Id),
		Title:   todo.Title,
		Content: todo.Content,
	})

	return repo.mapToDomain(res), err
}

func (repo TodoPersistence) Delete(id uuid.UUID) (domain.Todo, error) {
	res, err := repo.database.Queries.DeleteOneTodo(repo.database.Context, utils.UuidToPGUuid(id))

	return repo.mapToDomain(res), err
}

func (repo TodoPersistence) mapToDomain(todo sql.Todo) domain.Todo {
	return domain.Todo{
		Id:        utils.PGUuidToUuid(todo.ID),
		CreatedAt: todo.CreatedAt.Time,
		UpdatedAt: todo.UpdatedAt.Time,
		AuthId:    utils.PGUuidToUuid(todo.AuthID),
		Title:     todo.Title,
		Content:   todo.Content,
	}
}
