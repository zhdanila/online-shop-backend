package domain

import "time"

type Seller struct {
	ID        int       `json:"id,omitempty" db:"id"`
	Name      string    `json:"name,omitempty" validate:"max=100" db:"name"`
	Phone     string    `json:"phone,omitempty" validate:"max=20" db:"phone"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
