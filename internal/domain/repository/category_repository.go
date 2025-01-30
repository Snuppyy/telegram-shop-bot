package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c *CategoryRepository) CreateCategory(category models.Category) error {
	query := `INSERT INTO categories (name, parent_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id`
	err := c.db.QueryRow(query, category.Name, category.ParentID).Scan(&category.ID)
	return err
}

func (c *CategoryRepository) UpdateCategory(categoryID int64, category models.Category) error {
	query := `UPDATE categories SET name = $1, parent_id = $2, updated_at = NOW() WHERE id = $3`
	_, err := c.db.Exec(query, category.Name, category.ParentID, categoryID)
	return err
}

func (c *CategoryRepository) GetCategoryByID(categoryID int64) (models.Category, error) {
	query := `SELECT id, name, parent_id, created_at, updated_at FROM categories WHERE id = $1`
	category := models.Category{}
	err := c.db.QueryRow(query, categoryID).
		Scan(&category.ID, &category.Name, &category.ParentID, &category.CreatedAt, &category.UpdatedAt)
	return category, err
}

func (c *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	query := `SELECT id, name, parent_id, created_at, updated_at FROM categories`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.ParentID, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *CategoryRepository) DeleteCategory(categoryID int64) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := c.db.Exec(query, categoryID)
	return err
}
