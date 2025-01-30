package impl

import (
	"errors"
	"shop-bot/internal/domain/models"
	"shop-bot/internal/domain/repository"
	"shop-bot/internal/usecase/interfaces"
)

type CartUseCaseImpl struct {
	repoCart     repository.CartRepository
	repoCartItem repository.CartItemRepository
	repoProduct  repository.ProductRepository
}

func NewCartUseCase(repoCart repository.CartRepository, repoCartItem repository.CartItemRepository, repoProduct repository.ProductRepository) interfaces.CartUseCase {
	return &CartUseCaseImpl{
		repoCart:     repoCart,
		repoCartItem: repoCartItem,
		repoProduct:  repoProduct,
	}
}

func (c *CartUseCaseImpl) GetCartByUserID(userID int64) (models.Cart, error) {
	cart, err := c.repoCart.GetCartByUserID(userID)
	if err != nil {
		return models.Cart{}, err
	}

	items, err := c.repoCartItem.GetCartItemsByCartID(cart.ID)
	if err != nil {
		return models.Cart{}, err
	}
	cart.Items = items
	return cart, nil
}

func (c *CartUseCaseImpl) AddCartItem(userID, productID int64, quantity int) error {
	product, err := c.repoProduct.GetProductByID(productID)
	if err != nil {
		return err
	}
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	cart, err := c.repoCart.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	_, err = c.repoCartItem.AddCartItem(models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     product.Price * float64(quantity),
	})
	return err
}

func (c *CartUseCaseImpl) UpdateCartItemQuantity(userID, productID int64, quantity int) error {
	return nil
}

func (c *CartUseCaseImpl) DeleteCartItem(userID, productID int64) error {
	return nil
}

func (c *CartUseCaseImpl) ClearCart(userID int64) error {
	return nil
}
