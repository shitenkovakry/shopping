package products

import (
	"shopping/models"

	"github.com/pkg/errors"
)

type DB interface {
	Read() (models.Products, error)
	ReadStatusPublished() (models.Products, error)
	ReadProduct() (*models.Product, error)
}

type Products struct {
	db DB
}

func (products *Products) ListOfProductsForAdmin() (models.Products, error) {
	read, err := products.db.Read()
	if err != nil {
		return nil, errors.Wrapf(err, "can nor return list of products")
	}

	return read, nil
}

func (products *Products) ListOfProductsForPublic() (models.Products, error) {
	read, err := products.db.ReadStatusPublished()
	if err != nil {
		return nil, errors.Wrapf(err, "can nor return list of products for public")
	}

	return read, nil
}

func (products *Products) GetProduct(idProduct int) (*models.Product, error) {
	got, err := products.db.ReadProduct()
	if err != nil {
		return nil, errors.Wrapf(err, "can nor return products")
	}

	return got, nil
}

func New(db DB) *Products {
	return &Products{
		db: db,
	}
}
