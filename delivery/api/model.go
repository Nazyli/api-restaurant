package api

import "gopkg.in/guregu/null.v3"

//DataResponse json
type DataResponse struct {
	ID         interface{} `json:"id,omitempty"`
	Type       string      `json:"type,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}
type reqLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type reqUser struct {
	Username   string   `json:"username" validate:"required"`
	Email      string   `json:"email" validate:"required,email"`
	Password   string   `json:"password"`
	EmployeeID null.Int `json:"employee_id"`
	Scope      string   `json:"scope" validate:"required"`
}

type reqPosition struct {
	PositionName string `json:"position_name" validate:"required"`
}

type reqCategory struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type reqMenu struct {
	CategoryID int64    `json:"category_id" validate:"required"`
	Name       string   `json:"name" validate:"required"`
	Price      float64  `json:"price" validate:"required"`
	Discount   *float64 `json:"disc"`
}

type reqCustomer struct {
	Name    string  `json:"name" validate:"required"`
	Email   *string `json:"email" validate:"email"`
	Addreas *string `json:"addreas"`
}
