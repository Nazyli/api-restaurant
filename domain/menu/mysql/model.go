package mysql

import "gopkg.in/guregu/null.v3"

// category struct
type Menu struct {
	ID           int64     `db:"id"`
	CategoryID   int64     `db:"category_id"`
	Name         string    `db:"name"`
	Price        float64   `db:"price"`
	Discount     *float64  `db:"disc"`
	ShowMenu     int8      `db:"show_menu"`
	AppID        int64     `db:"app_id"`
	CreatedAt    null.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedAt    null.Time `db:"updated_at"`
	LastUpdateBy *string   `db:"last_update_by"`
	DeletedAt    null.Time `db:"deleted_at"`
	IsActive     int8      `db:"is_active"`
}

// Customer list
type Menues []Menu
