// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"back-end-website/graph/mypkg"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// All fields that represent a blog
type BlogResponse struct {
	UserID    mypkg.UUID `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Image     string     `json:"image"`
}

type CreateBlogInput struct {
	// title of the blog (required)
	Title string `json:"title"`
	// content of the blog (required)
	Content string `json:"content"`
	// image of the blog (required)
	Image string `json:"image"`
}

// Response when you get a blog
type GetBlogResponse struct {
	// if the request was successful or not, return always a value
	Success bool `json:"success"`
	// return the blog if the request was successful
	User *BlogResponse `json:"user"`
}

// Response when you get many blogs
type GetBlogsResponse struct {
	// if the request was successful or not, return always a value
	Success bool `json:"success"`
	// return an array of blog if the request was successful or null if there is an error or no users
	Users []*BlogResponse `json:"users"`
}

// Response when you get a user
type GetUserResponse struct {
	// if the request was successful or not, return always a value
	Success bool `json:"success"`
	// return the user if the request was successful
	User *UserResponse `json:"user"`
}

// Response when you get many users
type GetUsersResponse struct {
	// if the request was successful or not, return always a value
	Success bool `json:"success"`
	// return an array of user if the request was successful or null if there is an error or no users
	Users []*UserResponse `json:"users"`
}

type JWTResponse struct {
	// jwt token for user to authenticate, contains user id, role and expiry
	AccessToken mypkg.JWT `json:"access_token"`
	// use to refresh the access token
	RefreshToken mypkg.JWT `json:"refresh_token"`
	// true if the user can connect or false if not
	Success bool `json:"success"`
}

// All fields that represent a project
type ProjectResponse struct {
	UserID    mypkg.UUID `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Image     string     `json:"image"`
	Language  string     `json:"language"`
	URL       mypkg.URL  `json:"url"`
}

type SigninInput struct {
	// email of the user
	Email mypkg.Email `json:"email"`
	// password of the user
	Password string `json:"password"`
}

type SignupInput struct {
	// email of the user
	Email mypkg.Email `json:"email"`
	// password of the user
	Password string `json:"password"`
	// confirm password of the user
	ConfirmPassword string `json:"confirm_password"`
	// firstname of the user
	Firstname string `json:"firstname"`
	// lastname of the user
	Lastname string `json:"lastname"`
}

// payload send when you add a blog
type UpdateBlogInput struct {
	// title of the blog (required)
	Title string `json:"title"`
	// content of the blog (required)
	Content string `json:"content"`
	// id of the blog (required)
	ID mypkg.UUID `json:"id"`
	// image of the blog (required)
	Image string `json:"image"`
}

// payload send when you add a user
type UpdateUserInput struct {
	// name of the user (required)
	Name string `json:"name"`
	// email of the user (required)
	Email mypkg.Email `json:"email"`
	// firstname of the user (required)
	Firstname string `json:"firstname"`
	// lastname of the user (required)
	Lastname string `json:"lastname"`
	// id of the user (required)
	ID mypkg.UUID `json:"id"`
	// role of the user (required)
	Role UserType `json:"role"`
}

type UploadInput struct {
	// The file to upload
	File graphql.Upload `json:"file"`
	// width of the image if it needs to be resized
	Width *int `json:"width"`
	// height of the image if it needs to be resized
	Height *int `json:"height"`
}

// The `File` type, represents the response of uploading a file.
type UploadResponse struct {
	Name    string `json:"name"`
	Size    int    `json:"size"`
	URL     string `json:"url"`
	Success bool   `json:"success"`
}

// All fields that represent a user
type UserResponse struct {
	Firstname string      `json:"firstname"`
	Lastname  string      `json:"lastname"`
	Email     mypkg.Email `json:"email"`
	ID        mypkg.UUID  `json:"id"`
	Role      UserType    `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type UserType string

const (
	// User can have access to all data
	UserTypeAdmin UserType = "ADMIN"
	// User can access specific data but not all
	UserTypePro UserType = "PRO"
	// User can only see their own data
	UserTypeUser UserType = "USER"
)

var AllUserType = []UserType{
	UserTypeAdmin,
	UserTypePro,
	UserTypeUser,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeAdmin, UserTypePro, UserTypeUser:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
