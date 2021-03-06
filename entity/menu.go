package entity

import null "gopkg.in/guregu/null.v3"

// category struct
type Menu struct {
	ID           int64     `json:"id"`
	CategoryID   int64     `json:"category_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Discount     *float64  `json:"disc"`
	ShowMenu     int8      `json:"show_menu"`
	AppID        int64     `json:"app_id"`
	CreatedAt    null.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    null.Time `json:"updated_at"`
	LastUpdateBy *string   `json:"last_update_by"`
	DeletedAt    null.Time `json:"deleted_at"`
	IsActive     int8      `json:"is_active"`
}

// Customer list
type Menues []Menu
