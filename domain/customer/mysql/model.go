package mysql

import "gopkg.in/guregu/null.v3"

// Customer struct
type Customer struct {
	ID           int64     `db:"id"`
	Name         string    `db:"name"`
	Email        *string   `db:"email"`
	Addreas      *string   `db:"addreas"`
	AppID        int64     `db:"app_id"`
	CreatedAt    null.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedAt    null.Time `db:"updated_at"`
	LastUpdateBy *string   `db:"last_update_by"`
	DeletedAt    null.Time `db:"deleted_at"`
	IsActive     int8      `db:"is_active"`
}

// Customer list
type Customers []Customer
