// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CheckBlogByID(ctx context.Context, id uuid.UUID) (bool, error)
	CheckEmailExist(ctx context.Context, email string) (bool, error)
	CheckProjectByID(ctx context.Context, id uuid.UUID) (bool, error)
	CheckUserByID(ctx context.Context, id uuid.UUID) (bool, error)
	CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error)
	CreateFile(ctx context.Context, arg CreateFileParams) (CreateFileRow, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) error
	DeleteBlogByID(ctx context.Context, id uuid.UUID) error
	DeleteFile(ctx context.Context, url sql.NullString) error
	DeleteOldRefreshToken(ctx context.Context) error
	DeleteProjectByID(ctx context.Context, id uuid.UUID) error
	DeleteRefreshToken(ctx context.Context, id uuid.UUID) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	GetAllBlog(ctx context.Context, arg GetAllBlogParams) ([]Blog, error)
	GetAllProject(ctx context.Context, arg GetAllProjectParams) ([]Project, error)
	GetAllUser(ctx context.Context, arg GetAllUserParams) ([]User, error)
	GetBlogByID(ctx context.Context, id uuid.UUID) (Blog, error)
	GetFileByURL(ctx context.Context, url sql.NullString) (File, error)
	GetOldRefreshToken(ctx context.Context) (RefreshToken, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (Project, error)
	GetRefreshToken(ctx context.Context, token string) (GetRefreshTokenRow, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	InsertProject(ctx context.Context, arg InsertProjectParams) (Project, error)
	ListRefreshTokenByUserID(ctx context.Context, arg ListRefreshTokenByUserIDParams) ([]RefreshToken, error)
	LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserRow, error)
	Signup(ctx context.Context, arg SignupParams) (User, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) error
	UpdateProject(ctx context.Context, arg UpdateProjectParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
