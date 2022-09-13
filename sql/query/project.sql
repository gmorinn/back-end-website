-- name: GetAllProject :many
SELECT * FROM projects
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN sqlc.arg('title_asc')::bool THEN title END asc,
  CASE WHEN sqlc.arg('title_desc')::bool THEN title END desc,
  CASE WHEN sqlc.arg('created_at_asc')::bool THEN created_at END asc,
  CASE WHEN sqlc.arg('created_at_desc')::bool THEN created_at END desc,
  CASE WHEN sqlc.arg('tag_asc')::bool THEN tag END asc,
  CASE WHEN sqlc.arg('tag_desc')::bool THEN tag END desc,
  CASE WHEN sqlc.arg('client_asc')::bool THEN client END asc,
  CASE WHEN sqlc.arg('client_desc')::bool THEN client END desc,
  CASE WHEN sqlc.arg('language_asc')::bool THEN language END asc,
  CASE WHEN sqlc.arg('language_desc')::bool THEN language END desc,
  CASE WHEN sqlc.arg('content_asc')::bool THEN content END asc,
  CASE WHEN sqlc.arg('content_desc')::bool THEN content END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteProjectByID :exec
UPDATE
    projects
SET
    deleted_at = NOW()
WHERE
    id = $1;

-- name: UpdateProject :exec
UPDATE
    projects
SET
    title = $2,
    content = $3,
    language = $4,
    url = $5,
    img_cover = $6,
    img_description = $7,
    user_id = $8,
    tag = $9,
    client = $10,
    updated_at = NOW()
WHERE
    id = $1
RETURNING *;

-- name: InsertProject :one
INSERT INTO projects (user_id, title, content, language, url, img_cover, img_description, tag, client)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: CheckProjectByID :one
SELECT EXISTS(
    SELECT * FROM projects
    WHERE id = $1
    AND deleted_at IS NULL
);