package domain

type OrderItems struct {
	ID       int     `db:"id"`
	OrderID  int     `db:"order_id"`
	ItemID   int     `db:"item_id"`
	Quantity int     `db:"quantity"`
	Price    float64 `db:"price"`
}
