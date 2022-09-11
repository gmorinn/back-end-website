package service

import (
	config "back-end-website/config"
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	db "back-end-website/internal"
	"back-end-website/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type IBlogService interface {
	CreateBlog(ctx context.Context, input *model.CreateBlogInput) (*model.Blog, error)
	UpdateBlog(ctx context.Context, input *model.UpdateBlogInput) (*model.Blog, error)
	DeleteBlog(ctx context.Context, id mypkg.UUID) (*bool, error)
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

func SqlBlogToGraphBlog(sqlBlog *db.Blog) *model.Blog {
	if sqlBlog == nil {
		return nil
	}
	return &model.Blog{
		ID:        mypkg.UUID(sqlBlog.ID.String()),
		UserID:    mypkg.UUID(sqlBlog.UserID.String()),
		Title:     sqlBlog.Title,
		Content:   sqlBlog.Content,
		Image:     sqlBlog.Image,
		CreatedAt: sqlBlog.CreatedAt,
		UpdatedAt: sqlBlog.UpdatedAt,
		DeletedAt: &sqlBlog.DeletedAt.Time,
	}
}

func (s *BlogService) CreateBlog(ctx context.Context, input *model.CreateBlogInput) (*model.Blog, error) {
	var res *model.Blog

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// create blog
		newBlog, err := q.CreateBlog(ctx, db.CreateBlogParams{
			Title:   input.Title,
			Content: input.Content,
			Image:   input.Image,
			UserID:  uuid.MustParse(string(input.UserID)),
		})
		if err != nil {
			return err
		}

		// get blog
		blog, err := q.GetBlogByID(ctx, uuid.MustParse(string(newBlog.ID.String())))
		if err != nil {
			return err
		}

		// convert to graphql model
		res = SqlBlogToGraphBlog(&blog)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_CREATE_BLOG", err)
	}
	return res, nil
}

func (s *BlogService) UpdateBlog(ctx context.Context, input *model.UpdateBlogInput) (*model.Blog, error) {
	var res *model.Blog

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// update blog
		if err := q.UpdateBlog(ctx, db.UpdateBlogParams{
			ID:      uuid.MustParse(string(input.ID)),
			Title:   input.Title,
			Content: input.Content,
			Image:   input.Image,
		}); err != nil {
			return err
		}

		// get blog
		blog, err := q.GetBlogByID(ctx, uuid.MustParse(string(input.ID)))
		if err != nil {
			return err
		}

		// convert to graphql model
		res = SqlBlogToGraphBlog(&blog)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_BLOG", err)
	}
	return res, nil
}

func (s *BlogService) DeleteBlog(ctx context.Context, id mypkg.UUID) (*bool, error) {
	var res bool = false

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if blog exists
		isBlog, err := q.CheckBlogByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		if !isBlog {
			return utils.ErrorResponse("BLOG_NOT_FOUND", errors.New("blog not found"))
		}
		if err := q.DeleteBlogByID(ctx, uuid.MustParse(string(id))); err != nil {
			res = false
			return err
		}
		res = true
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_DELETE_BLOG", err)
	}
	return &res, nil
}

func (s *BlogService) GetBlogs(ctx context.Context, limit int, offset int) ([]*model.Blog, error) {
	var res []*model.Blog

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if limit is valid
		if limit < 0 {
			return utils.ErrorResponse("INVALID_LIMIT", errors.New("limit must be greater than 0"))
		}
		// check if offset is valid
		if offset < 0 {
			return utils.ErrorResponse("INVALID_OFFSET", errors.New("offset must be greater than 0"))
		}
		blogs, err := q.GetAllBlog(ctx, db.GetAllBlogParams{
			Limit:    int32(limit),
			Offset:   int32(offset),
			TitleAsc: utils.FilterOrderBy("title", "asc", "TitleAsc"),
		})
		if err != nil {
			return err
		}

		// convert to graphql model
		for _, blog := range blogs {
			res = append(res, SqlBlogToGraphBlog(&blog))
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_BLOGS", err)
	}
	return res, nil
}

func (s *BlogService) GetBlog(ctx context.Context, id mypkg.UUID) (*model.Blog, error) {
	var res *model.Blog

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		blog, err := q.GetBlogByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		res = SqlBlogToGraphBlog(&blog)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_BLOG", err)
	}
	return res, err
}
