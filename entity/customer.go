package entity

import null "gopkg.in/guregu/null.v3"

// Customer struct
type Customer struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        *string   `json:"email"`
	Addreas      *string   `json:"address"`
	AppID        int64     `json:"app_id"`
	CreatedAt    null.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    null.Time `json:"updated_at"`
	LastUpdateBy *string   `json:"last_update_by"`
	DeletedAt    null.Time `json:"deleted_at"`
	IsActive     int8      `json:"is_active"`
}

// Customer list
type Customers []Customer
