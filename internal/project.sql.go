// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: project.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const checkProjectByID = `-- name: CheckProjectByID :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, user_id, title, content, img_cover, img_description, language, client, tag, url FROM projects
    WHERE id = $1
    AND deleted_at IS NULL
)
`

func (q *Queries) CheckProjectByID(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkProjectByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const deleteProjectByID = `-- name: DeleteProjectByID :exec
UPDATE
    projects
SET
    deleted_at = NOW()
WHERE
    id = $1
`

func (q *Queries) DeleteProjectByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProjectByID, id)
	return err
}

const getAllProject = `-- name: GetAllProject :many
SELECT id, created_at, updated_at, deleted_at, user_id, title, content, img_cover, img_description, language, client, tag, url FROM projects
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN $1::bool THEN title END asc,
  CASE WHEN $2::bool THEN title END desc,
  CASE WHEN $3::bool THEN created_at END asc,
  CASE WHEN $4::bool THEN created_at END desc,
  CASE WHEN $5::bool THEN tag END asc,
  CASE WHEN $6::bool THEN tag END desc,
  CASE WHEN $7::bool THEN client END asc,
  CASE WHEN $8::bool THEN client END desc,
  CASE WHEN $9::bool THEN language END asc,
  CASE WHEN $10::bool THEN language END desc,
  CASE WHEN $11::bool THEN content END asc,
  CASE WHEN $12::bool THEN content END desc
LIMIT $14 OFFSET $13
`

type GetAllProjectParams struct {
	TitleAsc      bool  `json:"title_asc"`
	TitleDesc     bool  `json:"title_desc"`
	CreatedAtAsc  bool  `json:"created_at_asc"`
	CreatedAtDesc bool  `json:"created_at_desc"`
	TagAsc        bool  `json:"tag_asc"`
	TagDesc       bool  `json:"tag_desc"`
	ClientAsc     bool  `json:"client_asc"`
	ClientDesc    bool  `json:"client_desc"`
	LanguageAsc   bool  `json:"language_asc"`
	LanguageDesc  bool  `json:"language_desc"`
	ContentAsc    bool  `json:"content_asc"`
	ContentDesc   bool  `json:"content_desc"`
	Offset        int32 `json:"offset"`
	Limit         int32 `json:"limit"`
}

func (q *Queries) GetAllProject(ctx context.Context, arg GetAllProjectParams) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getAllProject,
		arg.TitleAsc,
		arg.TitleDesc,
		arg.CreatedAtAsc,
		arg.CreatedAtDesc,
		arg.TagAsc,
		arg.TagDesc,
		arg.ClientAsc,
		arg.ClientDesc,
		arg.LanguageAsc,
		arg.LanguageDesc,
		arg.ContentAsc,
		arg.ContentDesc,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.ImgCover,
			&i.ImgDescription,
			&i.Language,
			&i.Client,
			&i.Tag,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, created_at, updated_at, deleted_at, user_id, title, content, img_cover, img_description, language, client, tag, url FROM projects
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetProjectByID(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.ImgCover,
		&i.ImgDescription,
		&i.Language,
		&i.Client,
		&i.Tag,
		&i.Url,
	)
	return i, err
}

const insertProject = `-- name: InsertProject :one
INSERT INTO projects (user_id, title, content, language, url, img_cover, img_description, tag, client)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, created_at, updated_at, deleted_at, user_id, title, content, img_cover, img_description, language, client, tag, url
`

type InsertProjectParams struct {
	UserID         uuid.UUID      `json:"user_id"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	Language       sql.NullString `json:"language"`
	Url            string         `json:"url"`
	ImgCover       string         `json:"img_cover"`
	ImgDescription string         `json:"img_description"`
	Tag            ProjectTag     `json:"tag"`
	Client         sql.NullString `json:"client"`
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, insertProject,
		arg.UserID,
		arg.Title,
		arg.Content,
		arg.Language,
		arg.Url,
		arg.ImgCover,
		arg.ImgDescription,
		arg.Tag,
		arg.Client,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.ImgCover,
		&i.ImgDescription,
		&i.Language,
		&i.Client,
		&i.Tag,
		&i.Url,
	)
	return i, err
}

const updateProject = `-- name: UpdateProject :exec
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
RETURNING id, created_at, updated_at, deleted_at, user_id, title, content, img_cover, img_description, language, client, tag, url
`

type UpdateProjectParams struct {
	ID             uuid.UUID      `json:"id"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	Language       sql.NullString `json:"language"`
	Url            string         `json:"url"`
	ImgCover       string         `json:"img_cover"`
	ImgDescription string         `json:"img_description"`
	UserID         uuid.UUID      `json:"user_id"`
	Tag            ProjectTag     `json:"tag"`
	Client         sql.NullString `json:"client"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) error {
	_, err := q.db.ExecContext(ctx, updateProject,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.Language,
		arg.Url,
		arg.ImgCover,
		arg.ImgDescription,
		arg.UserID,
		arg.Tag,
		arg.Client,
	)
	return err
}
