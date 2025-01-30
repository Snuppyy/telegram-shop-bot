package models

type Category struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ParentID  *int64 `json:"parent_id,omitempty" db:"parent_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
