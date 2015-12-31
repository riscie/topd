package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/olekukonko/tablewriter"
	_ "github.com/denisenkom/go-mssqldb"
)

const dataSourceName = "SERVER=topdesk;DATABASE=TOPDESK_PROD;integrated security=true"

var db *sql.DB

func main() {

	var input = flag.String("h", "", "search a device via hostname")
	flag.Parse()

	db = initializeDB()
	defer db.Close()

	var data = getObjectByHostname(*input)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"HOSTNAME", "USER", "TYPE", "SPECIFICATION"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func getObjectByHostname(input string) [][]string {
	var data [][]string
	var naam, ref_gebruiker, objecttype, specificatie string

	rows, err := db.Query("select naam, ref_gebruiker, objecttype, specificatie FROM hardware where naam Like '%"+input+"%'")
	if err != nil {
		fmt.Println("Query Error1")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&naam, &ref_gebruiker, &objecttype, &specificatie)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Query Error2")
		}


		specificatie = shortenIfStringLongerThan(specificatie,23)
		row := []string{naam, ref_gebruiker, objecttype, specificatie}
		data = append(data, row)

	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Query Error3")
		log.Fatal(err)
	}
	return data
}

func initializeDB() *sql.DB {
	db, err := sql.Open("mssql", dataSourceName)
	if err != nil {
		fmt.Println("DB initialization Error")
		log.Fatal(err)
	}

	return db
}

func shortenIfStringLongerThan(s string, maxLength int) string{
	a := []rune(s)
	var returnValue string
	if len(a) > maxLength{
		for i:=0;i<maxLength;i++{
			returnValue+=string(a[i])
		}
		returnValue+="..."
	}
	return returnValue
}