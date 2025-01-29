package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
	"strings"
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

func (s SellerRepository) GetSeller(id int) (domain.Seller, error) {
	var seller domain.Seller
	query := fmt.Sprintf("SELECT id, name, phone, created_at FROM %s WHERE id = $1", SellerTable)

	err := s.db.Get(&seller, query, id)
	if err != nil {
		return seller, fmt.Errorf("could not get seller: %v", err)
	}

	return seller, nil
}

func (s SellerRepository) UpdateSeller(id int, data domain.Seller) error {
	var setClauses []string
	var params []interface{}

	if data.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", len(params)+1))
		params = append(params, data.Name)
	}

	if data.Phone != "" {
		setClauses = append(setClauses, fmt.Sprintf("phone = $%d", len(params)+1))
		params = append(params, data.Phone)
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", SellerTable,
		strings.Join(setClauses, ", "), len(params)+1)

	params = append(params, id)

	_, err := s.db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("could not update seller: %v", err)
	}

	return nil
}

func (s SellerRepository) DeleteSeller(id int) error {
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
