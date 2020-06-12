package mysql

import "gopkg.in/guregu/null.v3"

// OrderDetail struct
type OrderDetail struct {
	ID           int64     `db:"id"`
	InvoiceNum   string    `db:"invoice_num"`
	MenuID       int64     `db:"menu_id"`
	Amount       float64   `db:"amount"`
	Price        float64   `db:"price"`
	Discount     *float64  `db:"disc"`
	SubTotal     float64   `db:"sub_total"`
	AppID        int64     `db:"app_id"`
	CreatedAt    null.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedAt    null.Time `db:"updated_at"`
	LastUpdateBy *string   `db:"last_update_by"`
	DeletedAt    null.Time `db:"deleted_at"`
	IsActive     int8      `db:"is_active"`
}

// OrderDetail list
type OrderDetails []OrderDetail
