package repository

import "github.com/jmoiron/sqlx"

type BuyerRepository struct {
	db *sqlx.DB
}

func NewBuyerRepository(db *sqlx.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}
