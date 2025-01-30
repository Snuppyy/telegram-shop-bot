package interfaces

import "shop-bot/internal/domain/models"

type CategoryUseCase interface {
	CreateCategory(category models.Category) (int64, error)
	UpdateCategory(categoryID int64, category models.Category) error
	GetCategoryByID(categoryID int64) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	DeleteCategory(categoryID int64) error
}
