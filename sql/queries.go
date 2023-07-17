package sql

import (
	"database/sql"
	"shopping/models"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func GetAllPurchasesInfo(db *sql.DB) (models.Purchases, error) {
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

func GetBalanceOfBuyer(db *sql.DB, buyerId int) (decimal.Decimal, error) {
	var balance decimal.Decimal

	row := db.QueryRow(`select "balance" from "buyers" where id = $1`, buyerId)
	err := row.Scan(&balance)
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, " no row matches the query")
	}

	return balance, nil
}

func GetPriceOfItem(db *sql.DB, itemId int) (decimal.Decimal, error) {
	var price decimal.Decimal

	row := db.QueryRow(`select "price" from "items" where id = $1`, itemId)
	err := row.Scan(&price)
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, "no row matches the query")
	}

	return price, nil
}

func UpdateBalanceInTheThirdItem(db *sql.DB, buyerId int, itemId int, balanceAfterShopping decimal.Decimal) error {
	_, err := db.Exec(`insert into "purchases" ("buyer_id", "item_id", "balance") values ($1, $2, $3)`, buyerId, itemId, balanceAfterShopping)

	return err
}

func UpdateBalanceInTheFirstItem(db *sql.DB, buyerId int, balanceAfterShopping decimal.Decimal) error {
	_, err := db.Exec(`update "buyers" set "balance" = $1 where "id" = $2`, balanceAfterShopping, buyerId)

	return err
}
