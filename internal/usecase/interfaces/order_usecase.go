package interfaces

import "shop-bot/internal/domain/models"

type OrderUseCase interface {
	CreateOrder(order models.Order) (int64, error)
	GetOrderWithItems(orderID int64) (models.Order, error)
	GetOrdersByUserID(userID int64) ([]models.Order, error)
	UpdateOrderStatus(orderID int64, status string) error
	DeleteOrder(orderID int64) error
}
