package mysql

// category struct
type Category struct {
	ID           int64  `db:"id"`
	CategoryName string `db:"category_name"`
	AppID        int64  `db:"app_id"`
	IsActive     int8   `db:"is_active"`
}

// category list
type Categorys []Category

