package item

import "online-shop-backend/internal/repository"

type Service struct {
	repo repository.Item
}

func NewService(repo repository.Item) *Service {
	return &Service{repo: repo}
}
