package interfaces

import "shop-bot/internal/domain/models"

type UserUseCase interface {
	CreateUser(user models.User) (int64, error)
	GetUserByID(userID int64) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(userID int64, email, password string) error
	DeleteUser(userID int64) error
}
