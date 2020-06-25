package entity

import (
	null "gopkg.in/guregu/null.v3"
)

// User struct
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
	UserHash     string    `json:"user_hash"`
	EmployeeID   *int64    `json:"employee_id"`
	Scope        string    `json:"scope,omitempty"`
	AppID        int64     `json:"app_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
	LastUpdateBy *string   `json:"last_update_by"`
	DeletedAt    null.Time `json:"deleted_at"`
	IsActive     int8      `json:"is_active"`
}

// Users list
type Users []User

type UserByScope struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	UserHash     string    `json:"user_hash"`
	EmployeeID   *int64    `json:"employee_id"`
	AppID        int64     `json:"app_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
	LastUpdateBy *string   `json:"last_update_by"`
	DeletedAt    null.Time `json:"deleted_at"`
	IsActive     int8      `json:"is_active"`
	Scope        []Scopes  `json:"scope"`
}
type Scopes struct {
	RoleAcess string `json:"role_access"`
}
