package service

import (
	config "back-end-website/config"
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	"context"
)

type IBlogService interface {
	CreateBlog(ctx context.Context, input *model.CreateBlogInput) (*model.Blog, error)
	UpdateBlog(ctx context.Context, input *model.UpdateBlogInput) (*model.Blog, error)
	DeleteBlog(ctx context.Context, id mypkg.UUID) (*model.Blog, error)
	GetBlogs(ctx context.Context, limit int, offset int) ([]*model.Blog, error)
	GetBlog(ctx context.Context, id mypkg.UUID) (*model.Blog, error)
}

type BlogService struct {
	server *config.Server
}

func NewBlogService(server *config.Server) *BlogService {
	return &BlogService{
		server: server,
	}
}


