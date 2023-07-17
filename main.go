package main

import (
	"fmt"
	"log"
	"shopping/db"
	"shopping/sql"

	"github.com/pkg/errors"
)

func main() {
	db := db.CreateConnection()
	defer db.Close()

	purchases, err := sql.GetAllPurchasesInfo(db)
	if err != nil {
		log.Println(errors.Wrapf(err, "can not get info of purchases"))
	}

	for _, purchase := range purchases {
		fmt.Printf("%d покупатель %d, купил товар %d\n", purchase.IDPurchase, purchase.BuyerID, purchase.ItemID)

		buyerId := purchase.BuyerID
		itemId := purchase.ItemID

		remainigBalance, err := sql.UpdateBalanceOfByerAfterShopping(db, buyerId, itemId)
		if err != nil {
			fmt.Println(errors.Wrapf(err, "can not get onformation about remainig balance"))
		}

		fmt.Println(remainigBalance)
	}

}
