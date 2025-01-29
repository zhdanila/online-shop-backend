package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Auth
	Buyer
	Item
	Order
	Seller
}

type Auth interface {
}

type Buyer interface {
}

type Item interface {
}

type Order interface {
}

type Seller interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:   NewAuthRepository(db),
		Buyer:  NewBuyerRepository(db),
		Item:   NewItemRepository(db),
		Order:  NewOrderRepository(db),
		Seller: NewSellerRepository(db),
	}
}
