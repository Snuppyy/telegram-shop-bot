package unit

import (
	"shop-bot/internal/domain/models"
	"shop-bot/internal/usecase/impl"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(product models.Product) (int64, error) {
	args := m.Called(product)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockProductRepository) UpdateProduct(productID int64, product models.Product) error {
	args := m.Called(productID, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductByID(productID int64) (models.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) GetAllProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(productID int64) error {
	args := m.Called(productID)
	return args.Error(0)
}

func TestProductUseCase_CreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	usecase := impl.NewProductUseCase(mockRepo)

	inputProduct := models.Product{Name: "Laptop", Price: 1200}
	mockRepo.On("CreateProduct", inputProduct).Return(int64(1), nil)

	result, err := usecase.CreateProduct(inputProduct)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), result)
	mockRepo.AssertExpectations(t)
}

func TestProductUseCase_GetProductByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	usecase := impl.NewProductUseCase(mockRepo)

	mockProduct := models.Product{ID: 1, Name: "Laptop", Price: 1200}
	mockRepo.On("GetProductByID", int64(1)).Return(mockProduct, nil)

	result, err := usecase.GetProductByID(1)

	assert.Nil(t, err)
	assert.Equal(t, mockProduct, result)
	mockRepo.AssertExpectations(t)
}

func TestProductUseCase_GetAllProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	usecase := impl.NewProductUseCase(mockRepo)

	mockProducts := []models.Product{
		{ID: 1, Name: "Laptop", Price: 1200},
		{ID: 2, Name: "Smartphone", Price: 800},
	}
	mockRepo.On("GetAllProducts").Return(mockProducts, nil)

	result, err := usecase.GetAllProducts()

	assert.Nil(t, err)
	assert.Equal(t, mockProducts, result)
	mockRepo.AssertExpectations(t)
}
