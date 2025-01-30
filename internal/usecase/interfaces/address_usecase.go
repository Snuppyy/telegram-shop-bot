package interfaces

import "shop-bot/internal/domain/models"

type AddressUseCase interface {
	CreateAddress(address models.Address) (int64, error)
	UpdateAddress(addressID int64, address models.Address) error
	GetAddressesByUserID(userID int64) ([]models.Address, error)
	DeleteAddress(addressID int64) error
}
