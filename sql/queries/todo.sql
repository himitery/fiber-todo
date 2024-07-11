-- name: GetTodoMany :many
SELECT * FROM todo
ORDER BY updated_at DESC;

-- name: GetTodoByAuthId :many
SELECT * FROM todo
WHERE auth_id = $1
ORDER BY updated_at DESC;

-- name: GetTodoById :one
SELECT * FROM todo
WHERE id = $1
LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todo (
    auth_id, title, content
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todo
SET title = $2, content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteOneTodo :one
DELETE FROM todo
WHERE id = $1
RETURNING *;
