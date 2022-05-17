package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var Db *sql.DB

func InitDatabase() {
	Db, _ = sql.Open("sqlite3", "./sample.db")
	row, err := Db.Query("SELECT id, name, phone FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		var phone string
		row.Scan(&id, &name, &phone)
		log.Println("Customer: ", id, name, phone)
	}
}
