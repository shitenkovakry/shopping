package buyers

import (
	"shopping/models"

	"github.com/pkg/errors"
)

type DB interface {
	InsertBuyer(newBuyer *models.Buyer) (*models.Buyer, error)
	UpdateNameOfBuyer(idBuyer int, name string) (*models.Buyer, error)
	UpdateEmailOfBuyer(idBuyer int, email string) (*models.Buyer, error)
	UpdateStatusOfBuyer(idBuyer int, status string) (*models.Buyer, error)
	UpdateBalanceOfBuyerAfterShopping(idBuyer int, priceOfProduct float64) (*models.Buyer, error)
	DeleteAccount(idBuyer int) (*models.Buyer, error)
	GetBuyerByID(idBuyer int) (*models.Buyer, error)
}

type Buyers struct {
	db DB
}

func (buyers *Buyers) Register(newBuyer *models.Buyer) (*models.Buyer, error) {
	created, err := buyers.db.InsertBuyer(newBuyer)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create buyer")
	}

	return created, nil
}

func (buyers *Buyers) ChangeNameOfBuyer(idBuyer int, name string) (*models.Buyer, error) {
	updatedName, err := buyers.db.UpdateNameOfBuyer(idBuyer, name)
	if err != nil {
		return nil, errors.Wrapf(err, "can not change name")
	}

	return updatedName, nil
}

func (buyers *Buyers) ChangeEmailOfBuyer(idBuyer int, email string) (*models.Buyer, error) {
	updatedEmail, err := buyers.db.UpdateEmailOfBuyer(idBuyer, email)
	if err != nil {
		return nil, errors.Wrapf(err, "can not change email")
	}

	return updatedEmail, nil
}

func (buyers *Buyers) ChangeStatuslOfBuyer(idBuyer int, status string) (*models.Buyer, error) {
	updatedStatus, err := buyers.db.UpdateStatusOfBuyer(idBuyer, status)
	if err != nil {
		return nil, errors.Wrapf(err, "can not change status")
	}

	return updatedStatus, nil
}

func (buyers *Buyers) DeleteAccount(idBuyer int) (*models.Buyer, error) {
	deletedBuyer, err := buyers.db.DeleteAccount(idBuyer)
	if err != nil {
		return nil, errors.Wrapf(err, "can not delete buyer")
	}

	return deletedBuyer, nil
}

func (buyers *Buyers) ReplenishBalance(idBuyer int, priceOfProduct float64) (*models.Buyer, error) {
	replenishedBalance, err := buyers.db.UpdateBalanceOfBuyerAfterShopping(idBuyer, priceOfProduct)
	if err != nil {
		return nil, errors.Wrapf(err, "can not change balance after shopping")
	}

	return replenishedBalance, nil
}

func New(db DB) *Buyers {
	return &Buyers{
		db: db,
	}
}
