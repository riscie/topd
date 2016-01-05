// topd is a console application which make simple search querys against a TOPdesk Database
package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

// the dsn (parameters to connect to the TOPdesk DB)
const dataSourceName = "SERVER=10.197.11.97;DATABASE=TOPDESK_PROD;integrated security=true"

// the db object needed to query the database
var db *sql.DB

//  max length of string inside a colloumn, so that that there are no linebreaks within the table output
const rowMaxLenght = 23

func main() {
	// Command line flags parsing
	flag.Usage = usage
	flag.Parse()
	var searchString string
	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	} else {
		searchString = flag.Arg(0)
	}

	// DB Initialization
	db = initializeDB()
	defer db.Close()

	// Checking database connectivity first
	err := db.Ping()
	if err != nil {
		fmt.Println("Error: No connection to the database")
		os.Exit(1)
	}

	// TODO: Adding the searchString Parser here
	// We are only getting items where we find the object-name atm...
	var results = findByInventoryName(searchString)

	// Table-Forming and output of the result
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "OBJECT-ID", "USER", "TYPE", "SPECIFICATION", "LOCATION"})

	for i, r := range results {
		//Adding a number to each line
		line := []string{strconv.Itoa(i + 1)}
		//merging the whole line together and adding it to the table
		line = append(line, r...)
		table.Append(line)
	}
	table.Render()
}

// Displays the help message to the user
func usage() {
	fmt.Fprintf(os.Stderr, `
		Usage: 	topd <search>

		Examples:
			topd nb2737 		//search by topdesk inventory number
	`)
	// TODO: adding these options to the application
	// 		topd 10.197.10.200	//search by ip
	// 		topd 01-00-5e-7f-ff-fa	//search by mac
	// 		topd Bucher		//search by the users name
	// `)
}

// Finds results by inventory name
func findByInventoryName(input string) [][]string {
	//TODO: remove code duplication
	var data [][]string
	var naam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres string

	rows, err := db.Query("select naam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres FROM hardware where naam Like '%" + input + "%'")
	if err != nil {
		fmt.Println("Query Error1")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&naam, &ref_gebruiker, &objecttype, &specificatie, &ref_lokatie, &ipadres)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Query Error2")
		}

		// reformating strings
		naam = shortenStringsLongerThan(naam, rowMaxLenght)
		ref_gebruiker = shortenStringsLongerThan(ref_gebruiker, rowMaxLenght)
		objecttype = shortenStringsLongerThan(objecttype, rowMaxLenght)
		specificatie = shortenStringsLongerThan(specificatie, rowMaxLenght)
		ref_lokatie = shortenStringsLongerThan(ref_lokatie, rowMaxLenght)
		ipadres = shortenStringsLongerThan(ipadres, rowMaxLenght)


		row := []string{naam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres}
		data = append(data, row)

	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Query Error3")
		log.Fatal(err)
	}
	return data
}

// initializes the database connection
func initializeDB() *sql.DB {
	db, err := sql.Open("mssql", dataSourceName)
	if err != nil {
		fmt.Println("DB initialization Error")
		log.Fatal(err)
	}
	return db
}

// utlilty-function to help shorten strings so that there
// are no linebreaks within the table output
func shortenStringsLongerThan(input string, maxLength int) string {
	a := []rune(input)
	if len(a) < maxLength {
		return input
	}

	var output string
	for i := 0; i < maxLength; i++ {
		output += string(a[i])
	}
	output += "..."

	return output
}
