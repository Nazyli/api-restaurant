package user

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

// Repository Admin inteface
type Repository interface {
	Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (users entity.Users, err error)
	GetByEmail(ctx context.Context, app int64, email string) (user *entity.User, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (user *entity.User, err error)
	Insert(ctx context.Context, uid string, user *entity.User) (err error)
	Update(ctx context.Context, isAdmin bool, user *entity.User) (err error)
	Delete(ctx context.Context, isAdmin bool, user *entity.User) (err error)
}
