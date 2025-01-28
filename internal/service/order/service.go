package order

import "online-shop-backend/internal/repository"

type Service struct {
	repo repository.Order
}

func NewService(repo repository.Order) *Service {
	return &Service{repo: repo}
}
