package db

import (
	"shopping/models"

	"github.com/pkg/errors"
)

func (db *DB) Read() (models.Products, error) {
	rows, err := db.connection.Query("select * from 'items'")
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows, typically a SELECT.")
	}
	defer rows.Close()

	var products models.Products

	for rows.Next() {
		var product *models.Product

		if err := rows.Scan(&product.Name); err != nil {
			return nil, errors.Wrapf(err, "can not convert columns read from the database into the following common Go types")
		}

		products = append(products, product)
	}

	return products, nil
}
