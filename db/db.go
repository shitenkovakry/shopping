package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "shopping"
	password = "shopping"
	dbname   = "shopping"
)

func CreateConnection() *sql.DB {
	//connStr := "user=shopping password=shopping dbname=shopping sslmode=disable"
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "can not connect with database"))
	}

	if err := db.Ping(); err != nil {
		log.Fatal(errors.Wrapf(err, "can not ping to database"))
	}

	log.Printf("connected to database %s", dbname)

	return db
}

type DB struct {
	connection *sql.DB
}
