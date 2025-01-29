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

func (s Service) GetSeller(id int) (domain.Seller, error) {
	return s.repo.GetSeller(id)
}

func (s Service) UpdateSeller(id int, data domain.Seller) error {
	return s.repo.UpdateSeller(id, data)
}

func (s Service) DeleteSeller(id int) error {
	return s.repo.DeleteSeller(id)
}

func NewService(repo repository.Seller) *Service {
	return &Service{repo: repo}
}
