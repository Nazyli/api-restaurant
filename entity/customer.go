package entity

import null "gopkg.in/guregu/null.v3"

// Customer struct
type Customer struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        *string   `json:"email"`
	Addreas      *string   `json:"addreas"`
	AppID        int64     `json:"app_id"`
	CreatedAt    null.Time `json:"created_at,omitempty"`
	CreatedBy    string    `json:"created_by,omitempty"`
	UpdatedAt    null.Time `json:"updated_at,omitempty"`
	LastUpdateBy *string   `json:"last_update_by,omitempty"`
	DeletedAt    null.Time `json:"deleted_at,omitempty"`
	IsActive     int8      `json:"is_active,omitempty"`
}

// Customer list
type Customers []Customer
