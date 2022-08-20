-- name: GetAllBlog :many
SELECT * FROM blogs
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN sqlc.arg('title_asc')::bool THEN title END asc,
  CASE WHEN sqlc.arg('title_desc')::bool THEN title END desc,
  CASE WHEN sqlc.arg('created_at_asc')::bool THEN created_at END asc,
  CASE WHEN sqlc.arg('created_at_desc')::bool THEN created_at END desc,
  CASE WHEN sqlc.arg('content_asc')::bool THEN content END asc,
  CASE WHEN sqlc.arg('content_desc')::bool THEN content END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetBlogByID :one
SELECT * FROM blogs
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteBlogByID :exec
UPDATE
    blogs
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: UpdateBlog :exec
UPDATE 
    blogs
SET
    title = $2,
    content = $3,
    image = $4,
    updated_at = NOW()
WHERE
    id = $1;

-- name: CreateBlog :exec
INSERT INTO blogs (user_id, title, content, image)
VALUES ($1, $2, $3, $4);