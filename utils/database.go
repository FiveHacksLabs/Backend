package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/digital_banking")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized.")
	}
	return DB
}
