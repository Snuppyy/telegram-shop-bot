package impl

import (
	"shop-bot/internal/domain/models"
	"shop-bot/internal/domain/repository"
)

type CategoryUseCaseImpl struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository) *CategoryUseCaseImpl {
	return &CategoryUseCaseImpl{categoryRepo: categoryRepo}
}

func (c *CategoryUseCaseImpl) CreateCategory(category models.Category) error {
	return c.categoryRepo.CreateCategory(category)
}

func (c *CategoryUseCaseImpl) UpdateCategory(categoryID int64, category models.Category) error {
	return c.categoryRepo.UpdateCategory(categoryID, category)
}

func (c *CategoryUseCaseImpl) GetCategoryByID(categoryID int64) (models.Category, error) {
	return c.categoryRepo.GetCategoryByID(categoryID)
}

func (c *CategoryUseCaseImpl) GetAllCategories() ([]models.Category, error) {
	return c.categoryRepo.GetAllCategories()
}

func (c *CategoryUseCaseImpl) DeleteCategory(categoryID int64) error {
	return c.categoryRepo.DeleteCategory(categoryID)
}
