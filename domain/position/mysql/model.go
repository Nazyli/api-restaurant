package mysql

// Position struct
type Position struct {
	ID           int64  `db:"id"`
	PositionName string `db:"position_name"`
	AppID        int64  `db:"app_id"`
	IsActive     int8   `db:"is_active"`
}

// Users list
type Positions []Position
