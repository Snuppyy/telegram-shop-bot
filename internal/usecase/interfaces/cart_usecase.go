package interfaces

import "shop-bot/internal/domain/models"

type CartUseCase interface {
	GetCartByUserID(userID int64) (models.Cart, error)
	AddCartItem(userID, productID int64, quantity int) error
	UpdateCartItemQuantity(userID, productID int64, quantity int) error
	DeleteCartItem(userID, productID int64) error
	ClearCart(userID int64) error
}
