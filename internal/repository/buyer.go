package repository

import (
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

type BuyerRepository struct {
	db *sqlx.DB
}

func (b BuyerRepository) CreateBuyer(data domain.Buyer) error {
	//TODO implement me
	panic("implement me")
}

func (b BuyerRepository) GetBuyer(id string) (domain.Buyer, error) {
	//TODO implement me
	panic("implement me")
}

func (b BuyerRepository) UpdateBuyer(id string, data domain.Buyer) error {
	//TODO implement me
	panic("implement me")
}

func (b BuyerRepository) DeleteBuyer(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewBuyerRepository(db *sqlx.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}
