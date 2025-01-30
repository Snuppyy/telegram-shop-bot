package impl

import (
	"shop-bot/internal/domain/models"
	"shop-bot/internal/domain/repository"
	"shop-bot/internal/usecase/interfaces"
)

type AddressUseCaseImpl struct {
	repo repository.AddressRepository
}

func NewAddressUseCase(repo repository.AddressRepository) interfaces.AddressUseCase {
	return &AddressUseCaseImpl{repo: repo}
}

func (a *AddressUseCaseImpl) CreateAddress(address models.Address) (int64, error) {
	return address.ID, a.repo.CreateAddress(address)
}

func (a *AddressUseCaseImpl) UpdateAddress(addressID int64, address models.Address) error {
	return a.repo.UpdateAddress(addressID, address)
}

func (a *AddressUseCaseImpl) GetAddressesByUserID(userID int64) ([]models.Address, error) {
	return a.repo.GetAddressesByUserID(userID)
}

func (a *AddressUseCaseImpl) DeleteAddress(addressID int64) error {
	return a.repo.DeleteAddress(addressID)
}
