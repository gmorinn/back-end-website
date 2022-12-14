// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ProjectTag string

const (
	ProjectTagWebdevelopment ProjectTag = "webdevelopment"
	ProjectTagSocialmedia    ProjectTag = "socialmedia"
)

func (e *ProjectTag) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProjectTag(s)
	case string:
		*e = ProjectTag(s)
	default:
		return fmt.Errorf("unsupported scan type for ProjectTag: %T", src)
	}
	return nil
}

type Role string

const (
	RoleAdmin Role = "admin"
	RolePro   Role = "pro"
	RoleUser  Role = "user"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type Blog struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	UserID    uuid.UUID    `json:"user_id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	Image     string       `json:"image"`
}

type File struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	Name      sql.NullString `json:"name"`
	Url       sql.NullString `json:"url"`
	Mime      sql.NullString `json:"mime"`
	Size      sql.NullInt64  `json:"size"`
}

type Project struct {
	ID             uuid.UUID      `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      sql.NullTime   `json:"deleted_at"`
	UserID         uuid.UUID      `json:"user_id"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	ImgCover       string         `json:"img_cover"`
	ImgDescription string         `json:"img_description"`
	Language       sql.NullString `json:"language"`
	Client         sql.NullString `json:"client"`
	Tag            ProjectTag     `json:"tag"`
	Url            string         `json:"url"`
}

type RefreshToken struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Token     string       `json:"token"`
	Ip        string       `json:"ip"`
	UserAgent string       `json:"user_agent"`
	ExpirOn   time.Time    `json:"expir_on"`
	UserID    uuid.UUID    `json:"user_id"`
}

type User struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Firstname sql.NullString `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Role      Role           `json:"role"`
}
