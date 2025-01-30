package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) CreateOrder(order models.Order) (int64, error) {
	query := `INSERT INTO orders (user_id, total_amount, status, created_at, updated_at)
			  VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	var orderID int64
	err := o.db.QueryRow(query, order.UserID, order.TotalAmount, order.Status).Scan(&orderID)
	return orderID, err
}

func (o *OrderRepository) UpdateOrderStatus(orderID int64, status string) error {
	query := `UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := o.db.Exec(query, status, orderID)
	return err
}

func (o *OrderRepository) GetOrdersByUserID(userID int64) ([]models.Order, error) {
	query := `SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders WHERE user_id = $1`
	rows, err := o.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *OrderRepository) GetOrderByID(orderID int64) (models.Order, error) {
	query := `SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders WHERE id = $1`
	order := models.Order{}
	err := o.db.QueryRow(query, orderID).
		Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	return order, err
}

func (o *OrderRepository) DeleteOrder(orderID int64) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := o.db.Exec(query, orderID)
	return err
}

func (o *OrderRepository) GetOrderWithItems(orderID int64) (models.Order, error) {
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return models.Order{}, err
	}

	orderItemRepo := NewOrderItemRepository(o.db)
	items, err := orderItemRepo.GetOrderItemsByOrderID(orderID)
	if err != nil {
		return models.Order{}, err
	}
	order.Items = items
	return order, nil
}
