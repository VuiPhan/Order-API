package models

type Product struct {
	ProductID   int     `json:"product_id" db:"ProductID"`
	ProductName string  `json:"product_name" db:"ProductName"`
	Price       float32 `json:"price" db:"Price"`
	CategoryID  int     `json:"category_id" db:"CategoryID"`
}
