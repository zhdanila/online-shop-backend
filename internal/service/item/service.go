package item

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Item
}

func (s Service) CreateItem(data domain.Item) error {
	return s.repo.CreateItem(data)
}

func (s Service) GetItem(id int) (domain.Item, error) {
	return s.repo.GetItem(id)
}

func (s Service) UpdateItem(id int, data domain.Item) error {
	return s.repo.UpdateItem(id, data)
}

func (s Service) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

func NewService(repo repository.Item) *Service {
	return &Service{repo: repo}
}
