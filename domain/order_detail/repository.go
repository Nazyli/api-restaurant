package order_detail

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

type Repository interface {
	Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, err error)
	SelectByInv(ctx context.Context, app int64, inv string, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (orderDetail *entity.OrderDetail, err error)
	Insert(ctx context.Context, orderDetail *entity.OrderDetail) (err error)
	Update(ctx context.Context, isAdmin bool, orderDetail *entity.OrderDetail) (err error)
	Delete(ctx context.Context, isAdmin bool, orderDetail *entity.OrderDetail) (err error)
}
