package menu

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

// Repository Menu inteface
type Repository interface {
	Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (menus entity.Menues, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (menu *entity.Menu, err error)
	Insert(ctx context.Context, menu *entity.Menu) (err error)
	Update(ctx context.Context, isAdmin bool, menu *entity.Menu) (err error)
	Delete(ctx context.Context, isAdmin bool, menu *entity.Menu) (err error)
}
