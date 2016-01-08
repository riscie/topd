package model

import (
	"log"
)

// Hardware table contains the information for each Object
type Hardware struct {
	ObjectID    string `db:"naam"`
	Hostname    string `db:"hostnaam"`
	User        string `db:"ref_gebruiker"`
	Type        string `db:"objecttype"`
	Description string `db:"specificatie"`
	Location    string `db:"ref_lokatie"`
	IP          string `db:"ipadres"`
	MAC         string `db:"macadres"`
}

//FindHardware finds Hardware objects by ObjectID or Username and returns a slice of Hardware
func FindHardware(searchString string) []Hardware {
	result := []Hardware{}
	//TODO: Rewrite query with better security against sql injection
	//TODO: Check again if only active devices are shown (archive is shown as well imo)
	//TODO: Modifiy so that phones can be found as well (diffrent db table)
	//statusid = '11F18C35-FAB2-5802-86CA-B9DF68C41B8F' means the device has the status 'aktiv'
	err := db.Select(&result, "SELECT naam, hostnaam, ref_gebruiker, objecttype, specificatie, ref_lokatie, ipadres, macadres FROM hardware WHERE (naam Like '%"+searchString+"%' OR ref_gebruiker Like '%"+searchString+"%' OR ipadres Like '%"+searchString+"%' OR macadres Like '%"+searchString+"%' OR hostnaam Like '%"+searchString+"%') AND statusid = '11F18C35-FAB2-5802-86CA-B9DF68C41B8F'")
	if err != nil {
		log.Fatal(err)
	}
	return result
}
