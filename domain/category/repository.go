package category

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

// Repository category inteface
type Repository interface {
	Select(ctx context.Context, app int64, all bool) (categorys entity.Categorys, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool) (category *entity.Category, err error)
	Insert(ctx context.Context, c *entity.Category) (err error)
	Update(ctx context.Context, c *entity.Category) (err error)
	Delete(ctx context.Context, c *entity.Category) (err error)
}
