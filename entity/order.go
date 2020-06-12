package entity

import null "gopkg.in/guregu/null.v3"

// Order struct
type Order struct {
	ID            int64     `json:"id"`
	InvoiceNum    string    `json:"invoice_num"`
	SaleDate      null.Time `json:"sale_date"`
	SaleTime      null.Time `json:"sale_time"`
	SubTotal      float64   `json:"sub_total"`
	Tax           float32   `json:"tax"`
	Total         float64   `json:"total"`
	Cash          float64   `json:"cash"`
	Change        float64   `json:"change"`
	Other         string    `json:"other"`
	PaymentStatus int8      `json:"payment_status"`
	CusomerID     int8      `json:"customer_id"`
	EmployeeID    int8      `json:"employee_id"`
	AppID         int64     `json:"app_id"`
	CreatedAt     null.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedAt     null.Time `json:"updated_at"`
	LastUpdateBy  *string   `json:"last_update_by"`
	DeletedAt     null.Time `json:"deleted_at"`
	IsActive      int8      `json:"is_active"`
}

// Order list
type Orders []Order
