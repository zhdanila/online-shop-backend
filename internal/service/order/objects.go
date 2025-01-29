package order

import "time"

type Order struct {
	ID         int       `json:"id,omitempty"`
	BuyerID    int       `json:"buyer_id" validate:"required"`
	TotalPrice float64   `json:"total_price" validate:"required"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

type CreateOrderRequest struct {
	BuyerID    int     `json:"buyer_id" validate:"required"`
	TotalPrice float64 `json:"total_price" validate:"required"`
}

type CreateOrderResponse struct{}

type GetOrderRequest struct {
	Id int `json:"id"`
}

type GetOrderResponse struct {
	Order Order `json:"order"`
}

type ListOrderRequest struct {
}

type ListOrdersResponse struct {
	Orders []Order `json:"orders"`
}

type UpdateOrderRequest struct {
	Id         int     `json:"id"`
	BuyerID    int     `json:"buyer_id,omitempty" validate:"required"`
	TotalPrice float64 `json:"total_price,omitempty" validate:"required"`
}

type UpdateOrderResponse struct{}

type DeleteOrderRequest struct {
	Id int `json:"id"`
}

type DeleteOrderResponse struct{}

type OrderItems struct {
	ID       int     `json:"id,omitempty"`
	OrderID  int     `json:"order_id" validate:"required"`
	ItemID   int     `json:"item_id" validate:"required"`
	Quantity int     `json:"quantity" validate:"required,min=1"`
	Price    float64 `json:"price" validate:"required"`
}

type AddItemToOrderRequest struct {
	OrderID  int     `json:"order_id" validate:"required"`
	ItemID   int     `json:"item_id" validate:"required"`
	Quantity int     `json:"quantity" validate:"required,min=1"`
	Price    float64 `json:"price" validate:"required"`
}

type AddItemToOrderResponse struct{}
