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
	CreateBuyer(req *buyer.CreateBuyerRequest) (*buyer.CreateBuyerResponse, error)
	GetBuyer(req *buyer.GetBuyerRequest) (*buyer.GetBuyerResponse, error)
	UpdateBuyer(req *buyer.UpdateBuyerRequest) (*buyer.UpdateBuyerResponse, error)
	DeleteBuyer(req *buyer.DeleteBuyerRequest) (*buyer.DeleteBuyerResponse, error)
}

type Item interface {
	CreateItem(data domain.Item) error
	GetItem(id int) (domain.Item, error)
	UpdateItem(id int, data domain.Item) error
	DeleteItem(id int) error
}

type Order interface {
	CreateOrder(data domain.Order) error
	GetOrder(id int) (domain.Order, error)
	ListOrders() ([]domain.Order, error)
	UpdateOrder(id int, data domain.Order) error
	DeleteOrder(id int) error
	AddItemToOrder(item domain.OrderItems) error
}

type Seller interface {
	CreateSeller(data domain.Seller) error
	GetSeller(id int) (domain.Seller, error)
	UpdateSeller(id int, data domain.Seller) error
	DeleteSeller(id int) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Buyer:  buyer.NewService(repos.Buyer),
		Item:   item.NewService(repos.Item),
		Order:  order.NewService(repos.Order),
		Seller: seller.NewService(repos.Seller),
	}
}
