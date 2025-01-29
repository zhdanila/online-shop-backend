package domain

import "time"

type Item struct {
	ID          int       `json:"id,omitempty" db:"id"`
	SellerID    int       `json:"seller_id" validate:"required" db:"seller_id"`
	Name        string    `json:"name" validate:"required,min=1,max=255" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	Price       float64   `json:"price" validate:"required" db:"price"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
}
