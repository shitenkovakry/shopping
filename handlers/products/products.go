package products

import (
	"database/sql"
	"log"
)

type Products struct {
	Name   string
	Price  float64
	Status string
}

type HandlerListOfProducts struct {
	DB     *sql.DB
	Logger *log.Logger
}
