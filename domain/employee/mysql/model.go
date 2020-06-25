package mysql

import "gopkg.in/guregu/null.v3"

// Employee struct
type Employee struct {
	ID           int64     `db:"id"`
	PositionID   int64     `db:"position_id"`
	Name         string    `db:"name"`
	DateOfBirth  null.Time `db:"date_of_birth"`
	Address      *string   `db:"address"`
	Gender       *string   `db:"gender"`
	Email        *string   `db:"email"`
	Salary       float64   `db:"salary"`
	Bonus        *float64  `db:"bonus"`
	FromDate     null.Time `db:"from_date"`
	FinishDate   null.Time `db:"finish_date"`
	ShowEmployee int8      `db:"show_employee"`
	AppID        int64     `db:"app_id"`
	CreatedAt    null.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedAt    null.Time `db:"updated_at"`
	LastUpdateBy *string   `db:"last_update_by"`
	DeletedAt    null.Time `db:"deleted_at"`
	ImageUrl     *string   `db:"image_url"`
	ImageID      *string   `db:"image_id"`
	IsActive     int8      `db:"is_active"`
}

// Employee list
type Employees []Employee
