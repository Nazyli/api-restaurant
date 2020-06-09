package service

import (
	"context"

	_user "github.com/nazyli/api-restaurant/domain/user"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
)

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
	SignIn(ctx context.Context, email, password string, app int64) (token *auth.Token, errMsg string, status int)
	// User
	GetUserByID(ctx context.Context, all bool, uid string, id int64, app int64) (user *entity.User, status int)
	SelectUsers(ctx context.Context, all bool, uid string, app int64) (users entity.Users, status int)
}
