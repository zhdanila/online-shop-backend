package domain

import "time"

type User struct {
	ID        int       `json:"id,omitempty" db:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=100" db:"username"`
	Password  string    `json:"password" validate:"required,min=6,max=255" db:"password"`
	Role      string    `json:"role" validate:"required,oneof=admin seller buyer" db:"role"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
