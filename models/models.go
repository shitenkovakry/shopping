package models

type Purchase struct {
	BuyerID    int
	ItemID     int
	IDPurchase int
}

type Purchases []*Purchase
