package service

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
	"online-shop-backend/internal/service/buyer"
	"online-shop-backend/internal/service/item"
	"online-shop-backend/internal/service/order"
	"online-shop-backend/internal/service/seller"
)

type Service struct {
	Buyer
	Item
	Order
	Seller
}

type Buyer interface {
	CreateBuyer(data domain.Buyer) error
	GetBuyer(id string) (domain.Buyer, error)
	UpdateBuyer(id string, data domain.Buyer) error
	DeleteBuyer(id string) error
}

type Item interface {
	CreateItem(data domain.Item) error
	GetItem(id string) (domain.Item, error)
	UpdateItem(id string, data domain.Item) error
	DeleteItem(id string) error
}

type Order interface {
	CreateOrder(data domain.Order) error
	GetOrder(id string) (domain.Order, error)
	ListOrders() ([]domain.Order, error)
	UpdateOrder(id string, data domain.Order) error
	DeleteOrder(id string) error
	AddItemToOrder(item domain.OrderItems) error
}

type Seller interface {
	CreateSeller(data domain.Seller) error
	GetSeller(id string) (domain.Seller, error)
	UpdateSeller(id string, data domain.Seller) error
	DeleteSeller(id string) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Buyer:  buyer.NewService(repos.Buyer),
		Item:   item.NewService(repos.Item),
		Order:  order.NewService(repos.Order),
		Seller: seller.NewService(repos.Seller),
	}
}
