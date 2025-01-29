package repository

import (
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

type SellerRepository struct {
	db *sqlx.DB
}

func (s SellerRepository) CreateSeller(data domain.Seller) error {
	//TODO implement me
	panic("implement me")
}

func (s SellerRepository) GetSeller(id string) (domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (s SellerRepository) UpdateSeller(id string, data domain.Seller) error {
	//TODO implement me
	panic("implement me")
}

func (s SellerRepository) DeleteSeller(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewSellerRepository(db *sqlx.DB) *SellerRepository {
	return &SellerRepository{db: db}
}
