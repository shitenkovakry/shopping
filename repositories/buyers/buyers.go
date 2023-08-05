package buyers

import (
	"shopping/models"

	"github.com/pkg/errors"
)

type DB interface {
	InsertBuyer(newBuyer *models.Buyer) (*models.Buyer, error)
	UpdateNameOfBuyer(idBuyer int, name string) (*models.Buyer, error)
}

type Buyers struct {
	db DB
}

func (buyers *Buyers) Create(newBuyer *models.Buyer) (*models.Buyer, error) {
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

func New(db DB) *Buyers {
	return &Buyers{
		db: db,
	}
}
