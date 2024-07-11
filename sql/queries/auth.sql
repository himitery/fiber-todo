-- name: GetAuthById :one
SELECT * FROM auth
WHERE id = $1
LIMIT 1;

-- name: GetAuthByEmail :one
SELECT * FROM auth
WHERE email = $1
LIMIT 1;

-- name: CreateAuth :one
INSERT INTO auth (
    email, password, username
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateAuthPassword :one
UPDATE auth
SET password = $2
WHERE id = $1
RETURNING *;
