package position

import (
	"context"

	"github.com/nazyli/api-restaurant/entity"
)

// Repository Position inteface
type Repository interface {
	Select(ctx context.Context, app int64, all bool) (positions entity.Positions, err error)
	GetByID(ctx context.Context, app int64, id int64, all bool) (position *entity.Position, err error)
	Insert(ctx context.Context, position *entity.Position) (err error)
	Update(ctx context.Context, position *entity.Position) (err error)
	Delete(ctx context.Context, position *entity.Position) (err error)
}
