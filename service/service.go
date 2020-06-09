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
	SignIn(ctx context.Context, email, password string) (token *auth.Token, errMsg string, status int)
	GetByID(ctx context.Context, id int64) (user *entity.User, status int)
}
