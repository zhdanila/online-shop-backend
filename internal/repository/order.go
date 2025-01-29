package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

const OrderTable = "orders"
const OrderItemsTable = "order_items"

type OrderRepository struct {
	db *sqlx.DB
}

func (r *OrderRepository) AddItemToOrder(item domain.OrderItems) error {
	var orderExists bool
	err := r.db.Get(&orderExists, "SELECT EXISTS(SELECT 1 FROM orders WHERE id = $1)", item.OrderID)
	if err != nil {
		return fmt.Errorf("could not check if order exists: %v", err)
	}
	if !orderExists {
		return fmt.Errorf("order with ID %d does not exist", item.OrderID)
	}

	query := fmt.Sprintf("INSERT INTO %s (order_id, item_id, quantity, price) VALUES ($1, $2, $3, $4)", OrderItemsTable)

	_, err = r.db.Exec(query, item.OrderID, item.ItemID, item.Quantity, item.Price)
	if err != nil {
		return fmt.Errorf("could not add item to order: %v", err)
	}

	return nil
}

func (r *OrderRepository) CreateOrder(order domain.Order) error {
	query := fmt.Sprintf("INSERT INTO %s (buyer_id, total_price, created_at) VALUES ($1, $2, $3)", OrderTable)

	_, err := r.db.Exec(query, order.BuyerID, order.TotalPrice, order.CreatedAt)
	if err != nil {
		return fmt.Errorf("could not insert order: %v", err)
	}
	return nil
}

func (r *OrderRepository) GetOrder(id string) (domain.Order, error) {
	var order domain.Order
	query := fmt.Sprintf("SELECT id, buyer_id, total_price, created_at FROM %s WHERE id = $1", OrderTable)

	err := r.db.Get(&order, query, id)
	if err != nil {
		return order, fmt.Errorf("could not get order: %v", err)
	}
	return order, nil
}

func (r *OrderRepository) ListOrders() ([]domain.Order, error) {
	var orders []domain.Order
	query := fmt.Sprintf("SELECT id, buyer_id, total_price, created_at FROM %s", OrderTable)

	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, fmt.Errorf("could not get orders: %v", err)
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrder(id string, order domain.Order) error {
	query := fmt.Sprintf("UPDATE %s SET buyer_id = $1, total_price = $2 WHERE id = $3", OrderTable)

	_, err := r.db.Exec(query, order.BuyerID, order.TotalPrice, id)
	if err != nil {
		return fmt.Errorf("could not update order: %v", err)
	}
	return nil
}

func (r *OrderRepository) DeleteOrder(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", OrderTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete order: %v", err)
	}
	return nil
}

func (r *OrderRepository) CreateOrderItem(item domain.OrderItems) error {
	query := fmt.Sprintf("INSERT INTO %s (order_id, item_id, quantity, price) VALUES ($1, $2, $3, $4)", OrderItemsTable)

	_, err := r.db.Exec(query, item.OrderID, item.ItemID, item.Quantity, item.Price)
	if err != nil {
		return fmt.Errorf("could not insert order item: %v", err)
	}
	return nil
}

func (r *OrderRepository) GetOrderItems(orderID int) ([]domain.OrderItems, error) {
	var items []domain.OrderItems
	query := fmt.Sprintf("SELECT id, order_id, item_id, quantity, price FROM %s WHERE order_id = $1", OrderItemsTable)

	err := r.db.Select(&items, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("could not get order items: %v", err)
	}
	return items, nil
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
