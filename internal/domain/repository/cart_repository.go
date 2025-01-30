package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) GetCartByUserID(userID int64) (models.Cart, error) {
	query := `SELECT id, user_id, total, created_at, updated_at FROM carts WHERE user_id = $1`
	cart := models.Cart{}
	err := c.db.QueryRow(query, userID).
		Scan(&cart.ID, &cart.UserID, &cart.Total, &cart.CreatedAt, &cart.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.Cart{}, nil // у пользователя нет корзины — возвращаем пустую
	}
	return cart, err
}

func (c *CartRepository) CreateCart(userID int64, total float64) (int64, error) {
	query := `INSERT INTO carts (user_id, total, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id`
	var cartID int64
	err := c.db.QueryRow(query, userID, total).Scan(&cartID)
	return cartID, err
}

func (c *CartRepository) UpdateCartTotal(cartID int64, total float64) error {
	query := `UPDATE carts SET total = $1, updated_at = NOW() WHERE id = $2`
	_, err := c.db.Exec(query, total, cartID)
	return err
}

func (c *CartRepository) DeleteCart(cartID int64) error {
	query := `DELETE FROM carts WHERE id = $1`
	_, err := c.db.Exec(query, cartID)
	return err
}
