package helper

import (
	"github.com/riscie/topd/model"
	"strconv"
)

//  max length of string inside a colloumn, so that that there are no linebreaks within the table output
const rowMaxLenght = 23

// TODO: Make this dynamically from the struct, maybe return the header as well
func ProcessStructForTableOutput(hardware []model.Hardware) [][]string {
	var data [][]string
	for i,h := range hardware{
		//"#", "OBJECT-ID", "USER", "TYPE", "SPECIFICATION", "LOCATION", "IP"
		line := []string{
			strconv.Itoa(i),
			h.Name,
			shortenStringsLongerThan(h.User,rowMaxLenght),
			shortenStringsLongerThan(h.Type,rowMaxLenght),
			shortenStringsLongerThan(h.Description, rowMaxLenght),
			shortenStringsLongerThan(h.Location, rowMaxLenght),
			shortenStringsLongerThan(h.IP, rowMaxLenght),
			shortenStringsLongerThan(h.MAC, rowMaxLenght),
		}

		data = append(data, line)
	}
	return data
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
