package models

type Address struct {
	ID        int64  `json:"id" db:"id"`
	UserID    int64  `json:"user_id" db:"user_id"`
	Country   string `json:"country" db:"country"`
	City      string `json:"city" db:"city"`
	Street    string `json:"street" db:"street"`
	Building  string `json:"building" db:"building"`
	Apartment string `json:"apartment" db:"apartment"`
	ZipCode   string `json:"zip_code" db:"zip_code"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
