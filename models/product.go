package models

type ProductCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	CategoryID   int     `json:"category_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	UserName     string  `json:"user_name,omitempty"`
	CategoryName string  `json:"category_name,omitempty"`
}
