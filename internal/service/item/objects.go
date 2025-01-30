package item

import "time"

type Item struct {
	ID          int       `json:"id,omitempty"`
	SellerID    int       `json:"seller_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=1,max=255"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type CreateItemRequest struct {
	SellerID    int     `json:"seller_id" validate:"required"`
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price" validate:"required"`
}

type CreateItemResponse struct{}

type GetItemRequest struct {
	Id int `json:"id"`
}

type GetItemResponse struct {
	Item Item `json:"item"`
}

type ListItemsRequest struct {
}

type ListItemsResponse struct {
	Items []Item `json:"items"`
}

type UpdateItemRequest struct {
	Id          int     `json:"id"`
	SellerID    int     `json:"seller_id,omitempty"`
	Name        string  `json:"name,omitempty" validate:"max=255"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type UpdateItemResponse struct{}

type DeleteItemRequest struct {
	Id int `json:"id"`
}

type DeleteItemResponse struct{}
