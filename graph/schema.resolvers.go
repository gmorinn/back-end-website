package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	"context"
	"fmt"
)

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, input model.SigninInput) (*model.JWTResponse, error) {
	return r.AuthService.Signin(ctx, &input)
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (*model.JWTResponse, error) {
	return r.AuthService.Signup(ctx, &input)
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, refreshToken mypkg.JWT) (*model.JWTResponse, error) {
	return r.AuthService.RefreshToken(ctx, &refreshToken)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.GetUserResponse, error) {
	return r.UserService.UpdateRole(ctx, &role, &id)
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file model.UploadInput) (*model.UploadResponse, error) {
	return r.FileService.UploadSingleFile(ctx, &file)
}

// CreateBlog is the resolver for the createBlog field.
func (r *mutationResolver) CreateBlog(ctx context.Context, input model.CreateBlogInput) (*model.GetBlogResponse, error) {
	panic(fmt.Errorf("not implemented: CreateBlog - createBlog"))
}

// UpdateBlog is the resolver for the updateBlog field.
func (r *mutationResolver) UpdateBlog(ctx context.Context, input model.UpdateBlogInput) (*model.GetBlogResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateBlog - updateBlog"))
}

// DeleteBlog is the resolver for the deleteBlog field.
func (r *mutationResolver) DeleteBlog(ctx context.Context, id mypkg.UUID) (*model.GetBlogResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteBlog - deleteBlog"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id mypkg.UUID) (*model.GetUserResponse, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit int, offset int) (*model.GetUsersResponse, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Blogs is the resolver for the blogs field.
func (r *queryResolver) Blogs(ctx context.Context, limit int, offset int) (*model.GetBlogsResponse, error) {
	panic(fmt.Errorf("not implemented: Blogs - blogs"))
}

// Blog is the resolver for the blog field.
func (r *queryResolver) Blog(ctx context.Context, id mypkg.UUID) (*model.GetBlogResponse, error) {
	panic(fmt.Errorf("not implemented: Blog - blog"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
