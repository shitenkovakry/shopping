package db

import (
	"shopping/models"

	"github.com/pkg/errors"
)

func (db *DB) Read() (models.Products, error) {
	rows, err := db.connection.Query(`select "id", "name", "status" from "items"`)
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows, typically a SELECT.")
	}
	defer rows.Close()

	var products models.Products

	for rows.Next() {
		var product *models.Product

		if err := rows.Scan(&product.ID, &product.Name, product.Status); err != nil {
			return nil, errors.Wrapf(err, "can not convert columns read from the database into the following common Go types")
		}

		products = append(products, product)
	}

	return products, nil
}

func (db *DB) ReadStatusPublished() (models.Products, error) {
	rows, err := db.connection.Query(`select "id", "name" from "items" where "status" = 'publushed'`)
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows, typically a SELECT.")
	}
	defer rows.Close()

	var products models.Products

	for rows.Next() {
		var product *models.Product

		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, errors.Wrapf(err, "can not convert columns read from the database into the following common Go types")
		}

		products = append(products, product)
	}

	return products, nil
}

func (db *DB) ReadProduct() (*models.Product, error) {
	rows, err := db.connection.Query(`select "id", "name", "price", "status" from "items"`)
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows, typically a SELECT.")
	}
	defer rows.Close()

	var product *models.Product

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Status); err != nil {
			return nil, errors.Wrapf(err, "can not convert columns read from the database into the following common Go types")
		}
	}

	return product, nil
}

func (db *DB) ReadPublishedProduct() (*models.Product, error) {
	rows, err := db.connection.Query(`select "id", "name", "price" from "items" where "status" = 'published'`)
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows, typically a SELECT.")
	}
	defer rows.Close()

	var product *models.Product

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, errors.Wrapf(err, "can not convert columns read from the database into the following common Go types")
		}
	}

	return product, nil
}

func (db *DB) Insert(newProduct *models.Product) (*models.Product, error) {
	var id int

	err := db.connection.QueryRow(
		`insert into "items" ("name", "price", "status") values ($1, $2, $3) returning "id"`,
		newProduct.Name, newProduct.Price, newProduct.Status,
	).Scan(&id)
	if err != nil {
		return nil, errors.Wrapf(err, "can not insert new product")
	}

	createdProduct := &models.Product{
		ID:     id,
		Name:   newProduct.Name,
		Price:  newProduct.Price,
		Status: newProduct.Status,
	}

	return createdProduct, nil
}

func (db *DB) UpdatePrice(idProduct int, price float64) (*models.Product, error) {
	_, err := db.connection.Exec(
		`update "items" set "price" = $1 where "id" = $2`,
		price, idProduct,
	)

	if err != nil {
		return nil, errors.Wrapf(err, "can not update product price")
	}

	updatedPrice := &models.Product{
		ID:    idProduct,
		Price: price,
	}

	return updatedPrice, nil
}
