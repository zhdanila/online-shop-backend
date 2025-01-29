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
	GetBuyer(id int) (domain.Buyer, error)
	UpdateBuyer(id int, data domain.Buyer) error
	DeleteBuyer(id int) error
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Buyer:  NewBuyerRepository(db),
		Item:   NewItemRepository(db),
		Order:  NewOrderRepository(db),
		Seller: NewSellerRepository(db),
	}
}
