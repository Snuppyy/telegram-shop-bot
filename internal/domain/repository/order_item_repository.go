package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type OrderItemRepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (o *OrderItemRepository) AddOrderItem(item models.OrderItem) (int64, error) {
	query := `INSERT INTO order_items (order_id, product_id, quantity, price)
			  VALUES ($1, $2, $3, $4) RETURNING id`
	var itemID int64
	err := o.db.QueryRow(query, item.OrderID, item.ProductID, item.Quantity, item.Price).Scan(&itemID)
	return itemID, err
}

func (o *OrderItemRepository) GetOrderItemsByOrderID(orderID int64) ([]models.OrderItem, error) {
	query := `SELECT id, order_id, product_id, quantity, price FROM order_items WHERE order_id = $1`
	rows, err := o.db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (o *OrderItemRepository) DeleteOrderItemsByOrderID(orderID int64) error {
	query := `DELETE FROM order_items WHERE order_id = $1`
	_, err := o.db.Exec(query, orderID)
	return err
}
