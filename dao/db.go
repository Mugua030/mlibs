package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	ConnectDB()
}

func ConnectDB() {
	var err error
	Db, err = sql.Open("mysql", "root:telnetdb@tcp(127.0.0.1:3306)/mlibs")
	if err != nil {
		log.Fatal(err)
	}
}
