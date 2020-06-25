package entity

// category struct
type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	AppID        int64  `json:"app_id"`
	IsActive     int8   `json:"is_active"`
}

// category list
type Categorys []Category
