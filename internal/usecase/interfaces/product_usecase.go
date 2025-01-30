package interfaces

import "shop-bot/internal/domain/models"

type ProductUseCase interface {
	CreateProduct(product models.Product) (int64, error)
	UpdateProduct(productID int64, product models.Product) error
	GetProductByID(productID int64) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	DeleteProduct(productID int64) error
}
