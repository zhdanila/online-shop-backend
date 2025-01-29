package order

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Order
}

func (s Service) CreateOrder(data domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetOrder(id string) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListOrders() ([]domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateOrder(id string, data domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteOrder(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.Order) *Service {
	return &Service{repo: repo}
}
