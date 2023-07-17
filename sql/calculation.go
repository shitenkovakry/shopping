package sql

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func UpdateBalanceOfByerAfterShopping(db *sql.DB, idBuyer int, idItem int) (decimal.Decimal, error) {
	balanceOfBuyer, err := GetBalanceOfBuyer(db, idBuyer)
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, "can not get balance of buyer ", idBuyer)
	}

	price, err := GetPriceOfItem(db, idItem)
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, "can not get price of tem ", idItem)
	}

	balanceAfterShopping := balanceOfBuyer.Sub(price)

	err = UpdateBalanceInTheThirdItem(db, idBuyer, idItem, balanceAfterShopping)
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, "can not update balance in the third item")
	}

	err = UpdateBalanceInTheFirstItem(db, idBuyer, balanceAfterShopping)
	if err != nil {
		return decimal.Zero, errors.Wrapf(err, "can not update balance in the first item")
	}

	return balanceAfterShopping, nil
}
