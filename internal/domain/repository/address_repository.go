package repository

import (
	"database/sql"
	"shop-bot/internal/domain/models"
)

type AddressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (a *AddressRepository) CreateAddress(address models.Address) error {
	query := `INSERT INTO addresses (user_id, country, city, street, building, apartment, zip_code, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id`
	err := a.db.QueryRow(query, address.UserID, address.Country, address.City, address.Street, address.Building,
		address.Apartment, address.ZipCode).Scan(&address.ID)
	return err
}

func (a *AddressRepository) UpdateAddress(addressID int64, address models.Address) error {
	query := `UPDATE addresses
			  SET country = $1, city = $2, street = $3, building = $4, apartment = $5, zip_code = $6, updated_at = NOW()
			  WHERE id = $7`
	_, err := a.db.Exec(query, address.Country, address.City, address.Street, address.Building, address.Apartment, address.ZipCode, addressID)
	return err
}

func (a *AddressRepository) GetAddressesByUserID(userID int64) ([]models.Address, error) {
	query := `SELECT id, user_id, country, city, street, building, apartment, zip_code, created_at, updated_at
			  FROM addresses WHERE user_id = $1`
	rows, err := a.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var address models.Address
		if err := rows.Scan(&address.ID, &address.UserID, &address.Country, &address.City, &address.Street,
			&address.Building, &address.Apartment, &address.ZipCode, &address.CreatedAt, &address.UpdatedAt); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (a *AddressRepository) DeleteAddress(addressID int64) error {
	query := `DELETE FROM addresses WHERE id = $1`
	_, err := a.db.Exec(query, addressID)
	return err
}
