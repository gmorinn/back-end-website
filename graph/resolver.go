package graph

import (
	"back-end-website/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService service.IUserService
	BlogService service.IBlogService
	AuthService service.IAuthService
	FileService service.IFileService
}
