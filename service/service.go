package service

import (
	"context"

	_position "github.com/nazyli/api-restaurant/domain/position"
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
	AppID    int64
	user     _user.Repository
	position _position.Repository
}

// New init service
func New(AppID int64, user _user.Repository, position _position.Repository) Service {
	return &svc{
		AppID:    AppID,
		user:     user,
		position: position,
	}
}

type Service interface {
	SignIn(ctx context.Context, email, password string) (token *auth.Token, status Status)
	// User
	GetUserByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (user *entity.User, status Status)
	SelectUsers(ctx context.Context, all bool, isAdmin bool, uid string) (users entity.Users, status Status)
	InsertUser(ctx context.Context, user *entity.User) (userData *entity.User, status Status)
	UpdateUser(ctx context.Context, id int64, isAdmin bool, uid string, user *entity.User) (userData *entity.User, status Status)
	DeleteUser(ctx context.Context, id int64, isAdmin bool, uid string) (status Status)

	//Position
	SelectPosition(ctx context.Context, all bool) (positions entity.Positions, status Status)
	GetPositionByID(ctx context.Context, id int64, all bool) (position *entity.Position, status Status)
	InsertPosition(ctx context.Context, position *entity.Position) (positionData *entity.Position, status Status)
	UpdatePosition(ctx context.Context, position *entity.Position) (positionData *entity.Position, status Status)
	DeletePosition(ctx context.Context, position *entity.Position) (status Status)
}
