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

func New(db DB) *Buyers {
	return &Buyers{
		db: db,
	}
}
