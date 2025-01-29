package seller

import "time"

type Seller struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSellerRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=100"`
	Phone string `json:"phone" validate:"required,min=1,max=20"`
}

type CreateSellerResponse struct{}

type GetSellerRequest struct {
	Id int `json:"id"`
}

type GetSellerResponse struct {
	Seller Seller `json:"seller"`
}

type ListSellersRequest struct {
}

type ListSellersResponse struct {
	Sellers []Seller `json:"sellers"`
}

type UpdateSellerRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"omitempty,min=1,max=100"`
	Phone string `json:"phone" validate:"omitempty,min=1,max=20"`
}

type UpdateSellerResponse struct{}

type DeleteSellerRequest struct {
	Id int `json:"id"`
}

type DeleteSellerResponse struct{}
