package domain

import "time"

type Order struct {
	ID         int       `json:"id,omitempty" db:"id"`
	BuyerID    int       `json:"buyer_id" validate:"required" db:"buyer_id"`
	TotalPrice float64   `json:"total_price" validate:"required" db:"total_price"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
}
