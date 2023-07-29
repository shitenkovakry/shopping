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

func (db *DB) ReadForPublic() (models.Products, error) {
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
