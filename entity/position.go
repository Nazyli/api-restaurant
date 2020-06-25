package entity

// Position struct
type Position struct {
	ID           int64  `json:"id"`
	PositionName string `json:"position_name"`
	AppID        int64  `json:"app_id"`
	IsActive     int8   `json:"is_active"`
}

// Position list
type Positions []Position
