package domain

import "time"

type Order struct {
	ID         int       `db:"id"`
	BuyerID    int       `db:"buyer_id"`
	TotalPrice float64   `db:"total_price"`
	CreatedAt  time.Time `db:"created_at"`
}
