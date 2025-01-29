package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

const ItemTable = "items"

type ItemRepository struct {
	db *sqlx.DB
}

func (r *ItemRepository) ListItems() ([]domain.Item, error) {
	var item []domain.Item
	query := fmt.Sprintf("SELECT id, seller_id, name, description, price, created_at FROM %s", ItemTable)

	err := r.db.Select(&item, query)
	if err != nil {
		return item, fmt.Errorf("could not get item: %v", err)
	}

	return item, nil
}

func (r *ItemRepository) CreateItem(data domain.Item) error {
	query := fmt.Sprintf("INSERT INTO %s (name, description, price) VALUES ($1, $2, $3)", ItemTable)

	_, err := r.db.Exec(query, data.Name, data.Description, data.Price)
	if err != nil {
		return fmt.Errorf("could not insert item: %v", err)
	}

	return nil
}

func (r *ItemRepository) GetItem(id int) (domain.Item, error) {
	var item domain.Item
	query := fmt.Sprintf("SELECT id, seller_id, name, description, price, created_at FROM %s WHERE id = $1", ItemTable)

	err := r.db.Get(&item, query, id)
	if err != nil {
		return item, fmt.Errorf("could not get item: %v", err)
	}

	return item, nil
}

func (r *ItemRepository) UpdateItem(id int, data domain.Item) error {
	query := fmt.Sprintf("UPDATE %s SET seller_id = $1, name = $2, description = $3, price = $4 WHERE id = $5", ItemTable)

	_, err := r.db.Exec(query, data.SellerID, data.Name, data.Description, data.Price, id)
	if err != nil {
		return fmt.Errorf("could not update item: %v", err)
	}

	return nil
}

func (r *ItemRepository) DeleteItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ItemTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete item: %v", err)
	}

	return nil
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}
