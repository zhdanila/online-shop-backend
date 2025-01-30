package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
	"strings"
)

const BuyerTable = "buyers"

type BuyerRepository struct {
	db *sqlx.DB
}

func (b BuyerRepository) ListBuyers() ([]domain.Buyer, error) {
	var buyers []domain.Buyer
	query := fmt.Sprintf("SELECT id, name, phone, created_at FROM %s", BuyerTable)

	err := b.db.Select(&buyers, query)
	if err != nil {
		return nil, fmt.Errorf("could not get buyer: %v", err)
	}

	return buyers, nil
}

func (b BuyerRepository) CreateBuyer(data domain.Buyer) error {
	query := fmt.Sprintf("INSERT INTO %s (name, phone) VALUES ($1, $2)", BuyerTable)

	_, err := b.db.Exec(query, data.Name, data.Phone)
	if err != nil {
		return fmt.Errorf("could not insert buyer: %v", err)
	}

	return nil
}

func (b BuyerRepository) GetBuyer(id int) (domain.Buyer, error) {
	var buyer domain.Buyer
	query := fmt.Sprintf("SELECT id, name, phone, created_at FROM %s WHERE id = $1", BuyerTable)

	err := b.db.Get(&buyer, query, id)
	if err != nil {
		return buyer, fmt.Errorf("could not get buyer: %v", err)
	}

	return buyer, nil
}

func (b BuyerRepository) UpdateBuyer(id int, data domain.Buyer) error {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", BuyerTable, strings.Join(setClauses, ", "), len(params)+1)
	params = append(params, id)

	_, err := b.db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("could not update buyer: %v", err)
	}

	return nil
}

func (b BuyerRepository) DeleteBuyer(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", BuyerTable)

	_, err := b.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete buyer: %v", err)
	}

	return nil
}

func NewBuyerRepository(db *sqlx.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}
