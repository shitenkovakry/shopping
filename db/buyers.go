package db

import (
	"shopping/models"

	"github.com/pkg/errors"
)

func (db *DB) InsertBuyer(newBuyer *models.Buyer) (*models.Buyer, error) {
	var id int

	err := db.connection.QueryRow(
		`insert into "buyers" ("name", "email", "balance", "status") values ($1, $2, $3, $4) returning "id"`,
		newBuyer.Name, newBuyer.Email, newBuyer.Balance, newBuyer.Status,
	).Scan(&id)
	if err != nil {
		return nil, errors.Wrapf(err, "can not insert new user")
	}

	createdBuyer := &models.Buyer{
		ID:      id,
		Name:    newBuyer.Name,
		Email:   newBuyer.Email,
		Balance: newBuyer.Balance,
		Status:  newBuyer.Status,
	}

	return createdBuyer, nil
}

func (db *DB) UpdateNameOfBuyer(idBuyer int, name string) (*models.Buyer, error) {
	_, err := db.connection.Exec(
		`update "buyers" set "name" = $1 where "id" = $2`,
		name, idBuyer,
	)

	if err != nil {
		return nil, errors.Wrapf(err, "can not update buyer name")
	}

	updatedName := &models.Buyer{
		ID:   idBuyer,
		Name: name,
	}

	return updatedName, nil
}

func (db *DB) UpdateEmailOfBuyer(idBuyer int, email string) (*models.Buyer, error) {
	_, err := db.connection.Exec(
		`update "buyers" set "email" = $1 where "id" = $2`,
		email, idBuyer,
	)

	if err != nil {
		return nil, errors.Wrapf(err, "can not update buyer email")
	}

	updatedEmail := &models.Buyer{
		ID:    idBuyer,
		Email: email,
	}

	return updatedEmail, nil
}
