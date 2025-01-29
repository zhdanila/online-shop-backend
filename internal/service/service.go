package service

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
	"online-shop-backend/internal/service/auth"
	"online-shop-backend/internal/service/buyer"
	"online-shop-backend/internal/service/item"
	"online-shop-backend/internal/service/order"
	"online-shop-backend/internal/service/seller"
)

type Service struct {
	Auth
	Buyer
	Item
	Order
	Seller
}

type Auth interface {
	SignUp(person domain.User) (int, error)
}

type Buyer interface {
}

type Item interface {
}

type Order interface {
}

type Seller interface {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:   auth.NewService(repos.Auth),
		Buyer:  buyer.NewService(repos.Buyer),
		Item:   item.NewService(repos.Item),
		Order:  order.NewService(repos.Order),
		Seller: seller.NewService(repos.Seller),
	}
}
