// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: auth.sql

package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAuth = `-- name: CreateAuth :one
INSERT INTO auth (
    email, password, username
) VALUES (
    $1, $2, $3
)
RETURNING id, created_at, updated_at, email, password, username
`

type CreateAuthParams struct {
	Email    string
	Password string
	Username string
}

func (q *Queries) CreateAuth(ctx context.Context, arg CreateAuthParams) (Auth, error) {
	row := q.db.QueryRow(ctx, createAuth, arg.Email, arg.Password, arg.Username)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.Username,
	)
	return i, err
}

const getAuthByEmail = `-- name: GetAuthByEmail :one
SELECT id, created_at, updated_at, email, password, username FROM auth
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetAuthByEmail(ctx context.Context, email string) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuthByEmail, email)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.Username,
	)
	return i, err
}

const getAuthById = `-- name: GetAuthById :one
SELECT id, created_at, updated_at, email, password, username FROM auth
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAuthById(ctx context.Context, id pgtype.UUID) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuthById, id)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.Username,
	)
	return i, err
}

const updateAuthPassword = `-- name: UpdateAuthPassword :one
UPDATE auth
SET password = $2
WHERE id = $1
RETURNING id, created_at, updated_at, email, password, username
`

type UpdateAuthPasswordParams struct {
	ID       pgtype.UUID
	Password string
}

func (q *Queries) UpdateAuthPassword(ctx context.Context, arg UpdateAuthPasswordParams) (Auth, error) {
	row := q.db.QueryRow(ctx, updateAuthPassword, arg.ID, arg.Password)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.Username,
	)
	return i, err
}