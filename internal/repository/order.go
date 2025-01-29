package repository

import (
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

type OrderRepository struct {
	db *sqlx.DB
}

func (o OrderRepository) CreateOrder(data domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) GetOrder(id string) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) ListOrders() ([]domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) UpdateOrder(id string, data domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) DeleteOrder(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
