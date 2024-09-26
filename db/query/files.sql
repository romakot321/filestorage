-- name: CreateFile :one
INSERT INTO files(
    filename,
    owner_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetFileById :one
SELECT * FROM files
WHERE filename = $1 LIMIT 1;

-- name: ListFiles :many
SELECT * FROM files
ORDER BY filename
LIMIT $1
OFFSET $2;

-- name: UpdateFile :one
UPDATE files
SET
owner_id = coalesce(sqlc.narg('owner_id'), owner_id),
updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
WHERE filename = sqlc.arg('filename')
RETURNING *;

-- name: DeleteFile :exec
DELETE FROM files 
WHERE filename = $1;
