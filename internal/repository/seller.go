package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

const SellerTable = "sellers"

type SellerRepository struct {
	db *sqlx.DB
}

func (s SellerRepository) CreateSeller(data domain.Seller) error {
	query := fmt.Sprintf("INSERT INTO %s (name, phone) VALUES ($1, $2)", SellerTable)

	_, err := s.db.Exec(query, data.Name, data.Phone)
	if err != nil {
		return fmt.Errorf("could not insert seller: %v", err)
	}

	return nil
}

func (s SellerRepository) GetSeller(id string) (domain.Seller, error) {
	var seller domain.Seller
	query := fmt.Sprintf("SELECT id, name, phone, created_at FROM %s WHERE id = $1", SellerTable)

	err := s.db.Get(&seller, query, id)
	if err != nil {
		return seller, fmt.Errorf("could not get seller: %v", err)
	}

	return seller, nil
}

func (s SellerRepository) UpdateSeller(id string, data domain.Seller) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, phone = $2 WHERE id = $3", SellerTable)

	_, err := s.db.Exec(query, data.Name, data.Phone, id)
	if err != nil {
		return fmt.Errorf("could not update seller: %v", err)
	}

	return nil
}

func (s SellerRepository) DeleteSeller(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", SellerTable)

	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete seller: %v", err)
	}

	return nil
}

func NewSellerRepository(db *sqlx.DB) *SellerRepository {
	return &SellerRepository{db: db}
}
