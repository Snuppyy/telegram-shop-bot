package models

type Order struct {
	ID              int64       `json:"id" db:"id"`
	UserID          int64       `json:"user_id" db:"user_id"`
	Items           []OrderItem `json:"items" db:"-"`
	TotalAmount     float64     `json:"total_amount" db:"total_amount"`
	Status          string      `json:"status" db:"status"`
	CreatedAt       string      `json:"created_at" db:"created_at"`
	UpdatedAt       string      `json:"updated_at" db:"updated_at"`
	ShippingAddress Address     `json:"shipping_address" db:"-"`
	BillingAddress  Address     `json:"billing_address" db:"-"`
}

type OrderItem struct {
	ID        int64   `json:"id" db:"id"`
	OrderID   int64   `json:"order_id" db:"order_id"`
	ProductID int64   `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
}
