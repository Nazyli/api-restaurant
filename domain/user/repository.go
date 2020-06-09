package user

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

// Repository Admin inteface
type Repository interface {
	Select(ctx context.Context, all bool, uid string, app int64) (users entity.Users, err error)
	GetByEmail(ctx context.Context, email string, app int64) (user *entity.User, err error)
	GetByID(ctx context.Context, all bool, uid string, id int64, app int64) (user *entity.User, err error)
	// Insert(ctx context.Context, user *entity.User) (err error)
	// Update(ctx context.Context, user *entity.User, userID string) (err error)
	// Delete(ctx context.Context, pid int64, id int64, userID string) (err error)
	// Check(ctx context.Context, userID string) (user *entity.User, err error)
}
