package buyer

import "time"

type Buyer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateBuyerRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=100"`
	Phone string `json:"phone" validate:"required,min=1,max=20"`
}

type CreateBuyerResponse struct{}

type GetBuyerRequest struct {
	Id int `json:"id"`
}

type GetBuyerResponse struct {
	Buyer Buyer `json:"buyer"`
}

type UpdateBuyerRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"omitempty,min=1,max=100"`
	Phone string `json:"phone" validate:"omitempty,min=1,max=20"`
}

type UpdateBuyerResponse struct{}

type DeleteBuyerRequest struct {
	Id int `json:"id"`
}

type DeleteBuyerResponse struct{}
