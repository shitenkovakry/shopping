package sql

import (
	"database/sql"
	"shopping/models"

	"github.com/pkg/errors"
)

func getPurchaseInfo(db *sql.DB) ([]models.Purchase, error) {
	rows, err := db.Query("")
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows")
	}
	defer rows.Close()

	//var purchases []models.Purchase

	return nil, err
}
