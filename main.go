package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:gaurav11596@tcp(127.0.0.1:3306)/TEAMPROJECT")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
