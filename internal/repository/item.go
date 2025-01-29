package repository

import (
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

type ItemRepository struct {
	db *sqlx.DB
}

func (i ItemRepository) CreateItem(data domain.Item) error {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) GetItem(id string) (domain.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) UpdateItem(id string, data domain.Item) error {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) DeleteItem(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}
