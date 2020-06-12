package entity

import null "gopkg.in/guregu/null.v3"

// Order struct
type Order struct {
	ID            int64     `json:"id"`
	InvoiceNum    string    `json:"invoice_num"`
	SaleDate      *string   `json:"sale_date"`
	SaleTime      *string   `json:"sale_time"`
	SubTotal      *float64  `json:"sub_total"`
	Tax           *float64  `json:"tax"`
	Total         *float64  `json:"total"`
	Cash          *float64  `json:"cash"`
	ChangeMoney   *float64  `json:"change_money"`
	Other         *string   `json:"other"`
	PaymentStatus *int8     `json:"payment_status"`
	CustomerID    *int64    `json:"customer_id"`
	EmployeeID    int64     `json:"employee_id"`
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

type OrderID struct {
	ID         int64     `json:"id"`
	InvoiceNum string    `json:"invoice_num"`
	EmployeeID int64     `json:"employee_id"`
	AppID      int64     `json:"app_id"`
	CreatedAt  null.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	IsActive   int8      `json:"is_active"`
}
type CalculateOrder struct {
	ID          int64         `json:"id"`
	InvoiceNum  string        `json:"invoice_num"`
	SaleDate    string        `json:"sale_date"`
	SaleTime    string        `json:"sale_time"`
	SubTotal    float64       `json:"sub_total"`
	Tax         float64       `json:"tax"`
	Total       float64       `json:"total"`
	EmployeeID  int64         `json:"employee_id"`
	AppID       int64         `json:"app_id"`
	CreatedAt   null.Time     `json:"created_at"`
	CreatedBy   string        `json:"created_by"`
	OrderDetail []OrderDetail `json:"order_detail"`
}

// Order struct
type OrderData struct {
	ID            int64         `json:"id"`
	InvoiceNum    string        `json:"invoice_num"`
	SaleDate      *string       `json:"sale_date"`
	SaleTime      *string       `json:"sale_time"`
	SubTotal      *float64      `json:"sub_total"`
	Tax           *float64      `json:"tax"`
	Total         *float64      `json:"total"`
	Cash          *float64      `json:"cash"`
	ChangeMoney   *float64      `json:"change_money"`
	Other         *string       `json:"other"`
	PaymentStatus *int8         `json:"payment_status"`
	CustomerID    *int64        `json:"customer_id"`
	EmployeeID    int64         `json:"employee_id"`
	AppID         int64         `json:"app_id"`
	CreatedAt     null.Time     `json:"created_at"`
	CreatedBy     string        `json:"created_by"`
	UpdatedAt     null.Time     `json:"updated_at"`
	LastUpdateBy  *string       `json:"last_update_by"`
	DeletedAt     null.Time     `json:"deleted_at"`
	IsActive      int8          `json:"is_active"`
	OrderDetail   []OrderDetail `json:"order_detail"`
}
