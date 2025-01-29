package buyer

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Buyer
}

func (s Service) CreateBuyer(data domain.Buyer) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetBuyer(id string) (domain.Buyer, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateBuyer(id string, data domain.Buyer) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteBuyer(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.Buyer) *Service {
	return &Service{repo: repo}
}
