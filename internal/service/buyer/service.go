package buyer

import "online-shop-backend/internal/repository"

type Service struct {
	repo repository.Buyer
}

func NewService(repo repository.Buyer) *Service {
	return &Service{repo: repo}
}
