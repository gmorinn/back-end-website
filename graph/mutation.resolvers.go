package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	"context"
)

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, input model.SigninInput) (*model.JWTResponse, error) {
	return r.Resolver.AuthService.Signin(ctx, &input)
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (*model.JWTResponse, error) {
	return r.AuthService.Signup(ctx, &input)
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, refreshToken mypkg.JWT) (*model.JWTResponse, error) {
	return r.AuthService.RefreshToken(ctx, &refreshToken)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return r.UserService.UpdateUser(ctx, &input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id mypkg.UUID) (*bool, error) {
	return r.UserService.DeleteUser(ctx, id)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.User, error) {
	return r.UserService.UpdateRole(ctx, role, id)
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file model.UploadInput) (*model.UploadResponse, error) {
	return r.FileService.UploadSingleFile(ctx, &file)
}

// CreateBlog is the resolver for the createBlog field.
func (r *mutationResolver) CreateBlog(ctx context.Context, input model.CreateBlogInput) (*model.Blog, error) {
	return r.BlogService.CreateBlog(ctx, &input)
}

// UpdateBlog is the resolver for the updateBlog field.
func (r *mutationResolver) UpdateBlog(ctx context.Context, input model.UpdateBlogInput) (*model.Blog, error) {
	return r.BlogService.UpdateBlog(ctx, &input)
}

// DeleteBlog is the resolver for the deleteBlog field.
func (r *mutationResolver) DeleteBlog(ctx context.Context, id mypkg.UUID) (*bool, error) {
	return r.BlogService.DeleteBlog(ctx, id)
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	return r.ProjectService.CreateProject(ctx, &input)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	return r.ProjectService.UpdateProject(ctx, &input)
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, id mypkg.UUID) (*bool, error) {
	return r.ProjectService.DeleteProject(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
