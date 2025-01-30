package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) CreateProduct(product models.Product) error {
	query := `INSERT INTO products (category_id, name, description, price, stock, image_url, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id`
	err := p.db.QueryRow(query, product.CategoryID, product.Name, product.Description,
		product.Price, product.Stock, product.ImageURL).Scan(&product.ID)
	return err
}

func (p *ProductRepository) UpdateProduct(productID int64, product models.Product) error {
	query := `UPDATE products SET category_id = $1, name = $2, description = $3, price = $4, stock = $5, image_url = $6, updated_at = NOW()
			  WHERE id = $7`
	_, err := p.db.Exec(query, product.CategoryID, product.Name, product.Description,
		product.Price, product.Stock, product.ImageURL, productID)
	return err
}

func (p *ProductRepository) GetProductByID(productID int64) (models.Product, error) {
	query := `SELECT id, category_id, name, description, price, stock, image_url, created_at, updated_at
			  FROM products WHERE id = $1`
	product := models.Product{}
	err := p.db.QueryRow(query, productID).
		Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description,
			&product.Price, &product.Stock, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt)
	return product, err
}

func (p *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `SELECT id, category_id, name, description, price, stock, image_url, created_at, updated_at FROM products`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description,
			&product.Price, &product.Stock, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepository) DeleteProduct(productID int64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := p.db.Exec(query, productID)
	return err
}
