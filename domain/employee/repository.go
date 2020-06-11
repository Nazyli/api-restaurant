package employee

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

type Repository interface {
	Select(ctx context.Context, app int64, all bool, isAdmin bool, uid string) (employees entity.Employees, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool, isAdmin bool, uid string) (employee *entity.Employee, err error)
	Insert(ctx context.Context, employee *entity.Employee) (err error)
	Update(ctx context.Context, isAdmin bool, employee *entity.Employee) (err error)
	Delete(ctx context.Context, isAdmin bool, employee *entity.Employee) (err error)
}
