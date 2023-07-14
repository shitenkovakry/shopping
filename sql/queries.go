package sql

import (
	"database/sql"
	"shopping/models"

	"github.com/pkg/errors"
)

func GetPurchaseInfo(db *sql.DB) (models.Purchases, error) {
	var (
		purchases models.Purchases
	)

	rows, err := db.Query(`select "buyer_id", "item_id", "id_purchase" from "purchases"`)
	if err != nil {
		return nil, errors.Wrapf(err, "can not return rows")
	}

	defer rows.Close()

	for rows.Next() {
		purchase := &models.Purchase{}

		err := rows.Scan(&purchase.BuyerID, &purchase.ItemID, &purchase.IDPurchase)
		if err != nil {
			return nil, errors.Wrapf(err, "can not convert columns read from the database into the following common Go types")
		}

		purchases = append(purchases, purchase)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrapf(err, "error was encountered during iteration")
	}

	return purchases, nil
}
