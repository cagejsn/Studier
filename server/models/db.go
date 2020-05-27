package models

import (
    "database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// DBCon is the connection handle
	// for the database
	db *sql.DB
)

func OpenDatabaseConnection(dataSource string) {
    var err error
	db, err = sql.Open("sqlite3",
	dataSource)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}  else {
		fmt.Print("Database Connection Opened")
	}

	return
}
