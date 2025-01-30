package unit

import (
	"shop-bot/internal/domain/models"
	"shop-bot/internal/usecase/impl"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) CreateCategory(category models.Category) (int64, error) {
	args := m.Called(category)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCategoryRepository) UpdateCategory(categoryID int64, category models.Category) error {
	args := m.Called(categoryID, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetCategoryByID(categoryID int64) (models.Category, error) {
	args := m.Called(categoryID)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetAllCategories() ([]models.Category, error) {
	args := m.Called()
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepository) DeleteCategory(categoryID int64) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func TestCategoryUseCase_CreateCategory(t *testing.T) {
	mockCategoryRepo := new(MockCategoryRepository)
	usecase := impl.NewCategoryUseCase(mockCategoryRepo)

	inputCategory := models.Category{Name: "Books"}
	mockCategoryRepo.On("CreateCategory", inputCategory).Return(int64(1), nil)

	result, err := usecase.CreateCategory(inputCategory)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), result)
	mockCategoryRepo.AssertExpectations(t)
}

func TestCategoryUseCase_GetAllCategories(t *testing.T) {
	mockCategoryRepo := new(MockCategoryRepository)
	usecase := impl.NewCategoryUseCase(mockCategoryRepo)

	mockCategories := []models.Category{
		{ID: 1, Name: "Books"},
		{ID: 2, Name: "Electronics"},
	}
	mockCategoryRepo.On("GetAllCategories").Return(mockCategories, nil)

	result, err := usecase.GetAllCategories()

	assert.Nil(t, err)
	assert.Equal(t, mockCategories, result)
	mockCategoryRepo.AssertExpectations(t)
}
