package repository

import "github.com/jmoiron/sqlx"

type SellerRepository struct {
	db *sqlx.DB
}

func NewSellerRepository(db *sqlx.DB) *SellerRepository {
	return &SellerRepository{db: db}
}
