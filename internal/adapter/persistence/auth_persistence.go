package persistence

import (
	"github.com/google/uuid"
	"github.com/himitery/fiber-todo/config/database"
	"github.com/himitery/fiber-todo/internal/adapter/persistence/sql"
	"github.com/himitery/fiber-todo/internal/core/domain"
	"github.com/himitery/fiber-todo/internal/core/port"
	"github.com/himitery/fiber-todo/internal/utils"
)

type AuthPersistence struct {
	database *database.Database
}

func NewAuthPersistence(database *database.Database) port.AuthRepository {
	return &AuthPersistence{
		database: database,
	}
}

func (repo AuthPersistence) FindById(id uuid.UUID) (domain.Auth, error) {
	res, err := repo.database.Queries.GetAuthById(repo.database.Context, utils.UuidToPGUuid(id))

	return repo.mapToDomain(res), err
}

func (repo AuthPersistence) FindByEmail(email string) (domain.Auth, error) {
	res, err := repo.database.Queries.GetAuthByEmail(repo.database.Context, email)

	return repo.mapToDomain(res), err
}

func (repo AuthPersistence) Save(auth *domain.Auth) (domain.Auth, error) {
	res, err := repo.database.Queries.CreateAuth(repo.database.Context, sql.CreateAuthParams{
		Email:    auth.Email,
		Password: auth.Password,
		Username: auth.Username,
	})

	return repo.mapToDomain(res), err
}

func (repo AuthPersistence) UpdatePassword(auth *domain.Auth) (domain.Auth, error) {
	res, err := repo.database.Queries.UpdateAuthPassword(repo.database.Context, sql.UpdateAuthPasswordParams{
		ID:       utils.UuidToPGUuid(auth.Id),
		Password: auth.Password,
	})

	return repo.mapToDomain(res), err
}

func (repo AuthPersistence) mapToDomain(auth sql.Auth) domain.Auth {
	return domain.Auth{
		Id:        utils.PGUuidToUuid(auth.ID),
		CreatedAt: auth.CreatedAt.Time,
		UpdatedAt: auth.UpdatedAt.Time,
		Email:     auth.Email,
		Password:  auth.Password,
		Username:  auth.Username,
	}
}
