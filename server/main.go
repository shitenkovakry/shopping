package main

import (
	"log"
	"os"
	"shopping/db"

	"github.com/go-chi/chi/v5"
)

const (
	address = ":8080"
)

func main() {
	router := chi.NewRouter()

	dataBase := db.CreateConnection()

	defer dataBase.Close()

	logger := log.New(os.Stdout, "shopping", log.Flags())

	logger.Print("we are going to start")

}
