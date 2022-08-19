// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: students.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const countStudent = `-- name: CountStudent :one
SELECT COUNT(*) FROM students
WHERE deleted_at IS NULL
`

func (q *Queries) CountStudent(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countStudent)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteStudentByID = `-- name: DeleteStudentByID :exec
UPDATE
    students
SET
    deleted_at = NOW()
WHERE 
    id = $1
`

func (q *Queries) DeleteStudentByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteStudentByID, id)
	return err
}

const getStudentByID = `-- name: GetStudentByID :one
SELECT id, created_at, updated_at, deleted_at, email, password, name, role FROM students
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetStudentByID(ctx context.Context, id uuid.UUID) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByID, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const insertStudent = `-- name: InsertStudent :one
INSERT INTO students (email, name)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, deleted_at, email, password, name, role
`

type InsertStudentParams struct {
	Email string         `json:"email"`
	Name  sql.NullString `json:"name"`
}

func (q *Queries) InsertStudent(ctx context.Context, arg InsertStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, insertStudent, arg.Email, arg.Name)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const liststudents = `-- name: Liststudents :many
SELECT id, created_at, updated_at, deleted_at, email, password, name, role FROM students
WHERE deleted_at IS NULL
ORDER BY name
LIMIT $1 OFFSET $2
`

type ListstudentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) Liststudents(ctx context.Context, arg ListstudentsParams) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, liststudents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Student{}
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Email,
			&i.Password,
			&i.Name,
			&i.Role,
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

const updateRoleStudent = `-- name: UpdateRoleStudent :exec
UPDATE students
SET role = $1,
    updated_at = NOW()
WHERE id = $2
`

type UpdateRoleStudentParams struct {
	Role Role      `json:"role"`
	ID   uuid.UUID `json:"id"`
}

func (q *Queries) UpdateRoleStudent(ctx context.Context, arg UpdateRoleStudentParams) error {
	_, err := q.db.ExecContext(ctx, updateRoleStudent, arg.Role, arg.ID)
	return err
}

const updateStudent = `-- name: UpdateStudent :exec
UPDATE students
SET name = $1,
    email = $2,
    updated_at = NOW()
WHERE id = $3
`

type UpdateStudentParams struct {
	Name  sql.NullString `json:"name"`
	Email string         `json:"email"`
	ID    uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateStudent(ctx context.Context, arg UpdateStudentParams) error {
	_, err := q.db.ExecContext(ctx, updateStudent, arg.Name, arg.Email, arg.ID)
	return err
}
