package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Buyer
	Item
	Order
	Seller
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
		Buyer:  NewBuyerRepository(db),
		Item:   NewItemRepository(db),
		Order:  NewOrderRepository(db),
		Seller: NewSellerRepository(db),
	}
}
