// topd is a console application which make simple search querys against a TOPdesk Database
package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/riscie/topd/model"
	"github.com/riscie/topd/util"
	"os"
)

func main() {
	// Command line flags parsing
	flag.Usage = usage
	flag.Parse()
	var searchString string
	if flag.NArg() != 1 {
		//usage()
		//os.Exit(1)
		searchString = "nb276asdfadsf"
	} else {
		searchString = flag.Arg(0)
	}

	// DB Initialization
	model.InitializeDB()
	defer model.CloseDB()

	//Database Query
	//TODO: Fix error when no results
	results := model.FindHardware(searchString)

	if len(results) > 0 {
		// Table-Forming and output of the result
		table := tablewriter.NewWriter(os.Stdout)
		tableHeader := util.CreateTableHeaderFromQueryResult(results)
		tableData := util.CreateTableDataFromQueryResult(results)
		for _, t := range tableData {
			table.Append(t)
		}
		table.SetHeader(tableHeader)
		table.Render()
	} else {
		fmt.Println("Info: No Hardware found")
	}

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
