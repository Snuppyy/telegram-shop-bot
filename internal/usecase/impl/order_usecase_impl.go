package impl

import (
	"errors"
	"shop-bot/internal/domain/models"
	"shop-bot/internal/domain/repository"
)

type OrderUseCaseImpl struct {
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
	cartRepo      repository.CartRepository
	cartItemRepo  repository.CartItemRepository
}

func NewOrderUseCase(
	orderRepo repository.OrderRepository,
	orderItemRepo repository.OrderItemRepository,
	cartRepo repository.CartRepository,
	cartItemRepo repository.CartItemRepository,
) *OrderUseCaseImpl {
	return &OrderUseCaseImpl{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
	}
}

func (o *OrderUseCaseImpl) CreateOrder(order models.Order) (int64, error) {
	cart, err := o.cartRepo.GetCartByUserID(order.UserID)
	if err != nil {
		return 0, err
	}

	cartItems, err := o.cartItemRepo.GetCartItemsByCartID(cart.ID)
	if err != nil {
		return 0, err
	}
	if len(cartItems) == 0 {
		return 0, errors.New("cannot create order: cart is empty")
	}

	orderID, err := o.orderRepo.CreateOrder(order)
	if err != nil {
		return 0, err
	}

	for _, cartItem := range cartItems {
		orderItem := models.OrderItem{
			OrderID:   orderID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		}
		_, err := o.orderItemRepo.AddOrderItem(orderItem)
		if err != nil {
			return 0, err
		}
	}

	err = o.cartItemRepo.DeleteAllCartItems(cart.ID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (o *OrderUseCaseImpl) UpdateOrderStatus(orderID int64, status string) error {
	return o.orderRepo.UpdateOrderStatus(orderID, status)
}

func (o *OrderUseCaseImpl) GetOrderWithItems(orderID int64) (models.Order, error) {
	order, err := o.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return models.Order{}, err
	}

	orderItems, err := o.orderItemRepo.GetOrderItemsByOrderID(orderID)
	if err != nil {
		return models.Order{}, err
	}
	order.Items = orderItems

	return order, nil
}

func (o *OrderUseCaseImpl) GetOrdersByUserID(userID int64) ([]models.Order, error) {
	orders, err := o.orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}

	for i, order := range orders {
		orderItems, err := o.orderItemRepo.GetOrderItemsByOrderID(order.ID)
		if err != nil {
			return nil, err
		}
		orders[i].Items = orderItems
	}

	return orders, nil
}

func (o *OrderUseCaseImpl) DeleteOrder(orderID int64) error {
	err := o.orderItemRepo.DeleteOrderItemsByOrderID(orderID)
	if err != nil {
		return err
	}

	return o.orderRepo.DeleteOrder(orderID)
}
