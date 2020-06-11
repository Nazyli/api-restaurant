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
	CreatedAt    null.Time `json:"created_at,omitempty"`
	CreatedBy    string    `json:"created_by,omitempty"`
	UpdatedAt    null.Time `json:"updated_at,omitempty"`
	LastUpdateBy *string   `json:"last_update_by,omitempty"`
	DeletedAt    null.Time `json:"deleted_at,omitempty"`
	IsActive     int8      `json:"is_active,omitempty"`
}

// OrderDetail list
type OrderDetails []OrderDetail
