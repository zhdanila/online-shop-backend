package domain

import "time"

type Buyer struct {
	ID        int       `json:"id,omitempty" db:"id"`
	Name      string    `json:"name" validate:"required,min=1,max=100" db:"name"`
	Phone     string    `json:"phone" validate:"required,min=1,max=20" db:"phone"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
