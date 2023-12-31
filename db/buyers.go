package db

import (
	"shopping/models"

	"github.com/pkg/errors"
)

func (db *DB) InsertBuyer(newBuyer *models.Buyer) (*models.Buyer, error) {
	var id int

	newBuyer.Status = "active"

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

func (db *DB) UpdateStatusOfBuyer(idBuyer int, status string) (*models.Buyer, error) {
	_, err := db.connection.Exec(
		`update "buyers" set "status" = $1 where "id" = $2`,
		status, idBuyer,
	)

	if err != nil {
		return nil, errors.Wrapf(err, "can not update buyer status")
	}

	updatedStatus := &models.Buyer{
		ID:     idBuyer,
		Status: status,
	}

	return updatedStatus, nil
}
func (db *DB) GetBuyerByID(idBuyer int) (*models.Buyer, error) {
	row := db.connection.QueryRow(
		`select "id", "name", "email", "balance", "status" from "buyers" where "id" = $1`,
		idBuyer,
	)

	buyer := &models.Buyer{}
	err := row.Scan(&buyer.ID, &buyer.Name, &buyer.Email, &buyer.Balance, &buyer.Status)
	if err != nil {
		return nil, errors.Wrapf(err, "can not get buyer by ID")
	}

	return buyer, nil
}

func (db *DB) DeleteAccount(idBuyer int) (*models.Buyer, error) {
	deletedBuyer, err := db.GetBuyerByID(idBuyer)
	if err != nil {
		return nil, errors.Wrapf(err, "can not get buyer before deletion")
	}

	// Удаляем аккаунт из базы данных
	_, err = db.connection.Exec(
		`delete from "buyers" where "id" = $1`,
		idBuyer,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "can not delete buyer")
	}

	_, err = db.connection.Exec(
		`update "buyers" set "status" = 'deleted' where "id" = $1`,
		idBuyer,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "can not update buyer status")
	}

	return deletedBuyer, nil
}

func (db *DB) UpdateBalanceOfBuyerAfterShopping(idBuyer int, priceOfProduct float64) (*models.Buyer, error) {
	buyer, err := db.GetBuyerByID(idBuyer)
	if err != nil {
		return nil, errors.Wrapf(err, "can not get buyer")
	}

	balanceOfBuyerBeforShopping := buyer.Balance

	if balanceOfBuyerBeforShopping < priceOfProduct {
		return nil, errors.New("insufficient balance")
	}

	balanceAfterShopping := balanceOfBuyerBeforShopping - priceOfProduct

	_, err = db.connection.Exec(
		`update "buyers" set "balance" = $1 where "id" = $2`,
		balanceAfterShopping, idBuyer,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "can not update buyer balance")
	}

	buyer.Balance = balanceAfterShopping

	return buyer, nil
}
