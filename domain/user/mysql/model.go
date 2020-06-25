package mysql

import (
	"gopkg.in/guregu/null.v3"
)

// User struct
type User struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	UserHash     string    `db:"user_hash"`
	EmployeeID   *int64    `db:"employee_id"`
	Scope        string    `db:"scope"`
	AppID        int64     `db:"app_id"`
	CreatedAt    null.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedAt    null.Time `db:"updated_at"`
	LastUpdateBy *string   `db:"last_update_by"`
	DeletedAt    null.Time `db:"deleted_at"`
	IsActive     int8      `db:"is_active"`
}

// Users list
type Users []User
