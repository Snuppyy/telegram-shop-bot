package impl

import (
	"errors"
	"shop-bot/internal/domain/models"
	"shop-bot/internal/domain/repository"
)

type UserUseCaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{userRepo: userRepo}
}

func (u *UserUseCaseImpl) SignUp(username, email, password string) (models.User, error) {
	hashedPassword := hashPassword(password)
	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}
	err := u.userRepo.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserUseCaseImpl) Login(email, password string) (models.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if !checkPassword(password, user.Password) {
		return models.User{}, errors.New("invalid credentials")
	}
	return user, nil
}

func (u *UserUseCaseImpl) GetUserByID(userID int64) (models.User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *UserUseCaseImpl) UpdateUser(userID int64, email, password string) error {
	hashedPassword := hashPassword(password)
	return u.userRepo.UpdateUser(userID, email, hashedPassword)
}

func (u *UserUseCaseImpl) DeleteUser(userID int64) error {
	return u.userRepo.DeleteUser(userID)
}

func hashPassword(password string) string {
	return password
}

func checkPassword(password, hashedPassword string) bool {
	return password == hashedPassword
}
