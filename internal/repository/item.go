package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
	"strings"
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
	query := fmt.Sprintf("INSERT INTO %s (seller_id, name, description, price) VALUES ($1, $2, $3, $4)", ItemTable)

	_, err := r.db.Exec(query, data.SellerID, data.Name, data.Description, data.Price)
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
	var setClauses []string
	var params []interface{}

	if data.SellerID != 0 {
		setClauses = append(setClauses, fmt.Sprintf("seller_id = $%d", len(params)+1))
		params = append(params, data.SellerID)
	}

	if data.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", len(params)+1))
		params = append(params, data.Name)
	}

	if data.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", len(params)+1))
		params = append(params, data.Description)
	}

	if data.Price != 0 {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", len(params)+1))
		params = append(params, data.Price)
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", ItemTable, strings.Join(setClauses, ", "), len(params)+1)
	params = append(params, id)

	_, err := r.db.Exec(query, params...)
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
