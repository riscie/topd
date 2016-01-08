package util

import (
	"github.com/riscie/topd/model"
	"reflect"
	"strconv"
)

// max length of string inside a colloumn, so that that there are no linebreaks within the table output
const rowMaxLenght = 23

// CreateTableDataFromQueryResult takes a slice of Hardware and returns a slice
// of slice of string ([][]string) which will be passed to the table writer
func CreateTableDataFromQueryResult(hardware []model.Hardware) [][]string {
	var tableData [][]string
	for i, h := range hardware {
		//Using reflection to iterate over the hardware sttuct
		line := []string{}
		line = append(line, strconv.Itoa(i+1)) //adding the index
		value := reflect.Indirect(reflect.ValueOf(h))
		for i := 0; i < value.Type().NumField(); i++ {
			line = append(line, shortenStringsLongerThan(value.Field(i).String(), rowMaxLenght)) //adding each value to the table
		}
		tableData = append(tableData, line)
	}
	return tableData
}

// CreateTableHeaderFromQueryResult generates the slice of strings which is needed for the table header
func CreateTableHeaderFromQueryResult(result []model.Hardware) []string {
	//Title: "#", "OBJECT-ID", "USER", "TYPE", "SPECIFICATION", "LOCATION", "IP", "MAC"
	var tableHeader []string
	tableHeader = append(tableHeader, "  #") //Adding # for the index
	value := reflect.Indirect(reflect.ValueOf(result[0]))
	for i := 0; i < value.Type().NumField(); i++ {
		tableHeader = append(tableHeader, value.Type().Field(i).Name) //adding each variable Name to the Table Header
	}

	return tableHeader
}

// ShortenStringsLongerThan is an utlilty-function to help shorten
// strings so that there are no linebreaks within the table output
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
