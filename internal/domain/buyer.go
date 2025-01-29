package domain

import "time"

type Buyer struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}
