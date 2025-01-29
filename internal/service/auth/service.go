package auth

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
	"online-shop-backend/pkg/strings"
)

type Service struct {
	repo repository.Auth
}

func NewService(repo repository.Auth) *Service {
	return &Service{repo: repo}
}

func (a *Service) SignUp(person domain.User) (int, error) {
	person.Password = strings.GeneratePasswordHash(person.Password)
	return a.repo.SignUp(person)
}
