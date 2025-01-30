package repository

import (
	"database/sql"
	"errors"
	"shop-bot/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user models.User) error {
	query := `INSERT INTO users (username, email, password, role, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	err := u.db.QueryRow(query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
	return err
}

func (u *UserRepository) GetUserByID(userID int64) (models.User, error) {
	query := `SELECT id, username, email, password, role, created_at, updated_at
			  FROM users WHERE id = $1`
	user := models.User{}
	err := u.db.QueryRow(query, userID).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.User{}, errors.New("user not found")
	}
	return user, err
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, username, email, password, role, created_at, updated_at
			  FROM users WHERE email = $1`
	user := models.User{}
	err := u.db.QueryRow(query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.User{}, errors.New("user not found")
	}
	return user, err
}

func (u *UserRepository) UpdateUser(userID int64, email string, password string) error {
	query := `UPDATE users SET email = $1, password = $2, updated_at = NOW() WHERE id = $3`
	_, err := u.db.Exec(query, email, password, userID)
	return err
}

func (u *UserRepository) DeleteUser(userID int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.Exec(query, userID)
	return err
}
