package model

import (
	"log"
)

// Hardware table contains the information for each Object
type Hardware struct {
	Name        string `db:"naam"`
	User        string `db:"ref_gebruiker"`
	Type        string `db:"objecttype"`
	Description string `db:"specificatie"`
	Location    string `db:"ref_lokatie"`
	IP          string `db:"ipadres"`
}

//FindByName finds Hardware objects by Name and returns a slice of Hardware
func FindByName(searchString string) ([]Hardware, error) {
	result := []Hardware{}
	err := db.Select(&result, "SELECT naam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres FROM hardware where naam Like '%"+searchString+"%'")
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

/*
//FindByInventoryName finds results by inventory name
func FindByInventoryName(input string) [][]string {
	var data [][]string
	var naam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres string

	rows, err := db.Query("SELECT naam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres FROM hardware where naam Like '%" + input + "%'")

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
*/
