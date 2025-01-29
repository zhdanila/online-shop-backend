package domain

type OrderItems struct {
	ID       int     `json:"id,omitempty" db:"id"`
	OrderID  int     `json:"order_id" validate:"required" db:"order_id"`
	ItemID   int     `json:"item_id" validate:"required" db:"item_id"`
	Quantity int     `json:"quantity" validate:"required,min=1" db:"quantity"`
	Price    float64 `json:"price" validate:"required" db:"price"`
}
