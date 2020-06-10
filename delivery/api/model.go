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
