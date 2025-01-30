package models

type Cart struct {
	ID        int64      `json:"id" db:"id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	Items     []CartItem `json:"items" db:"-"`
	Total     float64    `json:"total" db:"total"`
	CreatedAt string     `json:"created_at" db:"created_at"`
	UpdatedAt string     `json:"updated_at" db:"updated_at"`
}

type CartItem struct {
	ID        int64   `json:"id" db:"id"`
	CartID    int64   `json:"cart_id" db:"cart_id"`
	ProductID int64   `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
}
