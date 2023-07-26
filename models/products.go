package models

type Product struct {
	ID     int
	Name   string
	Price  float64
	Status string
}

type Products []*Product
