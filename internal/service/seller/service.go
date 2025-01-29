package seller

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Seller
}

func (s Service) CreateSeller(data domain.Seller) error {
	return s.repo.CreateSeller(data)
}

func (s Service) GetSeller(id string) (domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateSeller(id string, data domain.Seller) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteSeller(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.Seller) *Service {
	return &Service{repo: repo}
}
