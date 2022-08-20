package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	"context"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id mypkg.UUID) (*model.User, error) {
	return r.UserService.GetUser(ctx, id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	return r.UserService.GetUsers(ctx, limit, offset)
}

// Blogs is the resolver for the blogs field.
func (r *queryResolver) Blogs(ctx context.Context, limit int, offset int) ([]*model.Blog, error) {
	return r.BlogService.GetBlogs(ctx, limit, offset)
}

// Blog is the resolver for the blog field.
func (r *queryResolver) Blog(ctx context.Context, id mypkg.UUID) (*model.Blog, error) {
	return r.BlogService.GetBlog(ctx, id)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
