package domain

import "time"

type Item struct {
	ID          int       `db:"id"`
	SellerID    int       `db:"seller_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
}
