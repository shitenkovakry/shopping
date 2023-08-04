package products

import (
	"shopping/models"

	"github.com/pkg/errors"
)

type DB interface {
	Read() (models.Products, error)
	ReadStatusPublished() (models.Products, error)
	ReadProduct() (*models.Product, error)
	ReadPublishedProduct() (*models.Product, error)
	Insert(newProduct *models.Product) (*models.Product, error)
	UpdatePrice(idProduct int, price float64) (*models.Product, error)
	UpdateName(idProduct int, name string) (*models.Product, error)
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

func (products *Products) GetPublishedProduct(idProduct int) (*models.Product, error) {
	got, err := products.db.ReadPublishedProduct()
	if err != nil {
		return nil, errors.Wrapf(err, "can nor return products")
	}

	return got, nil
}

func (products *Products) Create(newProduct *models.Product) (*models.Product, error) {
	created, err := products.db.Insert(newProduct)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create product")
	}

	return created, nil
}

func (products *Products) ChangePrice(idProduct int, price float64) (*models.Product, error) {
	updatedPrice, err := products.db.UpdatePrice(idProduct, price)
	if err != nil {
		return nil, errors.Wrapf(err, "can not change price")
	}

	return updatedPrice, nil
}

func (products *Products) ChangeName(idProduct int, name string) (*models.Product, error) {
	updatedName, err := products.db.UpdateName(idProduct, name)
	if err != nil {
		return nil, errors.Wrapf(err, "can not change name")
	}

	return updatedName, nil
}

func New(db DB) *Products {
	return &Products{
		db: db,
	}
}
