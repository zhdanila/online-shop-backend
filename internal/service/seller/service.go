package seller

import "online-shop-backend/internal/repository"

type Service struct {
	repo repository.Seller
}

func NewService(repo repository.Seller) *Service {
	return &Service{repo: repo}
}
