package item

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Item
}

func (s Service) CreateItem(data domain.Item) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetItem(id string) (domain.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateItem(id string, data domain.Item) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteItem(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.Item) *Service {
	return &Service{repo: repo}
}
