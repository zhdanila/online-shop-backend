package buyer

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Buyer
}

func (s Service) CreateBuyer(data domain.Buyer) error {
	return s.repo.CreateBuyer(data)
}

func (s Service) GetBuyer(id string) (domain.Buyer, error) {
	return s.repo.GetBuyer(id)
}

func (s Service) UpdateBuyer(id string, data domain.Buyer) error {
	return s.repo.UpdateBuyer(id, data)
}

func (s Service) DeleteBuyer(id string) error {
	return s.repo.DeleteBuyer(id)
}

func NewService(repo repository.Buyer) *Service {
	return &Service{repo: repo}
}
