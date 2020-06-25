package customer

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

type Repository interface {
	Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (customers entity.Customers, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (customer *entity.Customer, err error)
	Insert(ctx context.Context, customer *entity.Customer) (err error)
	Update(ctx context.Context, isAdmin bool, customer *entity.Customer) (err error)
	Delete(ctx context.Context, isAdmin bool, customer *entity.Customer) (err error)
}
