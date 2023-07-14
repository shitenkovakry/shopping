package main

import "shopping/db"

func main() {
	db := db.CreateConnection()
	defer db.Close()
}
