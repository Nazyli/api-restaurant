package entity

import null "gopkg.in/guregu/null.v3"

// OrderDetail struct
type OrderDetail struct {
	ID           int64     `json:"id"`
	InvoiceNum   string    `json:"invoice_num"`
	MenuID       int64     `json:"menu_id"`
	Amount       string    `json:"amount"`
	Price        float64   `json:"price"`
	Discount     float64   `json:"disc"`
	SubTotal     float64   `json:"sub_total"`
	AppID        int64     `json:"app_id"`
	CreatedAt    null.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    null.Time `json:"updated_at"`
	LastUpdateBy *string   `json:"last_update_by"`
	DeletedAt    null.Time `json:"deleted_at"`
	IsActive     int8      `json:"is_active"`
}

// OrderDetail list
type OrderDetails []OrderDetail
