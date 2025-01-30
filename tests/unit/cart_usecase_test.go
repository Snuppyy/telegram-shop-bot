package unit

import (
	"shop-bot/internal/domain/models"
	"shop-bot/internal/usecase/impl"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) GetCartByUserID(userID int64) (models.Cart, error) {
	args := m.Called(userID)
	return args.Get(0).(models.Cart), args.Error(1)
}

func (m *MockCartRepository) AddCartItem(userID, productID int64, quantity int) error {
	args := m.Called(userID, productID, quantity)
	return args.Error(0)
}

func (m *MockCartRepository) UpdateCartItemQuantity(userID, productID int64, quantity int) error {
	args := m.Called(userID, productID, quantity)
	return args.Error(0)
}

func (m *MockCartRepository) DeleteCartItem(userID, productID int64) error {
	args := m.Called(userID, productID)
	return args.Error(0)
}

func (m *MockCartRepository) ClearCart(userID int64) error {
	args := m.Called(userID)
	return args.Error(0)
}

func TestCartUseCase_GetCartByUserID(t *testing.T) {
	mockRepo := new(MockCartRepository)
	usecase := impl.NewCartUseCase(mockRepo)

	mockCart := models.Cart{ID: 1, UserID: 1}
	mockRepo.On("GetCartByUserID", int64(1)).Return(mockCart, nil)

	result, err := usecase.GetCartByUserID(1)

	assert.Nil(t, err)
	assert.Equal(t, mockCart, result)
	mockRepo.AssertExpectations(t)
}

func TestCartUseCase_AddCartItem(t *testing.T) {
	mockRepo := new(MockCartRepository)
	usecase := impl.NewCartUseCase(mockRepo)

	mockRepo.On("AddCartItem", int64(1), int64(2), 3).Return(nil)

	err := usecase.AddCartItem(1, 2, 3)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCartUseCase_ClearCart(t *testing.T) {
	mockRepo := new(MockCartRepository)
	usecase := impl.NewCartUseCase(mockRepo)

	mockRepo.On("ClearCart", int64(1)).Return(nil)

	err := usecase.ClearCart(1)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
