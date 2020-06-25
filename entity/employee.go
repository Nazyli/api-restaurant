package entity

import null "gopkg.in/guregu/null.v3"

// Employee struct
type Employee struct {
	ID         int64 `json:"id"`
	PositionID int64 `json:"position_id"`

	Name         string    `json:"name"`
	DateOfBirth  null.Time `json:"date_of_birth"`
	Address      *string   `json:"address"`
	Gender       *string   `json:"gender"`
	Email        *string   `json:"email"`
	Salary       float64   `json:"salary"`
	Bonus        *float64  `json:"bonus"`
	FromDate     null.Time `json:"from_date"`
	FinishDate   null.Time `json:"finish_date"`
	ShowEmployee int8      `json:"show_employee"`
	AppID        int64     `json:"app_id"`
	CreatedAt    null.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    null.Time `json:"updated_at"`
	LastUpdateBy *string   `json:"last_update_by"`
	DeletedAt    null.Time `json:"deleted_at"`
	ImageUrl     *string   `json:"image_url"`
	ImageID      *string   `json:"image_id"`
	IsActive     int8      `json:"is_active"`
}

// Employee list
type Employees []Employee
