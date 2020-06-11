package entity

import null "gopkg.in/guregu/null.v3"

// Order struct
type Order struct {
	ID            int64     `json:"id"`
	InvoiceNum    int64     `json:"invoice_num"`
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
	CreatedAt     null.Time `json:"created_at,omitempty"`
	CreatedBy     string    `json:"created_By,omitempty"`
	UpdatedAt     null.Time `json:"updated_at,omitempty"`
	LastUpdateBy  *string   `json:"last_update_by,omitempty"`
	DeletedAt     null.Time `json:"deleted_at,omitempty"`
	IsActive      int8      `json:"is_active,omitempty"`
}

// Order list
type Orders []Order
