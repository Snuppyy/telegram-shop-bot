package models

type Product struct {
	ID          int64   `json:"id" db:"id"`
	CategoryID  int64   `json:"category_id" db:"category_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	ImageURL    string  `json:"image_url" db:"image_url"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}
