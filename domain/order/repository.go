package order

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

type Repository interface {
	// Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (orders entity.Orders, err error)
	GetByInv(ctx context.Context, app int64, inv string, all bool, isAdmin bool, uid string) (order *entity.Order, err error)
	Insert(ctx context.Context, order *entity.Order) (err error)
	Update(ctx context.Context, isAdmin bool, order *entity.Order) (err error)
	// Delete(ctx context.Context, isAdmin bool, order *entity.Order) (err error)
}
