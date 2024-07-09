-- name: GetTodoMany :many
SELECT * FROM todo
ORDER BY updated_at DESC;

-- name: GetTodoOne :one
SELECT * FROM todo
WHERE id = $1
LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todo (
    title, content
) VALUES (
    $1, $2
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
