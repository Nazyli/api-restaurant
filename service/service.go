package service

import (
	"context"

	_user "github.com/nazyli/api-restaurant/domain/user"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
)

//DataResponse json

type Status struct {
	Code   int
	ErrMsg string
}
type svc struct {
	user _user.Repository
}

// New init sertvice
func New(user _user.Repository) Service {
	return &svc{
		user: user,
	}
}

type Service interface {
	SignIn(ctx context.Context, app int64, email, password string) (token *auth.Token, status Status)
	// User
	GetUserByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (user *entity.User, status Status)
	SelectUsers(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (users entity.Users, status Status)
	Insert(ctx context.Context, app int64, uid string, user *entity.User) (userData *entity.User, status Status)
	Update(ctx context.Context, app int64, id int64, isAdmin bool, uid string, user *entity.User) (userData *entity.User, status Status)
	Delete(ctx context.Context, app int64, id int64, isAdmin bool, uid string) (status Status)
}
