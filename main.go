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
	}
}
