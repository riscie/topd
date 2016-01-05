package model

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb" //database drivers need a blank import
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

//DB is the db object needed to query the database
var db *sqlx.DB

//The internal dsn (parameters to connect to the TOPdesk DB)
const dataSourceName = "SERVER=10.197.11.97;DATABASE=TOPDESK_PROD;integrated security=true"

// InitializeDB initializes the database connection
func InitializeDB() {
	var err error
	db, err = sqlx.Open("mssql", dataSourceName)
	if err != nil {
		fmt.Println("DB initialization Error")
		log.Fatal(err)
	}
	// Checking database connectivity first
	err = db.Ping()
	if err != nil {
		fmt.Println("Error: No connection to the database")
		os.Exit(1)
	}
}

//CloseDB is used to defer the closing of the db from outside the package
func CloseDB() {
	db.Close()
}
