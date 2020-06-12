package mysql

import "gopkg.in/guregu/null.v3"

// Order struct
type Order struct {
	ID            int64     `db:"id"`
	InvoiceNum    string    `db:"invoice_num"`
	SaleDate      *string   `db:"sale_date"`
	SaleTime      *string   `db:"sale_time"`
	SubTotal      *float64  `db:"sub_total"`
	Tax           *float64  `db:"tax"`
	Total         *float64  `db:"total"`
	Cash          *float64  `db:"cash"`
	ChangeMoney   *float64  `db:"change_money"`
	Other         *string   `db:"other"`
	PaymentStatus *int8     `db:"payment_status"`
	CustomerID    *int64    `db:"customer_id"`
	EmployeeID    int64     `db:"employee_id"`
	AppID         int64     `db:"app_id"`
	CreatedAt     null.Time `db:"created_at"`
	CreatedBy     string    `db:"created_by"`
	UpdatedAt     null.Time `db:"updated_at"`
	LastUpdateBy  *string   `db:"last_update_by"`
	DeletedAt     null.Time `db:"deleted_at"`
	IsActive      int8      `db:"is_active"`
}

// Order list
type Orders []Order
