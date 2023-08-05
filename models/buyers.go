package models

type Buyer struct {
	ID      int
	Name    string
	Email   string
	Balance float64
	Status  string
}

type Buyers []*Buyer
