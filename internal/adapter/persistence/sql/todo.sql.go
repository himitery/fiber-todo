// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: todo.sql

package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todo (
    title, content
) VALUES (
    $1, $2
)
RETURNING id, created_at, updated_at, title, content
`

type CreateTodoParams struct {
	Title   string
	Content string
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, arg.Title, arg.Content)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Content,
	)
	return i, err
}

const deleteOneTodo = `-- name: DeleteOneTodo :one
DELETE FROM todo
WHERE id = $1
RETURNING id, created_at, updated_at, title, content
`

func (q *Queries) DeleteOneTodo(ctx context.Context, id pgtype.UUID) (Todo, error) {
	row := q.db.QueryRow(ctx, deleteOneTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Content,
	)
	return i, err
}

const getTodoById = `-- name: GetTodoById :one
SELECT id, created_at, updated_at, title, content FROM todo
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTodoById(ctx context.Context, id pgtype.UUID) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Content,
	)
	return i, err
}

const getTodoMany = `-- name: GetTodoMany :many
SELECT id, created_at, updated_at, title, content FROM todo
ORDER BY updated_at DESC
`

func (q *Queries) GetTodoMany(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, getTodoMany)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Content,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todo
SET title = $2, content = $3
WHERE id = $1
RETURNING id, created_at, updated_at, title, content
`

type UpdateTodoParams struct {
	ID      pgtype.UUID
	Title   string
	Content string
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo, arg.ID, arg.Title, arg.Content)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Content,
	)
	return i, err
}
