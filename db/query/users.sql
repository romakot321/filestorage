-- name: CreateUser :one
INSERT INTO users(
    name,
    password_hash
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
name = coalesce(sqlc.narg('name'), name),
password_hash = coalesce(sqlc.narg('password_hash'), password_hash),
updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;
