package order

import (
	"errors"
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Order
}

func (s *Service) AddItemToOrder(item domain.OrderItems) error {
	if item.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}

	if err := s.repo.AddItemToOrder(item); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateOrder(data domain.Order) error {
	if err := s.repo.CreateOrder(data); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetOrder(id int) (domain.Order, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (s *Service) ListOrders() ([]domain.Order, error) {
	orders, err := s.repo.ListOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *Service) UpdateOrder(id int, data domain.Order) error {
	if err := s.repo.UpdateOrder(id, data); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteOrder(id int) error {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return errors.New("order not found")
	}

	if err := s.repo.DeleteOrder(order.ID); err != nil {
		return err
	}
	return nil
}

func NewService(repo repository.Order) *Service {
	return &Service{repo: repo}
}
