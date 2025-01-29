package repository

import (
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

type Repository struct {
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
}

type Seller interface {
	CreateSeller(data domain.Seller) error
	GetSeller(id string) (domain.Seller, error)
	UpdateSeller(id string, data domain.Seller) error
	DeleteSeller(id string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Buyer:  NewBuyerRepository(db),
		Item:   NewItemRepository(db),
		Order:  NewOrderRepository(db),
		Seller: NewSellerRepository(db),
	}
}
