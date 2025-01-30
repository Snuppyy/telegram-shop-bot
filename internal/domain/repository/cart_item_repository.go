package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) GetCartItemsByCartID(cartID int64) ([]models.CartItem, error) {
	query := `SELECT id, cart_id, product_id, quantity, price FROM cart_items WHERE cart_id = $1`
	rows, err := c.db.Query(query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []models.CartItem{}
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (c *CartItemRepository) AddCartItem(item models.CartItem) (int64, error) {
	query := `INSERT INTO cart_items (cart_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id`
	var itemID int64
	err := c.db.QueryRow(query, item.CartID, item.ProductID, item.Quantity, item.Price).Scan(&itemID)
	return itemID, err
}

func (c *CartItemRepository) UpdateCartItem(itemID int64, quantity int, price float64) error {
	query := `UPDATE cart_items SET quantity = $1, price = $2 WHERE id = $3`
	_, err := c.db.Exec(query, quantity, price, itemID)
	return err
}

func (c *CartItemRepository) DeleteCartItem(itemID int64) error {
	query := `DELETE FROM cart_items WHERE id = $1`
	_, err := c.db.Exec(query, itemID)
	return err
}

func (c *CartItemRepository) DeleteAllCartItems(cartID int64) error {
	query := `DELETE FROM cart_items WHERE cart_id = $1`
	_, err := c.db.Exec(query, cartID)
	return err
}
