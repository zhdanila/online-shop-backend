package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
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
	query := fmt.Sprintf("UPDATE %s SET name = $1, phone = $2 WHERE id = $3", BuyerTable)

	_, err := b.db.Exec(query, data.Name, data.Phone, id)
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
