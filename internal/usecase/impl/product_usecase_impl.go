package impl

import (
	"shop-bot/internal/domain/models"
	"shop-bot/internal/domain/repository"
)

type ProductUseCaseImpl struct {
	productRepo repository.ProductRepository
}

func NewProductUseCase(productRepo repository.ProductRepository) *ProductUseCaseImpl {
	return &ProductUseCaseImpl{productRepo: productRepo}
}

func (p *ProductUseCaseImpl) CreateProduct(product models.Product) error {
	return p.productRepo.CreateProduct(product)
}

func (p *ProductUseCaseImpl) UpdateProduct(productID int64, product models.Product) error {
	return p.productRepo.UpdateProduct(productID, product)
}

func (p *ProductUseCaseImpl) GetProductByID(productID int64) (models.Product, error) {
	return p.productRepo.GetProductByID(productID)
}

func (p *ProductUseCaseImpl) GetAllProducts() ([]models.Product, error) {
	return p.productRepo.GetAllProducts()
}

func (p *ProductUseCaseImpl) DeleteProduct(productID int64) error {
	return p.productRepo.DeleteProduct(productID)
}
